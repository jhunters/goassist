package compress_test

import (
	"testing"

	"github.com/jhunters/goassist/compress"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGZIP(t *testing.T) {

	Convey("TestGZIP", t, func() {
		ori := []byte("hello world!")
		compressed, err := compress.GZIP(ori)
		So(err, ShouldBeNil)
		decompressed, err := compress.GUNZIP(compressed)
		So(err, ShouldBeNil)
		So(ori, ShouldResemble, decompressed)

	})
}
