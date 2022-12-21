package syncx_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/jhunters/goassist/concurrent/syncx"
)

type ExamplePoolPojo struct {
	name string
}

func ExamplePool() {
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

func TestCond(t *testing.T) {
	var m sync.Mutex
	c := sync.NewCond(&m)

	for i := 0; i < 10; i++ {
		c.L.Lock()
		go func(c *sync.Cond, no int) {
			c.Wait()
			fmt.Println("no=", no, "finished")
			c.L.Unlock()
		}(c, i)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		c.Signal()
	}

}

func TestWait(t *testing.T) {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.example.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			fmt.Println(url)
			fmt.Println("done==>", url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
