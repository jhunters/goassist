package unsafes

import (
	"unsafe"
)

const (
	Max_Size = 2 << 31
)

// StructMappingToArray convert struct to array
func StructMappingToArray[T any](obj T) []byte {
	arr := (*[Max_Size]byte)(unsafe.Pointer(&obj))
	size := unsafe.Sizeof(obj)
	return arr[:size]
}

// ArrayMappingToStruct convert array to struct
func ArrayMappingToStruct[T any](bytes []byte) *T {
	ret := (*T)(unsafe.Pointer(&bytes[0]))
	return ret
}
