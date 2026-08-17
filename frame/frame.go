// Package frame defines immutable protocol frames accepted by relay.
package frame

import "fmt"

// Frame is one ordered payload in a named stream.
type Frame struct {
	Stream   string
	Sequence uint64
	Payload  []byte
	Headers  map[string]string
}

// Validate checks fields required for ordered delivery.
func (f Frame) Validate() error {
	if f.Stream == "" {
		return fmt.Errorf("frame stream is required")
	}
	if f.Sequence == 0 {
		return fmt.Errorf("frame sequence must be positive")
	}
	return nil
}

// Borrow keeps the caller-owned frame storage intact for low-overhead relaying.
func Borrow(value Frame) Frame {
	return value
}
