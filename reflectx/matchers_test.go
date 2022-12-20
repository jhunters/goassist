package reflectx_test

import (
	"testing"

	"github.com/jhunters/goassist/reflectx"
	. "github.com/smartystreets/goconvey/convey"
)

type EqualsPojo struct {
	Name    string
	Age     int
	Address string
	gender  bool
}

func NewEqualsPojo(name string, age int, address string, gender bool) EqualsPojo {
	return EqualsPojo{name, age, address, gender}
}

func TestNewDeepEquals(t *testing.T) {
	Convey("TestNewDeepEquals", t, func() {

		v1 := NewEqualsPojo("matt", 10, "shanghai pudong", true)
		v2 := NewEqualsPojo("matt", 10, "shanghai pudong", true)

		de := reflectx.NewDeepEquals(v1)
		equals := de.Matches(v2)
		So(equals, ShouldBeTrue)

		v2 = NewEqualsPojo("matt", 10, "shanghai pudong", false)
		equals = de.Matches(v2)
		So(equals, ShouldBeFalse)
	})
}
