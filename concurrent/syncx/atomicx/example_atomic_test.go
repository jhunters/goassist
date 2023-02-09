package atomicx_test

import (
	"fmt"

	"github.com/jhunters/goassist/concurrent/syncx/atomicx"
	"github.com/jhunters/goassist/conv"
)

func ExampleNewAtomicInt() {
	// int32
	atomInt := atomicx.NewAtomicInt(conv.ToPtr[int32](0))
	v := atomInt.Get()
	fmt.Println(v)

	// AddandGet
	v = atomInt.AddandGet(16)
	fmt.Println(v)
	v = atomInt.AddandGet(-16)
	fmt.Println(v)

	// CompareAndSet
	b := atomInt.CompareAndSet(0, -100)
	fmt.Println(atomInt.Get(), b)

	// GetAndSet
	v = atomInt.GetAndSet(100)
	fmt.Println(v, atomInt.Get())

	// IncrementAndGet
	v = atomInt.IncrementAndGet()
	fmt.Println(v)

	// decrement and get
	v = atomInt.AddandGet(-1)
	fmt.Println(v)
	// Output:
	// 0
	// 16
	// 0
	// -100 true
	// -100 100
	// 101
	// 100
}
