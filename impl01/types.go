package impl01

//go:generate gozz run -p "impl" ./

// +zz:impl:./implements
type ReadCloser interface {
	Read(b []byte) (int, error)
	Close() error
}
