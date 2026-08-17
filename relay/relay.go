package relay

import (
	"fmt"
	"sync"

	"example.com/frame-relay/window"
)

// Relay owns ordered windows for all active streams.
type Relay struct {
	mu      sync.RWMutex
	config  Config
	windows map[string]*window.Window
	closed  bool
}

// New creates a relay with validated defaults.
func New(config Config) *Relay {
	normalized, err := config.normalized()
	if err != nil {
		panic(fmt.Sprintf("invalid relay config: %v", err))
	}
	return &Relay{config: normalized, windows: make(map[string]*window.Window)}
}

// Close stops future frame admission while retaining snapshots for draining callers.
func (r *Relay) Close() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.closed = true
}
