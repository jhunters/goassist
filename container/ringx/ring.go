package ringx

import (
	"sort"

	"github.com/jhunters/goassist/base"
)

type sortableRing[E any] struct {
	data *Ring[E]
	cmp  base.CMP[E]
}

func (s sortableRing[E]) Len() int { return s.data.Len() }
func (s sortableRing[E]) Swap(i, j int) {
	e1 := s.data.get(i)
	e2 := s.data.get(j)
	e1.Value, e2.Value = e2.Value, e1.Value
}
func (s sortableRing[E]) Less(i, j int) bool {
	e1 := s.data.get(i)
	e2 := s.data.get(j)
	return s.cmp(e1.Value, e2.Value) <= 0
}

// A Ring is an element of a circular list, or ring.
// Rings do not have a beginning or end; a pointer to any ring element
// serves as reference to the entire ring. Empty rings are represented
// as nil Ring pointers. The zero value for a Ring is a one-element
// ring with a nil Value.
//
type Ring[E any] struct {
	next, prev *Ring[E]
	Value      E // for use by client; untouched by this library
}

func (r *Ring[E]) init() *Ring[E] {
	r.next = r
	r.prev = r
	return r
}

// Next returns the next ring element. r must not be empty.
func (r *Ring[E]) Next() *Ring[E] {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// Prev returns the previous ring element. r must not be empty.
func (r *Ring[E]) Prev() *Ring[E] {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

// Move moves n % r.Len() elements backward (n < 0) or forward (n >= 0)
// in the ring and returns that ring element. r must not be empty.
//
func (r *Ring[E]) Move(n int) *Ring[E] {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// Get to get n % r.Len() elements backward (n < 0) or forward (n >= 0)
// in the ring and returns that ring element value. r must not be empty.
func (r *Ring[E]) Get(n int) (ret E) {
	return r.get(n).Value
}

func (r *Ring[E]) get(n int) (ret *Ring[E]) {

	if r.next == nil {
		return r.init()
	}
	ret = r
	switch {
	case n < 0:
		for ; n < 0; n++ {
			ret = ret.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			ret = ret.next
		}
	}
	return ret
}

// New creates a ring of n elements.
func NewRing[E any](n int) *Ring[E] {
	if n <= 0 {
		return nil
	}
	r := new(Ring[E])
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring[E]{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

// New creates a ring of n elements.
func NewRingOf[E any](e ...E) *Ring[E] {
	if e == nil || len(e) == 0 {
		return nil
	}
	r := &Ring[E]{Value: e[0]}
	p := r

	for i := 1; i < len(e); i++ {
		p.next = &Ring[E]{prev: p, Value: e[i]}
		p = p.next
	}

	p.next = r
	r.prev = p
	return r
}

// Link connects ring r with ring s such that r.Next()
// becomes s and returns the original value for r.Next().
// r must not be empty.
//
// If r and s point to the same ring, linking
// them removes the elements between r and s from the ring.
// The removed elements form a subring and the result is a
// reference to that subring (if no elements were removed,
// the result is still the original value for r.Next(),
// and not nil).
//
// If r and s point to different rings, linking
// them creates a single ring with the elements of s inserted
// after r. The result points to the element following the
// last element of s after insertion.
//
func (r *Ring[E]) Link(s *Ring[E]) *Ring[E] {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		// Note: Cannot use multiple assignment because
		// evaluation order of LHS is not specified.
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

func (r *Ring[E]) LinkValue(e E) {
	nr := &Ring[E]{Value: e}
	r.Link(nr)
}

// Unlink removes n % r.Len() elements from the ring r, starting
// at r.Next(). If n % r.Len() == 0, r remains unchanged.
// The result is the removed subring. r must not be empty.
//
func (r *Ring[E]) Unlink(n int) *Ring[E] {
	if n <= 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}

// Len computes the number of elements in ring r.
// It executes in time proportional to the number of elements.
//
func (r *Ring[E]) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}

// Do calls function f on each element of the ring, in forward order.
// The behavior of Do is undefined if f changes *r.
func (r *Ring[E]) Do(f func(E)) {
	if r != nil {
		f(r.Value)
		for p := r.Next(); p != r; p = p.next {
			f(p.Value)
		}
	}
}

func (r *Ring[E]) Iterate(f func(E) bool) {
	if r != nil {
		f(r.Value)
		for p := r.Next(); p != r; p = p.next {
			if !f(p.Value) {
				return
			}
		}
	}
}

// Min to find the minimum one in the list
func (r *Ring[E]) Min(compare base.CMP[E]) (min E) {
	return selectByCompareRing(r, func(o1, o2 E) int {
		return compare(o1, o2)
	})
}

// Max to find the maximum one in the list
func (r *Ring[E]) Max(compare base.CMP[E]) (min E) {
	return selectByCompareRing(r, func(o1, o2 E) int {
		return compare(o2, o1)
	})
}

func selectByCompareRing[E any](r *Ring[E], compare base.CMP[E]) (v E) {
	i := 0
	r.Do(func(e E) {
		if i == 0 {
			v = e
		} else {
			if compare(v, e) > 0 {
				v = e
			}
		}
		i++
	})
	return
}

// Sort to sort ring elements order by compare condition.
func (r *Ring[E]) Sort(compare base.CMP[E]) {
	sortobject := sortableRing[E]{data: r, cmp: compare}
	sort.Sort(sortobject)
}

// Index return the index of the first matched object in list
func (r *Ring[E]) Index(v E, f base.EQL[E]) (index int) {
	index = -1
	matched := false
	r.Iterate(func(e E) bool {
		index++
		if f(v, e) {
			matched = true
			return false
		}
		return true
	})
	if !matched {
		index = -1
	}
	return
}

// Sort to sort ring elements order by compare condition.
func (r *Ring[E]) Contains(v E, f base.EQL[E]) bool {
	return r.Index(v, f) != -1
}

// ToArray to convert list elements to array
func (r *Ring[E]) ToArray() []E {
	ret := make([]E, 0)
	r.Do(func(e E) {
		ret = append(ret, e)
	})

	return ret
}

// WriteToArray extract list element to array
func (r *Ring[E]) WriteToArray(v []E) {
	size := len(v)
	pos := 0

	r.Iterate(func(e E) bool {
		if pos < size {
			v[pos] = e
		} else {
			return false
		}
		pos++
		return true
	})

}

// Copy to copy all elements to a new ring
func (r *Ring[E]) Copy() *Ring[E] {
	ret := NewRing[E](1)
	i := 0
	r.Do(func(e E) {
		if i == 0 {
			ret.Value = e
		} else {
			ret.LinkValue(e)
		}
		i++
	})
	return ret
}
