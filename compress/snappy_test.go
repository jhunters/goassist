package compress_test

import (
	"testing"

	"github.com/jhunters/goassist/compress"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSnappy(t *testing.T) {
	Convey("TestSnappy", t, func() {

		v := compress.SnappyEncode([]byte("hello world"))

		s, err := compress.SnappyDecode(v)
		So(err, ShouldBeNil)
		So(string(s), ShouldEqual, "hello world")
	})
}
