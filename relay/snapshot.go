package relay

import "example.com/frame-relay/frame"

// Snapshot returns an immutable view of a stream's pending frames.
func (r *Relay) Snapshot(stream string) []frame.Frame {
	r.mu.RLock()
	defer r.mu.RUnlock()
	streamWindow := r.windows[stream]
	if streamWindow == nil {
		return nil
	}
	snapshot := streamWindow.Snapshot()
	return append([]frame.Frame(nil), snapshot...)
}

// Acknowledged returns a stream watermark or zero for an unseen stream.
func (r *Relay) Acknowledged(stream string) uint64 {
	r.mu.RLock()
	defer r.mu.RUnlock()
	streamWindow := r.windows[stream]
	if streamWindow == nil {
		return 0
	}
	return streamWindow.Acknowledged()
}
