package mapx_test

import (
	"strings"
	"testing"

	"github.com/jhunters/goassist/container/mapx"
	. "github.com/smartystreets/goconvey/convey"
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

func TestNewMap(t *testing.T) {
	Convey("TestNewMap", t, func() {
		Convey("empty map", func() {
			mp := mapx.NewMap[string, *MapPojo]()
			So(mp.IsEmpty(), ShouldBeTrue)
			So(mp.Size(), ShouldBeZeroValue)
		})
	})
}

func TestMapPutGet(t *testing.T) {

	Convey("TestPutGet", t, func() {
		mp := createMap()
		v, ok := mp.Get("key1")
		So(ok, ShouldBeTrue)
		So(v, ShouldResemble, newMapPojo("hello"))

		mp.Put("newkey", newMapPojo("xmas"))
		v, ok = mp.Get("newkey")
		So(ok, ShouldBeTrue)
		So(v, ShouldResemble, newMapPojo("xmas"))

		Convey("not exist key", func() {
			v, ok = mp.Get("unknown")
			So(ok, ShouldBeFalse)
			So(v, ShouldBeNil)
		})
	})

}

func TestMapExist(t *testing.T) {
	Convey("TestMapExist", t, func() {
		mp := createMap()
		Convey("test key exist", func() {
			exist := mp.Exist("key3")
			So(exist, ShouldBeTrue)

			exist = mp.Exist("unknown")
			So(exist, ShouldBeFalse)
		})

		Convey("test value exist", func() {
			k, exist := mp.ExistValue(newMapPojo("to"))
			So(exist, ShouldBeTrue)
			So(k, ShouldEqual, "key3")

			k, exist = mp.ExistValue(newMapPojo("unknown"))
			So(exist, ShouldBeFalse)
			So(k, ShouldBeZeroValue)
		})

		Convey("test value exist with comparator", func() {
			k, exist := mp.ExistValueWithComparator(newMapPojo("to"), func(mp1, mp2 *MapPojo) bool {
				return strings.Compare(mp1.Name, mp2.Name) == 0
			})
			So(exist, ShouldBeTrue)
			So(k, ShouldEqual, "key3")

			k, exist = mp.ExistValueWithComparator(newMapPojo("unknown"), func(mp1, mp2 *MapPojo) bool {
				return strings.Compare(mp1.Name, mp2.Name) == 0
			})
			So(exist, ShouldBeFalse)
			So(k, ShouldBeZeroValue)
		})

	})

}

func TestMapRemoveAndClear(t *testing.T) {
	Convey("TestRemove", t, func() {
		mp := createMap()
		// key not found
		ok := mp.Remove("unknown")
		So(ok, ShouldBeFalse)

		ok = mp.Remove("key4")
		So(ok, ShouldBeTrue)

		v, ok := mp.Get("key4")
		So(ok, ShouldBeFalse)
		So(v, ShouldBeNil)
	})

	Convey("TestClear", t, func() {
		mp := createMap()
		So(mp.IsEmpty(), ShouldBeFalse)

		mp.Clear()
		So(mp.IsEmpty(), ShouldBeTrue)
		So(mp.Size(), ShouldBeZeroValue)
	})

}

func TestMapKeysAndValues(t *testing.T) {
	Convey("TestKeysAndValues", t, func() {
		Convey("test keys", func() {
			mp := createMap()
			keys := mp.Keys()
			So(len(keys), ShouldEqual, 5)

		})
		Convey("test values", func() {
			mp := createMap()
			values := mp.Values()
			So(len(values), ShouldEqual, 5)
		})
	})
}

func TestMapCopy(t *testing.T) {

	Convey("TestMapCopy", t, func() {
		mp := createMap()
		mp2 := mp.Copy()
		So(mp.Size(), ShouldEqual, mp2.Size())

		mp.Range(func(s string, mp *MapPojo) bool {
			ok := mp2.Exist(s)
			So(ok, ShouldBeTrue)
			return true
		})
	})
}

func TestMapMinMaxValue(t *testing.T) {
	Convey("TestMapMinMaxValue", t, func() {
		mp := createMap2()
		Convey("Test map min value", func() {
			k, v := mp.MinValue(func(mp1, mp2 *MapPojo) int {
				return strings.Compare(mp1.Name, mp2.Name)
			})
			So(k, ShouldEqual, "key4")
			So(v.Name, ShouldEqual, "1")
		})
		Convey("Test map max value", func() {
			k, v := mp.MaxValue(func(mp1, mp2 *MapPojo) int {
				return strings.Compare(mp1.Name, mp2.Name)
			})
			So(k, ShouldEqual, "key5")
			So(v.Name, ShouldEqual, "5")
		})

	})

}
func TestMapMinMaxKey(t *testing.T) {
	Convey("TestMapMinMaxKey", t, func() {
		mp := createMap2()
		Convey("Test map min key", func() {
			k, v := mp.MinKey(func(mp1, mp2 string) int {
				return strings.Compare(mp1, mp2)
			})
			So(k, ShouldEqual, "key1")
			So(v.Name, ShouldEqual, "2")
		})
		Convey("Test map max key", func() {
			k, v := mp.MaxKey(func(mp1, mp2 string) int {
				return strings.Compare(mp1, mp2)
			})
			So(k, ShouldEqual, "key5")
			So(v.Name, ShouldEqual, "5")
		})

	})

}
