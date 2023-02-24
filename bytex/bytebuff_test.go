package bytex_test

import (
	"fmt"
	"testing"

	"github.com/jhunters/goassist/bytex"
	. "github.com/smartystreets/goconvey/convey"
)

func TestReplaceByOffset(t *testing.T) {
	Convey("TestReplaceByOffset", t, func() {
		bbuf := bytex.NewByteBuffer([]byte("hello world"))
		err := bbuf.ReplaceByOffset(1, 2, []byte("yes"))
		So(err, ShouldBeNil)
		So(bbuf.Bytes(), ShouldResemble, []byte("hyesllo world"))
	})

	Convey("TestReplaceByOffset outof index", t, func() {
		bbuf := bytex.NewByteBuffer([]byte("hello world"))
		err := bbuf.ReplaceByOffset(2, 1, []byte("yes"))
		So(err, ShouldNotBeNil)

		err = bbuf.ReplaceByOffset(2, 100, []byte("yes"))
		So(err, ShouldNotBeNil)
	})

}

func TestInsert(t *testing.T) {

	Convey("TestInsert start", t, func() {
		bbuf := bytex.NewByteBuffer([]byte("hello world"))
		err := bbuf.Insert(0, []byte("yes"))
		So(err, ShouldBeNil)
		So(bbuf.String(), ShouldEqual, "yeshello world")

	})

	Convey("TestInsert end", t, func() {
		bbuf := bytex.NewByteBuffer([]byte("hello world"))
		err := bbuf.Insert(bbuf.Len(), []byte("yes"))
		So(err, ShouldBeNil)
		So(bbuf.String(), ShouldEqual, "hello worldyes")

	})

	Convey("TestInsert middle position", t, func() {
		bbuf := bytex.NewByteBuffer([]byte("hello world"))
		err := bbuf.Insert(3, []byte("yes"))
		So(err, ShouldBeNil)
		So(bbuf.String(), ShouldEqual, "helyeslo world")

	})

	Convey("TestInsert out of index", t, func() {
		bbuf := bytex.NewByteBuffer([]byte("hello world"))
		err := bbuf.Insert(bbuf.Len()+1, []byte("yes"))
		So(err, ShouldNotBeNil)
	})

}

func TestIndex(t *testing.T) {
	Convey("TestIndex", t, func() {

		bbuf := bytex.NewByteBuffer([]byte("hello world"))
		idx := bbuf.Index([]byte("lo wo"))
		So(idx, ShouldEqual, 3)

		// test not found
		idx = bbuf.Index([]byte("lo wol"))
		So(idx, ShouldEqual, -1)

	})

}

func TestSubBytes(t *testing.T) {
	Convey("TestSubBytes", t, func() {
		origin := []byte("hello world")
		bbuf := bytex.NewByteBuffer(origin)

		// sub all
		bb, err := bbuf.SubBytes(0, bbuf.Len())
		So(err, ShouldBeNil)
		So(bb, ShouldResemble, bbuf.Bytes())

		bb, err = bbuf.SubBytes(0, 50)
		So(err, ShouldBeNil)
		So(bb, ShouldResemble, bbuf.Bytes())

		// sub partial
		bb, err = bbuf.SubBytes(0, 5)
		So(err, ShouldBeNil)
		So(bb, ShouldResemble, origin[:5])

		// out of index
		bb, err = bbuf.SubBytes(25, 50)
		So(err, ShouldNotBeNil)
		So(bb, ShouldBeNil)
	})

}

func TestDelete(t *testing.T) {
	Convey("TestDelete", t, func() {
		origin := []byte("hello world")
		bbuf := bytex.NewByteBuffer(origin)

		// test delete all
		bbuf.Delete(0, bbuf.Len())
		So(bbuf.Bytes(), ShouldBeEmpty)

		bbuf = bytex.NewByteBuffer(origin)
		// test detele start
		bbuf.Delete(0, 5)
		So(bbuf.String(), ShouldEqual, " world")

		// test delete end with
		bbuf = bytex.NewByteBuffer(origin)
		bbuf.Delete(5, 15)
		So(bbuf.String(), ShouldEqual, "hello")

		// test delete out of index
		bbuf = bytex.NewByteBuffer(origin)
		bbuf.Delete(-1, 15)

	})

	Convey("Test DeleteIndex", t, func() {
		origin := []byte("hello world")
		bbuf := bytex.NewByteBuffer(origin)

		bbuf.DeleteIndex(-1)
		So(bbuf.String(), ShouldEqual, "hello world")

		bbuf.DeleteIndex(0)
		So(bbuf.String(), ShouldEqual, "ello world")

		bbuf.DeleteIndex(bbuf.Len() - 1)
		So(bbuf.String(), ShouldEqual, "ello worl")
	})

}

func TestReverse(t *testing.T) {
	Convey("TestReverse", t, func() {
		origin := []byte("hello world")
		bbuf := bytex.NewByteBuffer(origin)

		bbuf.Reverse()

		So(bbuf.Bytes(), ShouldResemble, []byte("dlrow olleh"))
	})

}

func TestCap(t *testing.T) {
	Convey("TestCap", t, func() {
		bb := make([]byte, 0, 100)
		bbuf := bytex.NewByteBuffer(bb)
		So(bbuf.Cap(), ShouldEqual, 100)
		So(bbuf.Len(), ShouldEqual, 0)
		bbuf.WriteString("1234567890")
		So(bbuf.Cap(), ShouldEqual, 100)
		So(bbuf.Len(), ShouldEqual, 10)
		bbuf.Truncate(0)
		So(bbuf.Cap(), ShouldEqual, 100)
		So(bbuf.Len(), ShouldEqual, 0)
	})

}

func TestComplexUsage(t *testing.T) {
	Convey("TestComplexUsage", t, func() {
		origin := []byte("hello world")
		bbuf := bytex.NewByteBuffer(origin)
		b, err := bbuf.ReadByte()
		So(b, ShouldEqual, 'h')
		So(err, ShouldBeNil)

		// insert one and reset offset
		bbuf.Insert(bbuf.Len(), []byte("!"))
		b, err = bbuf.ReadByte()
		So(b, ShouldEqual, 'e')
		So(err, ShouldBeNil)

		bb, err := bbuf.ReadBytes('w')
		So(string(bb), ShouldEqual, "llo w")
		So(err, ShouldBeNil)
		fmt.Println(bb, err)
	})
}

func ExampleNewByteBuffer() {
	origin := []byte("hello world")
	bbuf := bytex.NewByteBuffer(origin)
	err := bbuf.ReplaceByOffset(1, 2, []byte("yes"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bbuf)

	// insert
	err = bbuf.Insert(1, []byte("test"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bbuf)

	// index
	idx := bbuf.Index([]byte("lo wo"))
	fmt.Println(idx)

	// SubBytes
	bb, err := bbuf.SubBytes(2, 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bb)

	bbuf = bytex.NewByteBuffer(origin)
	// test detele start
	bbuf.Delete(0, 6)
	fmt.Println(bbuf)

	bbuf = bytex.NewByteBuffer(origin)
	bbuf.Reverse()
	fmt.Println(bbuf)

	// Output:
	// hyesllo world
	// htestyesllo world
	// 9
	// [101 115 116]
	// world
	// dlrow olleh
}

func ExampleNewByteBufferString() {
	origin := "hello world"
	bbuf := bytex.NewByteBufferString(origin)
	err := bbuf.ReplaceByOffset(1, 2, []byte("yes"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bbuf)

	// insert
	err = bbuf.Insert(1, []byte("test"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bbuf)

	// index
	idx := bbuf.Index([]byte("lo wo"))
	fmt.Println(idx)

	// SubBytes
	bb, err := bbuf.SubBytes(2, 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bb)

	bbuf = bytex.NewByteBufferString(origin)
	// test detele start
	bbuf.Delete(0, 6)
	fmt.Println(bbuf)

	bbuf = bytex.NewByteBufferString(origin)
	bbuf.Reverse()
	fmt.Println(bbuf)

	// Output:
	// hyesllo world
	// htestyesllo world
	// 9
	// [101 115 116]
	// world
	// dlrow olleh
}
