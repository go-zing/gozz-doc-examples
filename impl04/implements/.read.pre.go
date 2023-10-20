package implements

type ReadCloserImpl struct{}

func (impl *ReadCloserImpl) Read() (int, error) {
	return 0, nil
}
