package relay_test

import (
	"context"
	"testing"

	"example.com/frame-relay/frame"
	"example.com/frame-relay/relay"
)

func TestSnapshotDoesNotExposeStoredFrameMemory(t *testing.T) {
	r := relay.New(relay.Config{})
	if err := r.Publish(context.Background(), frame.Frame{Stream: "telemetry", Sequence: 1, Payload: []byte("ready"), Headers: map[string]string{"zone": "west"}}); err != nil {
		t.Fatal(err)
	}
	view := r.Snapshot("telemetry")
	view[0].Payload[0] = 'x'
	view[0].Headers["zone"] = "east"
	fresh := r.Snapshot("telemetry")
	if string(fresh[0].Payload) != "ready" || fresh[0].Headers["zone"] != "west" {
		t.Fatalf("snapshot mutation changed retained frame: %#v", fresh[0])
	}
}
