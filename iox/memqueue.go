package iox

import (
	"encoding/binary"
	"io"
)

type Queue interface {
	Enqueue(data []byte)

	Dequeue() ([]byte, error)

	Close()
}

type MemQueue struct {
	in  *io.PipeReader
	out *io.PipeWriter
}

func NewMemQueue() Queue {
	in, out := io.Pipe()
	mq := MemQueue{in, out}
	return &mq
}

func (mq *MemQueue) Enqueue(data []byte) {
	l := len(data)
	offset := make([]byte, 8)
	binary.BigEndian.PutUint64(offset, uint64(l))

	mq.out.Write(offset)
	mq.out.Write(data)
}

func (mq *MemQueue) Dequeue() ([]byte, error) {
	offset := make([]byte, 8)
	_, err := io.ReadFull(mq.in, offset)
	if err != nil {
		return nil, err
	}

	l := binary.BigEndian.Uint64(offset)
	data := make([]byte, l)
	_, err = io.ReadFull(mq.in, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (mq *MemQueue) Close() {
	mq.in.Close()
	mq.out.Close()
}
