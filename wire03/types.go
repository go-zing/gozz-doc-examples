package wire03

import (
	"bytes"
	"database/sql"
	"io"

	"github.com/google/wire"
)

//go:generate gozz run -p "wire" ./

// provide value and interface value
// +zz:wire:bind=io.Writer:aop
// +zz:wire
var Buffer = &bytes.Buffer{}

// provide referenced type
// +zz:wire
type NullString nullString

type nullString sql.NullString

// use provider function to provide referenced type alias
// +zz:wire
type String = string

func ProvideString() String {
	return ""
}

// provide value from implicit type
// +zz:wire
var Bool = false

// +zz:wire:inject=/
type Target struct {
	Buffer     *bytes.Buffer
	Writer     io.Writer
	NullString NullString
	Int        int
}

// origin wire set
// +zz:wire
var Set = wire.NewSet(wire.Value(Int))

var Int = 0

// mock set injector
// +zz:wire:inject=/:set=mock
type mockString sql.NullString

// mock set string
// provide type from function
// +zz:wire:set=mock
func MockString() String {
	return "mock"
}

// mock set struct type provide fields
// +zz:wire:set=mock:field=*
type MockConfig struct{ Bool bool }

// mock set value
// +zz:wire:set=mock
var mock = &MockConfig{Bool: true}
