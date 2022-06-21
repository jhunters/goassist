package poolx

import "sync"

type PoolX[E any] struct {
	New      func() E
	internal sync.Pool
}

// NewPoolX create a new PoolX
func NewPoolX[E any](f func() E) *PoolX[E] {
	p := PoolX[E]{New: f}
	p.internal = sync.Pool{
		New: func() any {
			return p.New()
		},
	}

	return &p
}

// Get selects an E generic type item from the Pool
func (p *PoolX[E]) Get() E {
	v := p.internal.Get()
	return v.(E)
}

// Put adds x to the pool.
func (p *PoolX[E]) Put(v E) {
	p.internal.Put(v)
}
