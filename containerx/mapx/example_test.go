package mapx_test

import (
	"fmt"
	"strings"

	"github.com/jhunters/goassist/containerx/mapx"
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

func ExampleExistValueComparable() {
	v := newMapExamplePojo("!")
	mp := mapx.NewMap[string, *MapExamplePojo]()
	mp.Store("hello", v)

	k, ok := mp.ExistValueComparable(v)
	fmt.Println(k, ok)

	// Output:
	// hello true
}
