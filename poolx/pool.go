package poolx

import "sync"

type Pool[E any] struct {
	New      func() E
	internal sync.Pool
}

// NewPoolX create a new PoolX
func NewPool[E any](f func() E) *Pool[E] {
	p := Pool[E]{New: f}
	p.internal = sync.Pool{
		New: func() any {
			return p.New()
		},
	}

	return &p
}

// Get selects an E generic type item from the Pool
func (p *Pool[E]) Get() E {
	v := p.internal.Get()
	return v.(E)
}

// Put adds x to the pool.
func (p *Pool[E]) Put(v E) {
	p.internal.Put(v)
}
