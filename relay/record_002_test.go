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

func TestCumulativeAcknowledgementAtMidSequenceKeepsLaterFrames(t *testing.T) {
	r := relay.New(relay.Config{})
	for _, sequence := range []uint64{1, 2, 3, 4, 5} {
		if err := r.Publish(context.Background(), frame.Frame{Stream: "telemetry", Sequence: sequence}); err != nil {
			t.Fatal(err)
		}
	}
	// Acknowledging 3 must purge 1, 2 and 3 but leave 4 and 5 pending, with
	// the watermark pinned to exactly 3. The earlier bug only deleted the
	// exact-match frame, so 1 and 2 leaked into the snapshot while the
	// watermark raced ahead of the purged state.
	if err := r.Acknowledge("telemetry", 3); err != nil {
		t.Fatal(err)
	}
	pending := r.Snapshot("telemetry")
	if len(pending) != 2 || pending[0].Sequence != 4 || pending[1].Sequence != 5 {
		t.Fatalf("unexpected pending frames: %#v", pending)
	}
	if got := r.Acknowledged("telemetry"); got != 3 {
		t.Fatalf("watermark = %d, want 3", got)
	}
}
