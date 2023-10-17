// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire01

// Injectors from wire_zinject.go:

// github.com/go-zing/gozz-doc-examples/wire01.Target
func Initialize_Target() (*Target, func(), error) {
	structA := StructA{}
	structC := &StructC{
		StructA: structA,
	}
	structB := StructB{
		InterfaceC: structC,
	}
	target := &Target{
		StructA:    structA,
		StructB:    structB,
		InterfaceC: structC,
	}
	return target, func() {
	}, nil
}