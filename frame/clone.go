package frame

// Clone returns a frame with independent payload and header storage so that
// mutations to the copy never reach the source. It is used both on intake (to
// detach caller-owned data) and on snapshot output (to detach the relay's
// stored memory from anything handed back to callers).
func Clone(source Frame) Frame {
	copied := source
	if source.Payload != nil {
		copied.Payload = make([]byte, len(source.Payload))
		copy(copied.Payload, source.Payload)
	}
	if source.Headers != nil {
		copied.Headers = make(map[string]string, len(source.Headers))
		for key, value := range source.Headers {
			copied.Headers[key] = value
		}
	}
	return copied
}
