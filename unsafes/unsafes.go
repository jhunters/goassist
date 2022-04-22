package unsafes

import (
	"unsafe"
)

const (
	Max_Size = 2 << 31
)

// ConvertStructToArray convert struct to array
func ConvertStructToArray[T any](obj T) []byte {
	arr := (*[Max_Size]byte)(unsafe.Pointer(&obj))
	size := unsafe.Sizeof(obj)
	return arr[:size]
}

// ConvertArrayToStruct convert array to struct
func ConvertArrayToStruct[T any](bytes []byte) *T {
	ret := (*T)(unsafe.Pointer(&bytes[0]))
	return ret
}
