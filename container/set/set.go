package set

import (
	"github.com/jhunters/goassist/base"
	"github.com/jhunters/goassist/maputil"
)

// Set provides one value container and keep no duplicate value
type Set[V comparable] struct {
	mp map[V]base.Null // internal mp
}

// NewMap create a new Map
func NewSet[K comparable]() *Set[K] {
	return &Set[K]{mp: make(map[K]base.Null)}
}

// Add add key into set if key exist return false and nothing changes
func (m *Set[K]) Add(key K) bool {
	_, ok := m.mp[key]
	if ok {
		return false
	}

	m.mp[key] = base.Empty
	return true
}

// IsEmpty return true if no keys
func (m *Set[K]) IsEmpty() (empty bool) {
	return m.mp == nil || len(m.mp) == 0
}

// Size return count of size
func (m *Set[K]) Size() int {
	if m.mp == nil {
		return 0
	}
	return len(m.mp)
}

// Range calls f sequentially for each key and value present in the map.
func (m *Set[K]) Range(f base.Func[bool, K]) {
	if m.mp == nil {
		return
	}

	for k := range m.mp {
		ok := f(k)
		if !ok {
			break
		}
	}
}

// Keys return all key as slice in map
func (m *Set[K]) ToArray() []K {
	ret := make([]K, m.Size())
	i := 0
	m.Range(func(key K) bool {
		ret[i] = key
		i++
		return true
	})
	return ret
}

// Clear remove all key and value
func (m *Set[K]) Clear() {
	maputil.Clear(m.mp)
}

// Copy all keys and values to a new Map
func (m *Set[K]) Copy() *Set[K] {
	ret := NewSet[K]()
	ret.mp = maputil.Clone(m.mp)
	return ret
}

// Exist return true if key exist
func (m *Set[K]) Exist(key K) bool {
	_, ok := m.mp[key]
	return ok
}

// Exist return true if key exist
func (m *Set[K]) Remove(key K) bool {
	_, ok := m.mp[key]
	if ok {
		delete(m.mp, key)
	}
	return ok
}
