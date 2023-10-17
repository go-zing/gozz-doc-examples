package implements

type ReadCloserImpl struct{}

func (impl *ReadCloserImpl) Read(b []byte) (int, error) {
	return len(b), nil
}

func (impl *ReadCloserImpl) Close() error {
	panic("not implemented")
}
