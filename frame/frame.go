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

// Borrow takes defensive ownership of value on intake by cloning its mutable
// storage, so the relay never aliases caller-owned memory. Callers may freely
// mutate or reuse their Payload and Headers inputs after Publish returns.
func Borrow(value Frame) Frame {
	return Clone(value)
}
