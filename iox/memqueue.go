package iox

import (
	"encoding/binary"
	"io"
	"sync"
)

type Queue interface {
	Enqueue(data []byte)

	Dequeue() ([]byte, error)

	Close()
}

type MemQueue struct {
	in  *io.PipeReader
	out *io.PipeWriter

	rwlocker sync.RWMutex
}

func NewMemQueue() Queue {
	in, out := io.Pipe()
	mq := MemQueue{in: in, out: out}
	return &mq
}

func (mq *MemQueue) Enqueue(data []byte) {
	mq.rwlocker.Lock()
	defer mq.rwlocker.Unlock()
	l := len(data)
	offset := make([]byte, 8)
	binary.BigEndian.PutUint64(offset, uint64(l))

	mq.out.Write(offset)
	mq.out.Write(data)
}

func (mq *MemQueue) Dequeue() ([]byte, error) {
	mq.rwlocker.RLock()
	defer mq.rwlocker.RUnlock()
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
	if mq.in != nil {
		mq.in.Close()
	}
	if mq.out != nil {
		mq.out.Close()
	}
}
