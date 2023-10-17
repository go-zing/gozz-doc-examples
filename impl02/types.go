package impl02

//go:generate gozz run -p "impl" ./

// +zz:impl:./implements:type=Impl:wire
type ReadCloser interface {
	Read(b []byte) (int, error)
	Close() error
}
