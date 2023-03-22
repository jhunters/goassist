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

	exist = mp.Remove("key3")
	fmt.Println(exist)

	// Output:
	// to true
	// 5
	// key5 !
	// key4 you
	// key3 true
	// true
}

func ExampleMap_IsEmpty() {
	mp := mapx.NewMap[string, string]()
	fmt.Println(mp.IsEmpty())

	// Output:
	// true
}

func ExampleMap_Equals() {
	mp1 := mapx.NewMap[string, string]()
	mp1.Put("key1", "value1")
	mp2 := mapx.NewMap[string, string]()
	mp2.Put("key1", "value1")

	fmt.Println(mp1.Equals(mp2, func(s1, s2 string) bool { return s1 == s2 }))
}

func ExampleMap_MaxKey() {
	mp := mapx.NewMap[string, string]()
	mp.Put("key1", ("2"))
	mp.Put("key2", ("3"))
	mp.Put("key3", ("4"))
	mp.Put("key4", ("1"))
	mp.Put("key5", ("5"))

	k, v := mp.MaxKey(func(s1, s2 string) int { return strings.Compare(s1, s2) })
	fmt.Println(k, v)

	// Output:
	// key5 5
}

func ExampleMap_MinKey() {
	mp := mapx.NewMap[string, string]()
	mp.Put("key1", ("2"))
	mp.Put("key2", ("3"))
	mp.Put("key3", ("4"))
	mp.Put("key4", ("1"))
	mp.Put("key5", ("5"))

	k, v := mp.MinKey(func(s1, s2 string) int { return strings.Compare(s1, s2) })
	fmt.Println(k, v)

	// Output:
	// key1 2
}

func ExampleMap_MinValue() {
	mp := mapx.NewMap[string, string]()
	mp.Put("key1", ("2"))
	mp.Put("key2", ("3"))
	mp.Put("key3", ("4"))
	mp.Put("key4", ("1"))
	mp.Put("key5", ("5"))

	k, v := mp.MinValue(func(s1, s2 string) int { return strings.Compare(s1, s2) })
	fmt.Println(k, v)

	// Output:
	// key4 1
}

func ExampleMap_MaxValue() {
	mp := mapx.NewMap[string, string]()
	mp.Put("key1", ("2"))
	mp.Put("key2", ("3"))
	mp.Put("key3", ("4"))
	mp.Put("key4", ("1"))
	mp.Put("key5", ("5"))

	k, v := mp.MaxValue(func(s1, s2 string) int { return strings.Compare(s1, s2) })
	fmt.Println(k, v)

	// Output:
	// key5 5
}

func ExampleMap_Exist() {
	mp := mapx.NewMap[string, string]()
	mp.Put("key1", ("2"))
	mp.Put("key2", ("3"))
	mp.Put("key3", ("4"))
	mp.Put("key4", ("1"))
	mp.Put("key5", ("5"))

	b := mp.Exist("key2")
	fmt.Println(b)

	// Output:
	// true
}

func ExampleMap_ExistValue() {
	mp := mapx.NewMap[string, string]()
	mp.Put("key1", ("2"))
	mp.Put("key2", ("3"))
	mp.Put("key3", ("4"))
	mp.Put("key4", ("1"))
	mp.Put("key5", ("5"))

	k, exist := mp.ExistValue("5")
	fmt.Println(k, exist)

	k, exist = mp.ExistValue("15")
	fmt.Println(k, exist)

	// Output:
	// key5 true
	//  false
}

func ExampleMap_ExistValueWithComparator() {
	mp := mapx.NewMap[string, string]()
	mp.Put("key1", ("2"))
	mp.Put("key2", ("3"))
	mp.Put("key3", ("4"))
	mp.Put("key4", ("1"))
	mp.Put("key5", ("5"))

	k, exist := mp.ExistValueWithComparator("5", func(s1, s2 string) bool { return s1 == s2 })
	fmt.Println(k, exist)

	k, exist = mp.ExistValueWithComparator("15", func(s1, s2 string) bool { return s1 == s2 })
	fmt.Println(k, exist)

	// Output:
	// key5 true
	//  false
}

func ExampleMap_Keys() {
	mp := mapx.NewMap[string, string]()
	mp.Put("key1", ("2"))
	mp.Put("key2", ("3"))
	mp.Put("key3", ("4"))
	mp.Put("key4", ("1"))
	mp.Put("key5", ("5"))

	keys := mp.Keys()
	fmt.Println(len(keys))

	// Output:
	// 5
}

func ExampleMap_Values() {
	mp := mapx.NewMap[string, string]()
	mp.Put("key1", ("2"))
	mp.Put("key2", ("3"))
	mp.Put("key3", ("4"))
	mp.Put("key4", ("1"))
	mp.Put("key5", ("5"))

	vals := mp.Values()
	fmt.Println(len(vals))

	// Output:
	// 5
}
