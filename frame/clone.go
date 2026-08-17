package frame

// Clone returns a frame with independent payload and header storage.
func Clone(source Frame) Frame {
	return Borrow(source)
}
