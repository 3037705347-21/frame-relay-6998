package relay

import "fmt"

// Acknowledge advances one stream's cumulative acknowledgement watermark to
// sequence. Every pending frame with sequence <= sequence is purged, so the
// watermark always reflects exactly which frames have been confirmed.
func (r *Relay) Acknowledge(stream string, sequence uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	streamWindow := r.windows[stream]
	if streamWindow == nil {
		return fmt.Errorf("stream %q has no pending frames", stream)
	}
	return streamWindow.Ack(sequence)
}
