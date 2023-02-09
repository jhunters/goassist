package mapx_test

import (
	"fmt"
	"strings"

	"github.com/jhunters/goassist/container/mapx"
)

type MapPojo struct {
	Name string
}

func (mp *MapPojo) CompareTo(v *MapPojo) int {
	return strings.Compare(mp.Name, v.Name)
}

func newMapPojo(name string) *MapPojo {
	return &MapPojo{name}
}

func createMap() *mapx.Map[string, *MapPojo] {
	mp := mapx.NewMap[string, *MapPojo]()
	mp.Put("key1", newMapPojo("hello"))
	mp.Put("key2", newMapPojo("world"))
	mp.Put("key3", newMapPojo("to"))
	mp.Put("key4", newMapPojo("you"))
	mp.Put("key5", newMapPojo("!"))
	return mp
}

func createMap2() *mapx.Map[string, *MapPojo] {
	mp := mapx.NewMap[string, *MapPojo]()
	mp.Put("key1", newMapPojo("2"))
	mp.Put("key2", newMapPojo("3"))
	mp.Put("key3", newMapPojo("4"))
	mp.Put("key4", newMapPojo("1"))
	mp.Put("key5", newMapPojo("5"))
	return mp
}

func ExampleNewMap() {
	mp := mapx.NewMap[string, *MapPojo]()
	mp.Put("key1", newMapPojo("hello"))
	mp.Put("key2", newMapPojo("world"))
	mp.Put("key3", newMapPojo("to"))
	mp.Put("key4", newMapPojo("you"))
	mp.Put("key5", newMapPojo("!"))

	v, exist := mp.Get("key3")
	fmt.Println(v.Name, exist)

	keys := mp.Keys()
	fmt.Println(len(keys))

	key, v := mp.MaxKey(func(s1, s2 string) int { return strings.Compare(s1, s2) })
	fmt.Println(key, v.Name)

	key, v = mp.MaxValue(func(mp1, mp2 *MapPojo) int { return mp1.CompareTo(mp2) })
	fmt.Println(key, v.Name)

	key, exist = mp.ExistValue(newMapPojo("to"))
	fmt.Println(key, exist)

	// Output:
	// to true
	// 5
	// key5 !
	// key4 you
	// key3 true
}
