package reflectutil

import (
	"bytes"
	"reflect"
)

var byteSliceType reflect.Type = reflect.TypeOf([]byte{})

// DeepEquals returns a matcher that matches based on 'deep equality', as
// defined by the reflect package. This matcher requires that values have
// identical types to x.
func NewDeepEquals(x interface{}) *DeepEquals {
	return &DeepEquals{x}
}

type DeepEquals struct {
	x interface{}
}

func (m *DeepEquals) Matches(c interface{}) bool {
	// Make sure the types match.
	ct := reflect.TypeOf(c)
	xt := reflect.TypeOf(m.x)

	if ct != xt {
		return false
	}

	// Special case: handle byte slices more efficiently.
	cValue := reflect.ValueOf(c)
	xValue := reflect.ValueOf(m.x)

	if ct == byteSliceType && !cValue.IsNil() && !xValue.IsNil() {
		xBytes := m.x.([]byte)
		cBytes := c.([]byte)
		return bytes.Equal(cBytes, xBytes)
	}

	// Defer to the reflect package.
	if reflect.DeepEqual(m.x, c) {
		return true
	}

	// Special case: if the comparison failed because c is the nil slice, given
	// an indication of this (since its value is printed as "[]").
	if cValue.Kind() == reflect.Slice && cValue.IsNil() {
		return false
	}

	return false
}
