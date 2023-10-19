package wire02

import (
	"context"
)

//go:generate gozz run -p "wire" ./

// +zz:wire:bind=Interface
type Implement struct{}

// +zz:wire:bind=InterfaceX
// +zz:wire:bind=InterfaceX2:aop
type Interface interface {
	Foo(ctx context.Context, param int) (result int, err error)
	Bar(ctx context.Context, param int) (result int, err error)
}

type InterfaceX Interface
type InterfaceX2 Interface

// +zz:wire:inject=/
type Target struct {
	Interface
	InterfaceX
	InterfaceX2
}

func (Implement) Foo(ctx context.Context, param int) (result int, err error) {
	return
}

func (Implement) Bar(ctx context.Context, param int) (result int, err error) {
	return
}
