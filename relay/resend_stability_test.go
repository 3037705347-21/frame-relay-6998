package relay_test

import (
	"context"
	"testing"

	"example.com/frame-relay/frame"
	"example.com/frame-relay/relay"
)

// TestPendingFrameStaysStableAfterConsumerMutation reproduces the telemetry
// resend scenario: a caller takes a pending frame for display, mutates the
// returned snapshot, and the same pending frame must read back unchanged so
// resend payload stays trustworthy. Previously the snapshot aliased the relay's
// stored Payload/Header memory, so the mutation corrupted the pending frame.
func TestPendingFrameStaysStableAfterConsumerMutation(t *testing.T) {
	r := relay.New(relay.Config{MaxPendingPerStream: 8})
	if err := r.Publish(context.Background(), frame.Frame{
		Stream:   "telemetry",
		Sequence: 1,
		Payload:  []byte("sensor-ok"),
		Headers:  map[string]string{"region": "west"},
	}); err != nil {
		t.Fatal(err)
	}

	// Caller takes the pending frame and rewrites the returned result's
	// contents, as the consumer did before resend.
	first := r.Snapshot("telemetry")
	if len(first) != 1 {
		t.Fatalf("expected 1 pending frame, got %d", len(first))
	}
	first[0].Payload[0] = 'X'
	first[0].Payload = append(first[0].Payload, "EXTRA"...)
	first[0].Headers["region"] = "east"
	first[0].Headers["injected"] = "yes"

	// The same pending frame, read back, must be untouched.
	second := r.Snapshot("telemetry")
	if len(second) != 1 {
		t.Fatalf("expected 1 pending frame after mutation, got %d", len(second))
	}
	if got := string(second[0].Payload); got != "sensor-ok" {
		t.Fatalf("pending payload corrupted by snapshot mutation: %q", got)
	}
	if got := second[0].Headers["region"]; got != "west" {
		t.Fatalf("pending header corrupted by snapshot mutation: %q", got)
	}
	if _, leaked := second[0].Headers["injected"]; leaked {
		t.Fatalf("pending frame inherited injected header: %#v", second[0].Headers)
	}
}
