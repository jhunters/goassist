package syncx_test

import (
	"fmt"
	"strings"

	"github.com/jhunters/goassist/concurrent/syncx"
)

type MapExamplePojo struct {
	Name string
}

func (mp *MapExamplePojo) CompareTo(v *MapExamplePojo) int {
	return strings.Compare(mp.Name, v.Name)
}

func newMapExamplePojo(name string) *MapExamplePojo {
	return &MapExamplePojo{name}
}

func ExampleNewMapByInitial() {
	mmp := map[string]*MapExamplePojo{
		"key1": newMapExamplePojo("hello"),
		"key2": newMapExamplePojo("world"),
	}

	mp := syncx.NewMapByInitial(mmp)
	mp.Range(func(s string, mep *MapExamplePojo) bool {
		// visit all elements here
		return true
	})

	fmt.Println(mp.Size())
	fmt.Println(mp.Exist("key1"))
	fmt.Println(mp.ExistValue(newMapExamplePojo("world")))

	// Output:
	// 2
	// true
	// key2 true
}

func ExampleNewMap() {
	mp := syncx.NewMap[string, *MapExamplePojo]()
	v := newMapExamplePojo("!")
	mp.Store("hello", v)
	v2, ok := mp.Load("hello")
	fmt.Println(v2.Name, ok)

	// Output:
	// ! true
}
