package reflectutil

import (
	"bytes"
	"fmt"
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

func (m *DeepEquals) Description() string {
	xDesc := fmt.Sprintf("%v", m.x)
	xValue := reflect.ValueOf(m.x)

	// Special case: fmt.Sprintf presents nil slices as "[]", but
	// reflect.DeepEqual makes a distinction between nil and empty slices. Make
	// this less confusing.
	if xValue.Kind() == reflect.Slice && xValue.IsNil() {
		xDesc = "<nil slice>"
	}

	return fmt.Sprintf("deep equals: %s", xDesc)
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

		if bytes.Equal(cBytes, xBytes) {
			return true
		}

		return false
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
