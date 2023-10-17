package wire03

import (
	"bytes"
	"database/sql"
	"io"
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
}

// mock set injector
// +zz:wire:inject=/:set=mock
type mockString sql.NullString

// mock set string
// provide type from function
// +zz:wire:set=mock
func MockString() String {
	return "mock"
}

// mock set value
// +zz:wire:set=mock
var MockBool = true
