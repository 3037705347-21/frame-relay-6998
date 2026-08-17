package relay_test

import (
	"context"
	"testing"

	"example.com/frame-relay/frame"
	"example.com/frame-relay/relay"
)

func TestPublishSnapshotAndAcknowledge(t *testing.T) {
	r := relay.New(relay.Config{MaxPendingPerStream: 2})
	for _, sequence := range []uint64{2, 1} {
		if err := r.Publish(context.Background(), frame.Frame{Stream: "telemetry", Sequence: sequence}); err != nil {
			t.Fatalf("publish %d: %v", sequence, err)
		}
	}

	snapshot := r.Snapshot("telemetry")
	if len(snapshot) != 2 || snapshot[0].Sequence != 1 || snapshot[1].Sequence != 2 {
		t.Fatalf("unexpected snapshot: %#v", snapshot)
	}
	if err := r.Acknowledge("telemetry", 1); err != nil {
		t.Fatalf("acknowledge: %v", err)
	}
	if got := r.Acknowledged("telemetry"); got != 1 {
		t.Fatalf("acknowledged = %d, want 1", got)
	}
}

func TestPublishCopiesCallerOwnedData(t *testing.T) {
	r := relay.New(relay.Config{})
	payload := []byte("ready")
	headers := map[string]string{"region": "west"}
	if err := r.Publish(context.Background(), frame.Frame{Stream: "telemetry", Sequence: 1, Payload: payload, Headers: headers}); err != nil {
		t.Fatal(err)
	}
	payload[0] = 'x'
	headers["region"] = "east"
	snapshot := r.Snapshot("telemetry")
	if string(snapshot[0].Payload) != "ready" || snapshot[0].Headers["region"] != "west" {
		t.Fatalf("relay retained caller data: %#v", snapshot[0])
	}
}

func TestCancelledPublishDoesNotCreateWindow(t *testing.T) {
	r := relay.New(relay.Config{})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := r.Publish(ctx, frame.Frame{Stream: "telemetry", Sequence: 1})
	if err == nil {
		t.Fatal("expected cancellation error")
	}
	if snapshot := r.Snapshot("telemetry"); snapshot != nil {
		t.Fatalf("cancelled publish created frames: %#v", snapshot)
	}
}
