package relay

import (
	"context"
	"fmt"

	"example.com/frame-relay/frame"
	"example.com/frame-relay/window"
)

// Publish admits one frame unless the caller has cancelled its work.
func (r *Relay) Publish(ctx context.Context, value frame.Frame) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if err := value.Validate(); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	if r.closed {
		return fmt.Errorf("relay is closed")
	}
	streamWindow := r.windows[value.Stream]
	if streamWindow == nil {
		streamWindow = window.New(value.Stream)
		r.windows[value.Stream] = streamWindow
	}
	return streamWindow.Add(value, r.config.MaxPendingPerStream)
}
