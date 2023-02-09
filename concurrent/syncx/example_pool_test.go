package syncx_test

import (
	"fmt"

	"github.com/jhunters/goassist/concurrent/syncx"
)

type ExamplePoolPojo struct {
	name string
}

func ExampleNewPool() {
	name1 := "matt"
	name2 := "matthew"
	p := syncx.NewPool(func() *ExamplePoolPojo {
		return &ExamplePoolPojo{name1}
	})

	p.Put(&ExamplePoolPojo{name2})

	get1 := p.Get()
	fmt.Println(get1.name)
	fmt.Println(p.Get().name)
	p.Put(get1)
	fmt.Println(p.Get().name)

	// Output:
	// matthew
	// matt
	// matthew
}
