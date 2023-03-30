package bytex

import (
	"bytes"
	"fmt"

	"github.com/jhunters/goassist/arrayutil"
)

// A ByteBuffer is composite with bytes.Buffer and provides some utility functions for variable-sized buffer of bytes.
type ByteBuffer struct {
	bytes.Buffer
}

// NewByteBuffer return a new ByteBuffer
func NewByteBuffer(bb []byte) *ByteBuffer {
	cp := arrayutil.Clone(bb)
	ret := &ByteBuffer{*bytes.NewBuffer(cp)}
	return ret
}

// NewByteBufferString return a new ByteBuffer
func NewByteBufferString(s string) *ByteBuffer {
	ret := &ByteBuffer{*bytes.NewBuffer([]byte(s))}
	return ret
}

// Delete some bytes between begin and end index. note: this operation will reset offset
func (bbuf *ByteBuffer) Delete(begin, end int) {
	if begin < 0 || begin > bbuf.Len() || begin > end {
		return
	}
	if end > bbuf.Len() {
		bb := bbuf.Bytes()[:begin]
		bbuf.reset(bb)
		return
	}

	bb := bbuf.Bytes()[:begin]
	ret := make([]byte, len(bb))
	copy(ret, bb)
	ret = append(ret, bbuf.Bytes()[end:]...)
	bbuf.reset(ret)
}

// Delete one byte at index offset. note: this operation will reset offset
func (bbuf *ByteBuffer) DeleteIndex(index int) {
	bbuf.Delete(index, index+1)
}

// ReplaceByOffset to replace sub slice by offset index. note: this operation will reset offset
func (bbuf *ByteBuffer) ReplaceByOffset(begin, end int, sub []byte) (err error) {
	if begin > bbuf.Len() || end > bbuf.Len() {
		err = fmt.Errorf("out of index. begin=%d, end=%d", begin, end)
		return
	}
	if begin > end {
		err = fmt.Errorf("begin should lower than end. begin=%d, end=%d", begin, end)
		return
	}

	bb := bbuf.Bytes()
	b1 := bb[:begin]
	ret := make([]byte, len(b1))
	copy(ret, b1)
	ret = append(ret, sub...)
	ret = append(ret, bb[end:]...)

	return bbuf.reset(ret)
}

func (bbuf *ByteBuffer) reset(value []byte) error {
	bbuf.Reset()
	bbuf.Grow(len(value))
	_, err := bbuf.Write(value)
	return err
}

// Index returns the index of the first instance of sub in s, or -1 if sub is not present in s.
func (bbuf *ByteBuffer) Index(sub []byte) int {
	return bytes.Index(bbuf.Bytes(), sub)
}

// Index returns the index of the fromIndex index instance of sub in s, or -1 if sub is not present in s.
func (bbuf *ByteBuffer) IndexOffset(sub []byte, fromIndex int) int {
	if fromIndex < 0 || fromIndex+len(sub) >= bbuf.Len() {
		return -1
	}
	left, _ := bbuf.SubBytes(fromIndex, bbuf.Len())
	index := bytes.Index(left, sub)
	if index != -1 {
		index += fromIndex
	}
	return index
}

// Insert the sub slice into the target offset. note: this operation will reset offset
func (bbuf *ByteBuffer) Insert(offset int, sub []byte) (err error) {
	if offset < 0 || offset > bbuf.Len() {
		return fmt.Errorf("out of index, offset=%d", offset)
	}

	if offset == bbuf.Len() {
		return bbuf.reset(append(bbuf.Bytes(), sub...))
	}

	bb := bbuf.Bytes()
	b1 := bb[:offset]
	ret := make([]byte, len(b1))
	copy(ret, b1)
	ret = append(ret, sub...)
	ret = append(ret, bb[offset:]...)
	return bbuf.reset(ret)
}

// SubString a string that is a substring of this string.
func (bbuf *ByteBuffer) SubBytes(beginIndex, endIndex int) ([]byte, error) {
	if beginIndex < 0 || beginIndex > bbuf.Len() {
		return nil, fmt.Errorf("out of index")
	}
	if endIndex > bbuf.Len() || endIndex < 0 {
		return bbuf.Bytes()[beginIndex:], nil
	}

	return bbuf.Bytes()[beginIndex:endIndex], nil
}

// Reverse the order to slice. note: this operation will reset offset
func (bbuf *ByteBuffer) Reverse() {
	bb := bbuf.Bytes()
	arrayutil.Reverse(bb)
	bbuf.reset(bb)
}
