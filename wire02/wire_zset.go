// Code generated by gozz:wire github.com/go-zing/gozz. DO NOT EDIT.

//go:build wireinject
// +build wireinject

package wire02

import (
	wire "github.com/google/wire"
)

var (
	_Set = wire.NewSet(
		// github.com/go-zing/gozz-doc-examples/wire02.Implement
		wire.Bind(new(Interface), new(*Implement)),
		wire.Struct(new(Implement), "*"),

		// github.com/go-zing/gozz-doc-examples/wire02.Interface
		wire.Bind(new(InterfaceX), new(Interface)),
		wire.Bind(new(_aop_InterfaceX2), new(Interface)),
		wire.Struct(new(_impl_aop_InterfaceX2), "*"),
		wire.Bind(new(InterfaceX2), new(*_impl_aop_InterfaceX2)),

		// github.com/go-zing/gozz-doc-examples/wire02.Target
		wire.Struct(new(Target), "*"),
	)
)
