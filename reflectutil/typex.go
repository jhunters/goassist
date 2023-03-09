package reflectutil

import "reflect"

// IsByteType check if byte type
func IsByteType(v any) bool {
	typ := reflect.TypeOf(v)
	return typ.Kind() == reflect.Uint8
}

// TypeOfName return the type name of v
func TypeOfName[E any](v E) string {
	t := reflect.TypeOf(v)
	return t.Name()
}

// ConvertibleTo reports whether v is convertible to checked.
func ConvertibleTo[E any](v E, checked any) bool {
	t := reflect.TypeOf(v)
	return t.ConvertibleTo(reflect.TypeOf(checked))
}

// AssignableIfConvertibleTo do assign checked to v if type Convertibleable
// return converted value and true if convertible
func AssignIfConvertibleTo[E any](v E, checked any) (E, bool) {
	if ConvertibleTo(v, checked) {
		va := reflect.ValueOf(checked)
		newVa := va.Convert(reflect.TypeOf(v))
		r := newVa.Interface().(E)
		return r, true
	}

	return v, false
}
