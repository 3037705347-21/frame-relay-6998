package relay_test

import (
	"context"
	"testing"

	"example.com/frame-relay/frame"
	"example.com/frame-relay/relay"
)

func TestCumulativeAcknowledgementPurgesEveryEarlierFrame(t *testing.T) {
	r := relay.New(relay.Config{})
	for _, sequence := range []uint64{1, 2, 3} {
		if err := r.Publish(context.Background(), frame.Frame{Stream: "telemetry", Sequence: sequence}); err != nil {
			t.Fatal(err)
		}
	}
	if err := r.Acknowledge("telemetry", 3); err != nil {
		t.Fatal(err)
	}
	if pending := r.Snapshot("telemetry"); len(pending) != 0 {
		t.Fatalf("acknowledged frames remain pending: %#v", pending)
	}
	if got := r.Acknowledged("telemetry"); got != 3 {
		t.Fatalf("watermark = %d, want 3", got)
	}
}
