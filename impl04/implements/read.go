package implements

type ReadCloserImpl struct{}

func (impl *ReadCloserImpl) Read(b []byte) (int, error) {
	return 0, nil
}
