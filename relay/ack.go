package relay

import (
	"fmt"

	"example.com/frame-relay/frame"
)

// Acknowledge advances one stream's acknowledgement watermark.
func (r *Relay) Acknowledge(stream string, sequence uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	streamWindow := r.windows[stream]
	if streamWindow == nil {
		return fmt.Errorf("stream %q has no pending frames", stream)
	}
	return streamWindow.Ack(frame.Previous(sequence))
}
