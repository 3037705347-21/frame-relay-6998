package relay_test

import (
	"context"
	"testing"

	"example.com/frame-relay/frame"
	"example.com/frame-relay/relay"
)

func TestCancelledPublishIsRejectedWithoutStateChange(t *testing.T) {
	r := relay.New(relay.Config{})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := r.Publish(ctx, frame.Frame{Stream: "telemetry", Sequence: 1})
	if err == nil {
		t.Fatal("cancelled publish was accepted")
	}
	if pending := r.Snapshot("telemetry"); pending != nil {
		t.Fatalf("cancelled publish retained state: %#v", pending)
	}
}
