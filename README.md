# Frame Relay

Frame Relay is an in-memory Go library for coordinating ordered protocol frames
across independent streams. It is intended for services that need bounded frame
buffers, acknowledgement tracking, and immutable snapshots without a user
interface or an external dependency.

## Packages

- `frame`: protocol frame values and defensive copy helpers.
- `window`: ordered per-stream pending and acknowledged frame state.
- `relay`: the public coordinator that admits frames, acknowledges them, and
  exposes read-only snapshots.

## Usage

```go
r := relay.New(relay.Config{MaxPendingPerStream: 64})
err := r.Publish(ctx, frame.Frame{Stream: "telemetry", Sequence: 1, Payload: []byte("ok")})
```

## Commands

```text
go build ./...
go test ./...
```

The library has no required environment variables.
