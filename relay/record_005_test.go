package relay_test

import (
	"context"
	"testing"

	"example.com/frame-relay/frame"
	"example.com/frame-relay/relay"
)

func TestInvalidFrameDoesNotCreateAnAnonymousStream(t *testing.T) {
	r := relay.New(relay.Config{})
	err := r.Publish(context.Background(), frame.Frame{Sequence: 1, Payload: []byte("bad")})
	if err == nil {
		t.Fatal("frame without stream was accepted")
	}
	if pending := r.Snapshot(""); pending != nil {
		t.Fatalf("invalid frame created anonymous state: %#v", pending)
	}
}
