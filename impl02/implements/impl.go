package implements

import impl02 "github.com/go-zing/gozz-doc-examples/impl02"

var (
	_ impl02.ReadCloser = (*Impl)(nil)
)

// +zz:wire:bind=impl02.ReadCloser
type Impl struct{}

func (impl Impl) Read(b []byte) (int, error) {
	panic("not implemented")
}

func (impl Impl) Close() error {
	panic("not implemented")
}
