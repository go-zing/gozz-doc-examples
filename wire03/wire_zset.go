// Code generated by gozz:wire github.com/go-zing/gozz. DO NOT EDIT.

//go:build wireinject
// +build wireinject

package wire03

import (
	wire "github.com/google/wire"
	"io"
)

var (
	_Set = wire.NewSet(
		// github.com/go-zing/gozz-doc-examples/wire03.Buffer
		wire.InterfaceValue(new(_aop_io_Writer), Buffer),
		wire.Struct(new(_impl_aop_io_Writer), "*"),
		wire.Bind(new(io.Writer), new(*_impl_aop_io_Writer)),
		wire.Value(Buffer),

		// github.com/go-zing/gozz-doc-examples/wire03.NullString
		wire.Struct(new(NullString), "*"),

		// github.com/go-zing/gozz-doc-examples/wire03.String
		ProvideString,

		// github.com/go-zing/gozz-doc-examples/wire03.Bool
		wire.Value(Bool),

		// github.com/go-zing/gozz-doc-examples/wire03.Target
		wire.Struct(new(Target), "*"),
	)

	_mockSet = wire.NewSet(
		// github.com/go-zing/gozz-doc-examples/wire03.mockString
		wire.Struct(new(mockString), "*"),

		// github.com/go-zing/gozz-doc-examples/wire03.MockString
		MockString,

		// github.com/go-zing/gozz-doc-examples/wire03.MockConfig
		wire.FieldsOf(new(MockConfig), "Bool"),

		// github.com/go-zing/gozz-doc-examples/wire03.mock
		wire.Value(mock),
	)
)
