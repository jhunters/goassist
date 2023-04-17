package compress_test

import (
	"testing"

	"github.com/jhunters/goassist/compress"
	. "github.com/smartystreets/goconvey/convey"
)

func TestToBytesWithCompress(t *testing.T) {
	Convey("TestToBytesWithCompress", t, func() {

		b := compress.CompressUint64(1)
		So(b, ShouldResemble, []byte{1})
		v, n := compress.DecompressUint6(b)
		So(v, ShouldEqual, 1)
		So(n, ShouldEqual, 1)

		b = compress.CompressUint64(100)
		So(b, ShouldResemble, []byte{100})
		v, n = compress.DecompressUint6(b)
		So(v, ShouldEqual, 100)
		So(n, ShouldEqual, 1)

		b = compress.CompressUint64(250)
		So(b, ShouldResemble, []byte{250, 1})
		v, n = compress.DecompressUint6(b)
		So(v, ShouldEqual, 250)
		So(n, ShouldEqual, 2)

		b = compress.CompressUint64(1 << 14)
		So(b, ShouldResemble, []byte{128, 128, 1})
		v, n = compress.DecompressUint6(b)
		So(v, ShouldEqual, 1<<14)
		So(n, ShouldEqual, 3)

		b = compress.CompressUint64(1 << 21)
		So(b, ShouldResemble, []byte{128, 128, 128, 1})
		v, n = compress.DecompressUint6(b)
		So(v, ShouldEqual, 1<<21)
		So(n, ShouldEqual, 4)

		b = compress.CompressUint64(1 << 21)
		So(b, ShouldResemble, []byte{128, 128, 128, 1})
		v, n = compress.DecompressUint6(b)
		So(v, ShouldEqual, 1<<21)
		So(n, ShouldEqual, 4)

		b = compress.CompressUint64(1 << 28)
		So(b, ShouldResemble, []byte{128, 128, 128, 128, 1})
		v, n = compress.DecompressUint6(b)
		So(v, ShouldEqual, 1<<28)
		So(n, ShouldEqual, 5)

		b = compress.CompressUint64(1 << 35)
		So(b, ShouldResemble, []byte{128, 128, 128, 128, 128, 1})
		v, n = compress.DecompressUint6(b)
		So(v, ShouldEqual, 1<<35)
		So(n, ShouldEqual, 6)

		b = compress.CompressUint64(1 << 42)
		So(b, ShouldResemble, []byte{128, 128, 128, 128, 128, 128, 1})
		v, n = compress.DecompressUint6(b)
		So(v, ShouldEqual, 1<<42)
		So(n, ShouldEqual, 7)

		b = compress.CompressUint64(1 << 49)
		So(b, ShouldResemble, []byte{128, 128, 128, 128, 128, 128, 128, 1})
		v, n = compress.DecompressUint6(b)
		So(v, ShouldEqual, 1<<49)
		So(n, ShouldEqual, 8)

		b = compress.CompressUint64(1 << 56)
		So(b, ShouldResemble, []byte{128, 128, 128, 128, 128, 128, 128, 128, 1})
		v, n = compress.DecompressUint6(b)
		So(v, ShouldEqual, 1<<56)
		So(n, ShouldEqual, 9)

		b = compress.CompressUint64(1 << 63)
		So(b, ShouldResemble, []byte{128, 128, 128, 128, 128, 128, 128, 128, 128, 1})
		v, n = compress.DecompressUint6(b)
		So(v, ShouldEqual, uint(1<<63))
		So(n, ShouldEqual, 10)

	})

}
