/*
 * Package unsafex to provides utility api for applying unsafe package
 */
package unsafex

import (
	"reflect"
	"unsafe"
)

const (
	Max_Size = 2 << 31
)

// MappingToArray convert any type to array
func MappingToArray[T any](obj T) []byte {
	arr := (*[Max_Size]byte)(unsafe.Pointer(&obj))
	size := unsafe.Sizeof(obj)
	return arr[:size]
}

// ArrayMapping convert array to any type
func ArrayMapping[T any](bytes []byte) *T {
	ret := (*T)(unsafe.Pointer(&bytes[0]))
	return ret
}

// Slice returns a slice whose underlying array starts at ptr
func Slice[T, R any](ptr *T, size int) []R {
	ret := (*[Max_Size]R)(unsafe.Pointer(ptr))[:size]
	return ret
}

// As to convert type by unsafe way
func As[E, R any](ptr *E) *R {
	return (*R)(unsafe.Pointer(ptr))
}

// Offset return the converted target value by point offset of starting ptr
func OffsetValue[T, R any](ptr *T, offsetN int) *R {
	var r R
	return (*R)(unsafeIndex(unsafe.Pointer(ptr), 0, ValueSizeof(r), offsetN))
}

func unsafeIndex(base unsafe.Pointer, offset uintptr, elemsz uintptr, n int) unsafe.Pointer {
	return unsafe.Pointer(uintptr(base) + offset + uintptr(n)*elemsz)
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

// Copy a new value. if value is a pointer,so only pointer reference copyied. note if has sync.Mutex field will get a warnning
func Copy[E any](v *E) *E {
	n := *v
	return &n
}
