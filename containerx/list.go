package containerx

import (
	"sort"

	"github.com/jhunters/goassist/base"
)

type sortableList[E any] struct {
	data *List[E]
	cmp  base.CMP[E]
}

func (s sortableList[E]) Len() int { return s.data.Len() }
func (s sortableList[E]) Swap(i, j int) {
	e1, _ := s.data.get(i)
	e2, _ := s.data.get(j)
	e1.Value, e2.Value = e2.Value, e1.Value
}
func (s sortableList[E]) Less(i, j int) bool {
	e1, _ := s.data.Get(i)
	e2, _ := s.data.Get(j)
	return s.cmp(e1, e2) <= 0
}

// Element is an element of a linked list.
type Element[E any] struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element[E]

	// The list to which this element belongs.
	list *List[E]

	// The value stored with this element.
	Value E
}

// Next returns the next list element or nil.
func (e *Element[E]) Next() *Element[E] {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Prev returns the previous list element or nil.
func (e *Element[E]) Prev() *Element[E] {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type List[E any] struct {
	root Element[E] // sentinel list element, only &root, root.prev, and root.next are used
	len  int        // current list length excluding (this) sentinel element
}

// Init initializes or clears list l.
func (l *List[E]) Init() *List[E] {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// New returns an initialized list.
func NewList[E any]() *List[E] { return new(List[E]).Init() }

// NewFromArray returns an initialized list and set elements from target array.
func NewListFromArray[E any](arr []E) *List[E] {
	l := new(List[E]).Init()
	if arr != nil {
		for _, v := range arr {
			l.PushBack(v)
		}
	}

	return l
}

// NewOf returns an initialized list and set elements from target variable parameter.
func NewListOf[E any](e ...E) *List[E] {
	l := new(List[E]).Init()
	if e != nil {
		for _, v := range e {
			l.PushBack(v)
		}
	}
	return l
}

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *List[E]) Len() int { return l.len }

// IsEmpty returns true if len is zero
func (l *List[E]) IsEmpty() bool { return l.len == 0 }

// Front returns the first element of list l or nil if the list is empty.
func (l *List[E]) Front() *Element[E] {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// FrontValue returns the first element of list value
func (l *List[E]) FrontValue() (v E) {
	if e := l.Front(); e != nil {
		v = e.Value
	}
	return
}

// Back returns the last element of list l or nil if the list is empty.
func (l *List[E]) Back() *Element[E] {
	if l.IsEmpty() {
		return nil
	}
	return l.root.prev
}

// Back returns the last element of list l or nil if the list is empty.
func (l *List[E]) BackValue() (v E) {
	if e := l.Back(); e != nil {
		v = e.Value
	}
	return
}

// lazyInit lazily initializes a zero List value.
func (l *List[E]) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert inserts e after at, increments l.len, and returns e.
func (l *List[E]) insert(e, at *Element[E]) *Element[E] {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *List[E]) insertValue(v E, at *Element[E]) *Element[E] {
	return l.insert(&Element[E]{Value: v}, at)
}

// remove removes e from its list, decrements l.len
func (l *List[E]) remove(e *Element[E]) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
}

// move moves e to next to at.
func (l *List[E]) move(e, at *Element[E]) {
	if e == at {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (l *List[E]) RemoveElement(e *Element[E]) E {
	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.remove(e)
	}
	return e.Value
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *List[E]) PushFront(v E) *Element[E] {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *List[E]) PushBack(v E) *Element[E] {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List[E]) InsertBefore(v E, mark *Element[E]) *Element[E] {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark.prev)
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List[E]) InsertAfter(v E, mark *Element[E]) *Element[E] {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark)
}

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List[E]) MoveToFront(e *Element[E]) {
	if e.list != l || l.root.next == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, &l.root)
}

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List[E]) MoveToBack(e *Element[E]) {
	if e.list != l || l.root.prev == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, l.root.prev)
}

// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List[E]) MoveBefore(e, mark *Element[E]) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark.prev)
}

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List[E]) MoveAfter(e, mark *Element[E]) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark)
}

// PushBackList inserts a copy of another list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List[E]) PushBackList(other *List[E]) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// PushFrontList inserts a copy of another list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List[E]) PushFrontList(other *List[E]) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}

// ToArray to convert list elements to array
func (l *List[E]) ToArray() []E {
	if l.IsEmpty() {
		return []E{}
	}
	ret := make([]E, l.len)
	e := l.Front()
	i := 0
	for e != nil {
		ret[i] = e.Value
		e = e.Next()
		i++
	}
	return ret
}

// WriteToArray extract list element to array
func (l *List[E]) WriteToArray(v []E) {
	if l.IsEmpty() || v == nil || len(v) == 0 {
		return
	}

	size := len(v)
	pos := 0
	l.iterate(func(e *Element[E]) bool {
		if pos < size {
			v[pos] = e.Value
		} else {
			return false
		}
		pos++
		return true
	})
}

// Iterator to iterate all elements
func (l *List[E]) Iterate(f func(E) bool) {
	if l.IsEmpty() {
		return
	}
	l.iterate(func(e *Element[E]) bool {
		return f(e.Value)
	})
}

// Iterator to iterate all elements
func (l *List[E]) IterateReverse(f func(E) bool) {
	if l.IsEmpty() {
		return
	}
	l.iterateReverse(func(e *Element[E]) bool {
		return f(e.Value)
	})
}

// Iterator to iterate all elements
func (l *List[E]) iterate(f func(e *Element[E]) bool) {
	e := l.Front()
	for e != nil {
		next := e.Next()
		if !f(e) {
			break
		}
		e = next
	}
}

// Iterator to iterate all elements
func (l *List[E]) iterateReverse(f func(e *Element[E]) bool) {
	e := l.Back()
	for e != nil {
		next := e.Prev()
		if !f(e) {
			break
		}
		e = next
	}
}

// Contains to check if contains target element value in list.
func (l *List[E]) Contains(v E, f base.EQL[E]) (contains bool) {
	return l.Index(v, f) != -1
}

// Index return the index of the first matched object in list
func (l *List[E]) Index(v E, f base.EQL[E]) (index int) {
	index = -1
	matched := false
	l.Iterate(func(e E) bool {
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

// Index return the last index of the first matched object in list
func (l *List[E]) LastIndex(v E, f base.EQL[E]) (index int) {
	index = l.Len()
	matched := false
	l.IterateReverse(func(e E) bool {
		index--
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

// Clear to remove all elements
func (l *List[E]) Clear() {
	l.iterate(func(e *Element[E]) bool {
		l.remove(e)
		return true
	})
}

// Contains to check if contains target element value in list.
func (l *List[E]) Remove(v E, f base.EQL[E]) (ret E, removed bool) {
	return l.removeMatches(v, f, true)
}

// Contains to check if contains target element value in list.
func (l *List[E]) RemoveAll(v E, f base.EQL[E]) (ret E, removed bool) {
	return l.removeMatches(v, f, false)
}

// Contains to check if contains target element value in list.
func (l *List[E]) removeMatches(v E, f base.EQL[E], first bool) (ret E, removed bool) {
	l.iterate(func(e *Element[E]) bool {
		if f(v, e.Value) {
			removed = true
			ret = l.RemoveElement(e)
			return !first
		}
		return true
	})
	return
}

// Get the element value at target index
func (l *List[E]) Get(index int) (E, bool) {
	var ret E
	if !l.isElementIndex(index) {
		return ret, false
	}

	e, b := l.get(index)
	return e.Value, b
}

// Get the element value at target index
func (l *List[E]) get(index int) (*Element[E], bool) {
	if !l.isElementIndex(index) {
		return nil, false
	}

	return l.node(index), true
}

// Set set the value at target index
func (l *List[E]) Set(index int, v E) bool {
	if !l.isElementIndex(index) {
		return false
	}
	e := l.node(index)
	e.Value = v
	return true
}

// Set set the value at target index
func (l *List[E]) Add(index int, v E) bool {
	if !l.isElementIndex(index) {
		return false
	}
	e := l.node(index)
	l.InsertAfter(v, e)
	return true
}

// RemoveFront remove front element
func (l *List[E]) RemoveFront() E {
	var ret E
	if l.IsEmpty() {
		return ret
	}
	e := l.Front()
	ret = e.Value
	l.remove(e)
	return ret
}

// RemoveBack remove back element
func (l *List[E]) RemoveBack() E {
	var ret E
	if l.IsEmpty() {
		return ret
	}
	e := l.Back()
	ret = e.Value
	l.remove(e)
	return ret
}

func (l *List[E]) node(index int) *Element[E] {
	if index < (l.len >> 1) { // pos value is before middle value
		e := l.Front()
		for i := 0; i < index; i++ {
			e = e.next
		}
		return e
	} else {
		e := l.Back()
		for i := l.len - 1; i > index; i-- {
			e = e.prev
		}
		return e
	}

}

func (l *List[E]) isElementIndex(index int) bool {
	return index >= 0 && index < l.len
}

// Filter do filter element action and return a new list
func (l *List[E]) Filter(test base.Evaluate[E]) *List[E] {
	ret := NewList[E]()
	if test == nil {
		return ret
	}

	l.iterate(func(e *Element[E]) bool {
		if test(e.Value) {
			ret.PushBack(e.Value)
		}
		return true
	})

	return ret
}

// Min to find the minimum one in the list
func (l *List[E]) Min(compare base.CMP[E]) (min E) {
	return selectByCompare(l, func(o1, o2 E) int {
		return compare(o1, o2)
	})
}

// Max to find the maximum one in the list
func (l *List[E]) Max(compare base.CMP[E]) (min E) {
	return selectByCompare(l, func(o1, o2 E) int {
		return compare(o2, o1)
	})
}

func selectByCompare[E any](l *List[E], compare base.CMP[E]) (v E) {
	i := 0
	l.iterate(func(e *Element[E]) bool {
		if i == 0 {
			v = e.Value
		} else {
			if compare(v, e.Value) > 0 {
				v = e.Value
			}
		}
		i++
		return true
	})
	return
}

// Sort to sort list elements order by compare condition.
func (l *List[E]) Sort(compare base.CMP[E]) {
	sortobject := sortableList[E]{data: l, cmp: compare}
	sort.Sort(sortobject)
}
