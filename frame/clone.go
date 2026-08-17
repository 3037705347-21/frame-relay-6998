package frame

// Clone returns a frame with independent payload and header storage.
func Clone(source Frame) Frame {
	copy := Frame{
		Stream:   source.Stream,
		Sequence: source.Sequence,
		Payload:  append([]byte(nil), source.Payload...),
	}
	if len(source.Headers) != 0 {
		copy.Headers = make(map[string]string, len(source.Headers))
		for key, value := range source.Headers {
			copy.Headers[key] = value
		}
	}
	return copy
}
