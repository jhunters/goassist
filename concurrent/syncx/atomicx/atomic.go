// Package atomicx provides utility api for sync atomic operation.
package atomicx

import (
	"fmt"
	"reflect"
	"sync/atomic"

	"github.com/jhunters/goassist/unsafex"
)

type AtomicSigned interface {
	~int32 | ~int64
}

// Unsigned is a constraint that permits any unsigned integer type.
// If future releases of Go add new predeclared unsigned integer types,
// this constraint will be modified to include them.
type AtomicUnsigned interface {
	~uint32 | ~uint64 | ~uintptr
}

type AtomicInteger interface {
	AtomicSigned | AtomicUnsigned
}

type AtomicInt[E AtomicInteger] struct {
	value *E
}

// NewAtomicInt return a AtomicInt object
func NewAtomicInt[E AtomicInteger](v *E) *AtomicInt[E] {
	return &AtomicInt[E]{v}
}

// Get the current value.
func (ai *AtomicInt[E]) Get() E {
	return *ai.value
}

func (ai *AtomicInt[E]) AddandGet(v E) E {
	val := reflect.ValueOf(*ai.value)

	switch val.Kind() {
	case reflect.Int32:
		var old *int32 = unsafex.As[E, int32](ai.value)
		v := atomic.AddInt32(old, int32(v))
		return *unsafex.As[int32, E](&v)
	case reflect.Int64:
		var old *int64 = unsafex.As[E, int64](ai.value)
		v := atomic.AddInt64(old, int64(v))
		return *unsafex.As[int64, E](&v)
	case reflect.Uint32:
		var old *uint32 = unsafex.As[E, uint32](ai.value)
		v := atomic.AddUint32(old, uint32(v))
		return *unsafex.As[uint32, E](&v)
	case reflect.Uint64:
		var old *uint64 = unsafex.As[E, uint64](ai.value)
		v := atomic.AddUint64(old, uint64(v))
		return *unsafex.As[uint64, E](&v)
	case reflect.Uintptr:
		var old *uintptr = unsafex.As[E, uintptr](ai.value)
		v := atomic.AddUintptr(old, uintptr(v))
		return *unsafex.As[uintptr, E](&v)
	}

	// should not go here
	panic(fmt.Sprintf("invalid value type, %s", val.Kind().String()))
}

// IncrementAndGet increments by one the current value
func (ai *AtomicInt[E]) IncrementAndGet() E {
	return ai.AddandGet(1)
}

// CompareAndSet  executes the compare-and-swap operation for an new value.
func (ai *AtomicInt[E]) CompareAndSet(expect, update E) bool {
	val := reflect.ValueOf(*ai.value)

	switch val.Kind() {
	case reflect.Int32:
		var v *int32 = unsafex.As[E, int32](ai.value)
		return atomic.CompareAndSwapInt32(v, int32(expect), int32(update))
	case reflect.Int64:
		var v *int64 = unsafex.As[E, int64](ai.value)
		return atomic.CompareAndSwapInt64(v, int64(expect), int64(update))
	case reflect.Uint32:
		var v *uint32 = unsafex.As[E, uint32](ai.value)
		return atomic.CompareAndSwapUint32(v, uint32(expect), uint32(update))
	case reflect.Uint64:
		var v *uint64 = unsafex.As[E, uint64](ai.value)
		return atomic.CompareAndSwapUint64(v, uint64(expect), uint64(update))
	case reflect.Uintptr:
		var v *uintptr = unsafex.As[E, uintptr](ai.value)
		return atomic.CompareAndSwapUintptr(v, uintptr(expect), uintptr(update))
	}
	// should not go here
	panic(fmt.Sprintf("invalid value type, %s", val.Kind().String()))
}

// Set the value
func (ai *AtomicInt[E]) Set(update E) {
	ai.GetAndSet(update)
}

// CompareAndSet  executes the compare-and-swap operation for an new value.
func (ai *AtomicInt[E]) GetAndSet(update E) E {
	val := reflect.ValueOf(*ai.value)

	switch val.Kind() {
	case reflect.Int32:
		var old *int32 = unsafex.As[E, int32](ai.value)
		v := atomic.SwapInt32(old, int32(update))
		return *unsafex.As[int32, E](&v)
	case reflect.Int64:
		var old *int64 = unsafex.As[E, int64](ai.value)
		v := atomic.SwapInt64(old, int64(update))
		return *unsafex.As[int64, E](&v)
	case reflect.Uint32:
		var old *uint32 = unsafex.As[E, uint32](ai.value)
		v := atomic.SwapUint32(old, uint32(update))
		return *unsafex.As[uint32, E](&v)
	case reflect.Uint64:
		var old *uint64 = unsafex.As[E, uint64](ai.value)
		v := atomic.SwapUint64(old, uint64(update))
		return *unsafex.As[uint64, E](&v)
	case reflect.Uintptr:
		var old *uintptr = unsafex.As[E, uintptr](ai.value)
		v := atomic.SwapUintptr(old, uintptr(update))
		return *unsafex.As[uintptr, E](&v)
	}
	// should not go here
	panic(fmt.Sprintf("invalid value type, %s", val.Kind().String()))
}
