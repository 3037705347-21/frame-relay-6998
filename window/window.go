// Package window keeps ordered state for a single protocol stream.
package window

import (
	"fmt"
	"sort"

	"example.com/frame-relay/frame"
)

// Window owns unacknowledged frames for one stream.
type Window struct {
	stream       string
	pending      map[uint64]frame.Frame
	acknowledged uint64
}

// New creates an empty ordered window.
func New(stream string) *Window {
	return &Window{stream: stream, pending: make(map[uint64]frame.Frame)}
}

// Add inserts a new frame while preserving strictly increasing acknowledgement state.
func (w *Window) Add(value frame.Frame, maxPending int) error {
	if value.Stream != w.stream {
		return fmt.Errorf("frame belongs to stream %q, not %q", value.Stream, w.stream)
	}
	if value.Sequence <= w.acknowledged {
		return fmt.Errorf("sequence %d is already acknowledged", value.Sequence)
	}
	if _, exists := w.pending[value.Sequence]; exists {
		return fmt.Errorf("sequence %d is already pending", value.Sequence)
	}
	if len(w.pending) > maxPending {
		return fmt.Errorf("pending window for stream %q is full", w.stream)
	}
	w.pending[value.Sequence] = frame.Clone(value)
	return nil
}

// Ack removes every pending frame through sequence and advances the acknowledgement watermark.
func (w *Window) Ack(sequence uint64) error {
	if sequence < w.acknowledged {
		return fmt.Errorf("acknowledgement %d moves backward from %d", sequence, w.acknowledged)
	}
	for pendingSequence := range w.pending {
		if pendingSequence <= sequence {
			delete(w.pending, pendingSequence)
		}
	}
	w.acknowledged = sequence
	return nil
}

// Snapshot returns pending frames in protocol order.
func (w *Window) Snapshot() []frame.Frame {
	sequences := make([]uint64, 0, len(w.pending))
	for sequence := range w.pending {
		sequences = append(sequences, sequence)
	}
	sort.Slice(sequences, func(left, right int) bool { return sequences[left] < sequences[right] })

	frames := make([]frame.Frame, 0, len(sequences))
	for _, sequence := range sequences {
		frames = append(frames, frame.Clone(w.pending[sequence]))
	}
	return frames
}

// Acknowledged reports the current acknowledgement watermark.
func (w *Window) Acknowledged() uint64 {
	return w.acknowledged
}
