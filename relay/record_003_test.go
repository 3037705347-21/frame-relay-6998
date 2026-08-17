package relay_test

import (
	"context"
	"testing"

	"example.com/frame-relay/frame"
	"example.com/frame-relay/relay"
)

func TestAdmissionHonorsCapacityAndClose(t *testing.T) {
	r := relay.New(relay.Config{MaxPendingPerStream: 1})
	first := frame.Frame{Stream: "telemetry", Sequence: 1}
	if err := r.Publish(context.Background(), first); err != nil {
		t.Fatal(err)
	}
	if err := r.Publish(context.Background(), frame.Frame{Stream: "telemetry", Sequence: 2}); err == nil {
		t.Fatal("second frame bypassed the configured capacity")
	}
	r.Close()
	if err := r.Publish(context.Background(), frame.Frame{Stream: "telemetry", Sequence: 3}); err == nil {
		t.Fatal("closed relay admitted a frame")
	}
}
