package unsafes

import (
	"reflect"
	"unsafe"
)

const (
	Max_Size = 2 << 31
)

// StructMappingToArray convert struct to array
func MappingToArray[T any](obj T) []byte {
	arr := (*[Max_Size]byte)(unsafe.Pointer(&obj))
	size := unsafe.Sizeof(obj)
	return arr[:size]
}

// ArrayMappingToStruct convert array to struct
func ArrayMapping[T any](bytes []byte) *T {
	ret := (*T)(unsafe.Pointer(&bytes[0]))
	return ret
}

// ValueSizeof to cal real value size
func ValueSizeof(v any) uintptr {
	typ := reflect.TypeOf(v)
	if typ.Kind() == reflect.Pointer {
		return typ.Elem().Size()
	}

	return typ.Size()
}
