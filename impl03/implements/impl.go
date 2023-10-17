package implements

import "github.com/go-zing/gozz-doc-examples/impl03"

type ReadCloserImpl struct{}

func (impl *ReadCloserImpl) Read() {}

var (
	_ impl03.ReadCloser = (*Impl)(nil)
)

// +zz:wire:bind=impl03.ReadCloser
type Impl struct{}

func (impl Impl) Read(b []byte) (int, error) {
	panic("not implemented")
}

func (impl Impl) Close() error {
	panic("not implemented")
}
