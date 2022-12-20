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
