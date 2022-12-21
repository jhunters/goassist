package syncx_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jhunters/goassist/concurrent/syncx"
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

func createMap() *syncx.Map[string, *MapPojo] {
	mp := syncx.NewMap[string, *MapPojo]()
	mp.Store("key1", newMapPojo("hello"))
	mp.Store("key2", newMapPojo("world"))
	mp.Store("key3", newMapPojo("to"))
	mp.Store("key4", newMapPojo("you"))
	mp.Store("key5", newMapPojo("!"))
	return mp
}

func TestNewMap(t *testing.T) {
	Convey("TestNewMap", t, func() {
		mp := syncx.NewMap[string, *MapPojo]()
		So(mp, ShouldNotBeNil)
		So(mp.IsEmpty(), ShouldBeTrue)

	})

	Convey("TestNewMapNewMapByInitial", t, func() {

		mp := syncx.NewMapByInitial[string, *MapPojo](nil)
		So(mp, ShouldNotBeNil)
		So(mp.IsEmpty(), ShouldBeTrue)

		mp1 := createMap()
		mp = syncx.NewMapByInitial(mp1.ToMap())
		So(mp, ShouldNotBeNil)
		So(mp.IsEmpty(), ShouldBeFalse)
		So(mp.Size(), ShouldEqual, mp1.Size())

	})
}

func TestMapLoad(t *testing.T) {
	Convey("TestMapLoad", t, func() {
		mp := createMap()
		So(mp.IsEmpty(), ShouldBeFalse)

		Convey("TestMapLoad key", func() {
			v, ok := mp.Load("key1")
			So(ok, ShouldBeTrue)
			So(v, ShouldNotBeNil)
			So(v.Name, ShouldEqual, "hello")
		})

		Convey("TestMapLoad LoadOrStore", func() {
			v, ok := mp.LoadOrStore("key1", newMapPojo("unknown"))
			So(ok, ShouldBeTrue)
			So(v, ShouldNotBeNil)
			So(v.Name, ShouldEqual, "hello")

			v, ok = mp.LoadOrStore("key10", newMapPojo("unknown"))
			So(ok, ShouldBeFalse)
			So(v, ShouldNotBeNil)
			So(v.Name, ShouldEqual, "unknown")
		})

		Convey("TestMapLoad LoadAndDelete", func() {
			v, ok := mp.LoadAndDelete("key1")
			So(ok, ShouldBeTrue)
			So(v, ShouldNotBeNil)

			v, ok = mp.LoadAndDelete("key10")
			So(ok, ShouldBeFalse)
			So(v, ShouldBeNil)

		})
	})
}

func TestMapStore(t *testing.T) {
	Convey("TestMapStore", t, func() {
		Convey("Test store one", func() {
			mp := createMap()
			So(mp.IsEmpty(), ShouldBeFalse)

			v, ok := mp.Load("key1")
			So(ok, ShouldBeTrue)
			So(v, ShouldNotBeNil)
			So(v.Name, ShouldEqual, "hello")

			So(mp.Exist("key1"), ShouldBeTrue)
		})

		Convey("Test store all", func() {
			mp := createMap()
			So(mp.IsEmpty(), ShouldBeFalse)

			mp2 := syncx.NewMap[string, *MapPojo]()
			So(mp2.IsEmpty(), ShouldBeTrue)
			So(mp2.Size(), ShouldBeZeroValue)

			mp2.StoreAll(nil)
			So(mp2.IsEmpty(), ShouldBeTrue)
			So(mp2.Size(), ShouldBeZeroValue)

			mp2.StoreAll(mp)
			So(mp2.Size(), ShouldEqual, mp.Size())

		})

		Convey("Test store all from origin map", func() {

			mp2 := syncx.NewMap[string, *MapPojo]()
			So(mp2.IsEmpty(), ShouldBeTrue)
			So(mp2.Size(), ShouldBeZeroValue)

			mp2.StoreAllOrigin(nil)
			So(mp2.IsEmpty(), ShouldBeTrue)
			So(mp2.Size(), ShouldBeZeroValue)

			mmp := make(map[string]*MapPojo)
			mmp["key1"] = newMapPojo("hello")
			mp2.StoreAllOrigin(mmp)
			So(mp2.Size(), ShouldEqual, len(mmp))
			v, ok := mp2.Load("key1")
			So(ok, ShouldBeTrue)
			So(v, ShouldResemble, mmp["key1"])

		})
	})
}

func TestMapDelete(t *testing.T) {
	Convey("TestMapDelete", t, func() {
		mp := createMap()
		So(mp.IsEmpty(), ShouldBeFalse)

		mp.Delete("key1")
		So(mp.Exist("key1"), ShouldBeFalse)

		mp.Delete("key10")
		So(mp.Exist("key10"), ShouldBeFalse)

	})
}

func TestMapRange(t *testing.T) {
	Convey("TestMapRange", t, func() {
		mp := createMap()
		So(mp.IsEmpty(), ShouldBeFalse)

		Convey("TestMapRange quick break", func() {
			mmp := make(map[string]*MapPojo)
			var foundkey string
			mp.Range(func(key string, value *MapPojo) bool {
				foundkey = key
				mmp[key] = value
				return false
			})
			So(len(mmp), ShouldEqual, 1)
			v, _ := mp.Load(foundkey)
			So(mmp[foundkey], ShouldResemble, v)
		})

		Convey("TestMapRange all", func() {
			mmp := make(map[string]*MapPojo)
			var foundkey string
			mp.Range(func(key string, value *MapPojo) bool {
				foundkey = key
				mmp[key] = value
				return true
			})
			So(len(mmp), ShouldEqual, mp.Size())
			v, _ := mp.Load(foundkey)
			So(mmp[foundkey], ShouldResemble, v)
		})

	})
}

func TestMapRelace(t *testing.T) {
	Convey("TestMapRelace", t, func() {
		mp := createMap()
		So(mp.IsEmpty(), ShouldBeFalse)

		Convey("TestMapRelace exist", func() {
			ok := mp.Replace("key1", newMapPojo("hello"), newMapPojo("unknown"), func(mp1, mp2 *MapPojo) bool {
				return strings.Compare(mp1.Name, mp2.Name) == 0
			})
			So(ok, ShouldBeTrue)

			v, ok := mp.Load("key1")
			So(ok, ShouldBeTrue)
			So(v, ShouldResemble, newMapPojo("unknown"))

		})

		Convey("TestMapRelace not exist", func() {
			ok := mp.Replace("key1", newMapPojo("unknown"), newMapPojo("unknown"), func(mp1, mp2 *MapPojo) bool {
				return strings.Compare(mp1.Name, mp2.Name) == 0
			})
			So(ok, ShouldBeFalse)

			v, ok := mp.Load("key1")
			So(ok, ShouldBeTrue)
			So(v, ShouldResemble, newMapPojo("hello"))

		})

		Convey("ReplaceByCondition exist", func() {
			ok := mp.ReplaceByCondition("key1", func(s string, mp *MapPojo) *MapPojo {
				if strings.Compare(mp.Name, "hello") == 0 {
					return newMapPojo("unknown")
				}
				return mp
			})
			So(ok, ShouldBeTrue)

			v, ok := mp.Load("key1")
			So(ok, ShouldBeTrue)
			So(v, ShouldResemble, newMapPojo("unknown"))

		})
	})
}

func TestMapToMap(t *testing.T) {

	Convey("TestMapToMap", t, func() {
		mp := createMap()
		So(mp.IsEmpty(), ShouldBeFalse)

		mmp := mp.ToMap()
		So(len(mmp), ShouldEqual, mp.Size())
		mp.Range(func(key string, value *MapPojo) bool {
			v, ok := mmp[key]
			So(ok, ShouldBeTrue)
			So(v, ShouldResemble, value)
			return true
		})
	})
}

func TestMapValues(t *testing.T) {

	Convey("TestMapValues", t, func() {
		mp := createMap()
		So(mp.IsEmpty(), ShouldBeFalse)

		values := mp.Values()
		So(len(values), ShouldEqual, mp.Size())
	})
}

func TestMapKeys(t *testing.T) {

	Convey("TestMapKeys", t, func() {
		mp := createMap()
		So(mp.IsEmpty(), ShouldBeFalse)

		keys := mp.Keys()
		So(len(keys), ShouldEqual, mp.Size())
	})
}

func TestMapClear(t *testing.T) {

	Convey("TestMapClear", t, func() {
		mp := createMap()
		So(mp.IsEmpty(), ShouldBeFalse)

		mp.Clear()
		So(mp.IsEmpty(), ShouldBeTrue)
	})
}

func TestMapExistValue(t *testing.T) {
	Convey("TestMapExistValue", t, func() {
		mp := createMap()

		k, ok := mp.ExistValueWithComparator(newMapPojo("!"), func(mp1, mp2 *MapPojo) bool {
			return strings.Compare(mp1.Name, mp2.Name) == 0
		})
		So(ok, ShouldBeTrue)
		So(k, ShouldEqual, "key5")

		k, ok = mp.ExistValue(newMapPojo("!"))
		So(ok, ShouldBeTrue)
		So(k, ShouldEqual, "key5")

	})

	Convey("ExistValueComparable", t, func() {
		mp := createMap()

		k, ok := mp.ExistValueComparable(newMapPojo("!"))
		So(ok, ShouldBeTrue)
		So(k, ShouldEqual, "key5")

	})

}

func TestMapCopy(t *testing.T) {
	Convey("TestMapCopy", t, func() {
		mp := createMap()

		newMp := mp.Copy()
		So(newMp.Size(), ShouldEqual, mp.Size())

	})
}

func TestXxx(t *testing.T) {

	mp := make(map[any]string)

	mp[MapPojo{Name: "hello"}] = "1"
	mp[MapPojo{Name: "world"}] = "2"

	fmt.Println(mp[MapPojo{Name: "world"}])

	mp2 := syncx.NewMap[MapPojo, *MapPojo]()
	mp2.Store(MapPojo{Name: "world"}, &MapPojo{Name: "world"})
	fmt.Println(mp2.Load(MapPojo{Name: "world"}))

}
