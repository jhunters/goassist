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
	visitedCount := 0
	mp.Range(func(s string, mep *MapExamplePojo) bool {
		// visit all elements here
		visitedCount++
		return true
	})
	fmt.Println(visitedCount)

	fmt.Println(mp.Size())
	fmt.Println(mp.Exist("key1"))
	fmt.Println(mp.ExistValue(newMapExamplePojo("world")))

	// LoadOrStore
	v, loaded := mp.LoadOrStore("key1", newMapExamplePojo("value1"))
	fmt.Println(v.Name, loaded)
	v, loaded = mp.LoadOrStore("key3", newMapExamplePojo("value3"))
	fmt.Println(v.Name, loaded)

	// LoadAndDelete
	v, loaded = mp.LoadAndDelete("key3")
	fmt.Println(v.Name, loaded)

	// Output:
	// 2
	// 2
	// true
	// key2 true
	// hello true
	// value3 false
	// value3 true
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

func ExampleMap_Store() {
	mmp := map[string]string{
		"key1": ("hello"),
		"key2": ("world"),
	}

	mp := syncx.NewMapByInitial(mmp)
	mp.Store("key3", "!")

	v, exist := mp.Load("key3")
	fmt.Println(v, exist)

	// Output:
	// ! true
}

func ExampleMap_StoreAllOrigin() {
	mmp := map[string]string{
		"key1": ("hello"),
		"key2": ("world"),
	}

	mp := syncx.NewMap[string, string]()
	mp.StoreAllOrigin(mmp)

	v, exist := mp.Load("key1")
	fmt.Println(v, exist)

	// Output:
	// hello true
}

func ExampleMap_StoreAll() {
	mmp := map[string]string{
		"key1": ("hello"),
		"key2": ("world"),
	}

	mp := syncx.NewMap[string, string]()

	mp1 := syncx.NewMapByInitial(mmp)
	mp.StoreAll(mp1)

	v, exist := mp.Load("key1")
	fmt.Println(v, exist)

	// Output:
	// hello true
}

func ExampleMap_Replace() {
	mmp := map[string]string{
		"key1": ("hello"),
		"key2": ("world"),
	}
	mp := syncx.NewMapByInitial(mmp)

	success := mp.Replace("key1", "hello", "welcome", func(s1, s2 string) bool { return s1 == s2 })
	v, exist := mp.Load("key1")
	fmt.Println(v, exist)
	fmt.Println(success)

	// Output:
	// welcome true
	// true
}

func ExampleMap_ReplaceByCondition() {
	mmp := map[string]string{
		"key1": ("hello"),
		"key2": ("world"),
	}
	mp := syncx.NewMapByInitial(mmp)

	success := mp.ReplaceByCondition("key1", func(key, oldvalue string) string {
		if oldvalue == "hello" {
			return "welcome"
		}
		return oldvalue
	})
	v, exist := mp.Load("key1")
	fmt.Println(v, exist)
	fmt.Println(success)

	// Output:
	// welcome true
	// true
}

func ExampleMap_ToMap() {
	mmp := map[string]string{
		"key1": ("hello"),
		"key2": ("world"),
	}
	mp := syncx.NewMapByInitial(mmp)

	mmp1 := mp.ToMap()
	fmt.Println(len(mmp1))

	// Output:
	// 2
}

func ExampleMap_Load() {
	mmp := map[string]string{
		"key1": ("hello"),
		"key2": ("world"),
	}
	mp := syncx.NewMapByInitial(mmp)

	v, exist := mp.Load("key1")
	fmt.Println(v, exist)

	v, exist = mp.Load("key12")
	fmt.Println(v, exist)

	// Output:
	// hello true
	//  false
}

func ExampleMap_LoadOrStore() {
	mmp := map[string]string{
		"key1": ("hello"),
		"key2": ("world"),
	}
	mp := syncx.NewMapByInitial(mmp)

	v, loaded := mp.LoadOrStore("key1", "welcome")
	fmt.Println(v, loaded)

	v, loaded = mp.LoadOrStore("key3", "welcome")
	fmt.Println(v, loaded)

	// Output:
	// hello true
	// welcome false
}

func ExampleMap_LoadAndDelete() {
	mmp := map[string]string{
		"key1": ("hello"),
		"key2": ("world"),
	}
	mp := syncx.NewMapByInitial(mmp)

	v, loaded := mp.LoadAndDelete("key1")
	fmt.Println(v, loaded)

	v, loaded = mp.LoadAndDelete("key1")
	fmt.Println(v, loaded)

	// Output:
	// hello true
	//  false
}

func ExampleMap_MaxKey() {
	mmp := map[string]string{
		"key1": ("hello"),
		"key2": ("world"),
	}
	mp := syncx.NewMapByInitial(mmp)

	k, v := mp.MaxKey(func(s1, s2 string) int { return strings.Compare(s1, s2) })
	fmt.Println(k, v)

	// Output:
	// key2 world
}

func ExampleMap_MinKey() {
	mmp := map[string]string{
		"key1": ("hello"),
		"key2": ("world"),
	}
	mp := syncx.NewMapByInitial(mmp)

	k, v := mp.MinKey(func(s1, s2 string) int { return strings.Compare(s1, s2) })
	fmt.Println(k, v)

	// Output:
	// key1 hello
}

func ExampleMap_MinValue() {
	mmp := map[string]string{
		"key1": ("hello"),
		"key2": ("world"),
	}
	mp := syncx.NewMapByInitial(mmp)

	k, v := mp.MinValue(func(s1, s2 string) int { return strings.Compare(s1, s2) })
	fmt.Println(k, v)

	// Output:
	// key1 hello
}

func ExampleMap_MaxValue() {
	mmp := map[string]string{
		"key1": ("hello"),
		"key2": ("world"),
	}
	mp := syncx.NewMapByInitial(mmp)

	k, v := mp.MaxValue(func(s1, s2 string) int { return strings.Compare(s1, s2) })
	fmt.Println(k, v)

	// Output:
	// key2 world
}

func ExampleMap_Equals() {
	mmp := map[string]string{
		"key1": ("hello"),
		"key2": ("world"),
	}
	mp := syncx.NewMapByInitial(mmp)

	mp2 := mp.Copy()

	equal := mp.Equals(mp2, func(s1, s2 string) bool { return s1 == s2 })
	fmt.Println(equal)

	// Output:
	// true
}

func ExampleMap_Range() {
	mmp := map[string]string{
		"key1": ("hello"),
		"key2": ("world"),
	}
	mp := syncx.NewMapByInitial(mmp)

	count := 0
	mp.Range(func(s1, s2 string) bool {
		count++
		return true
	})
	fmt.Println(count)

	// Output:
	// 2
}
