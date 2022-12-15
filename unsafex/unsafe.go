package unsafex

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

// StringToSlice base on unsafe package to convert string to []byte without copy action
// key point: copy string's Data and Len to slice's Data and Len, and append Cap value
func StringToSlice(value string) []byte {
	// create a new []byte
	var ret []byte

	// 把string的引用指向 ret的空间
	*(*string)(unsafe.Pointer(&ret)) = value

	// 设置slice的Cap值 ，用unsafe.Add操作，执行偏移操作 16个字节
	*(*int)(unsafe.Add(unsafe.Pointer(&ret), uintptr(8)*2)) = len(value)

	return ret
}

// SliceToString base on unsafe packge to convert []byte to string without copy action
// key point: copy slice's Data and Len to string's Data and Len.
func SliceToString(b []byte) string {
	if b == nil {
		return ""
	}

	// just share Slice's Data and Len content
	return *(*string)(unsafe.Pointer(&b))
}
