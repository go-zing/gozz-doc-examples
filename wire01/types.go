package wire01

//go:generate gozz run -p "wire" ./

// +zz:wire
type StructA struct{}

// +zz:wire
type StructB struct {
	InterfaceC
}

// +zz:wire:bind=InterfaceC
type StructC struct {
	StructA
}

func (StructC) Foo() {}

type InterfaceC interface {
	Foo()
}

// +zz:wire:inject=./
type Target struct {
	StructA
	StructB
	InterfaceC
}
