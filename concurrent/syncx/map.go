package syncx

import (
	"sync"

	"github.com/jhunters/goassist/base"
	"github.com/jhunters/goassist/reflectx"
)

// Map is like a Go map[interface{}]interface{} but is safe for concurrent use
// by multiple goroutines without additional locking or coordination.
// Loads, stores, and deletes run in amortized constant time.
//
// The Map type is specialized. Most code should use a plain Go map instead,
// with separate locking or coordination, for better type safety and to make it
// easier to maintain other invariants along with the map content.
//
// The Map type is optimized for two common use cases: (1) when the entry for a given
// key is only ever written once but read many times, as in caches that only grow,
// or (2) when multiple goroutines read, write, and overwrite entries for disjoint
// sets of keys. In these two cases, use of a Map may significantly reduce lock
// contention compared to a Go map paired with a separate Mutex or RWMutex.
//
// The zero Map is empty and ready for use. A Map must not be copied after first use.
type Map[K comparable, V any] struct {
	mp    sync.Map
	empty V
	mu    sync.Mutex
}

// NewMap create a new map
func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{mp: sync.Map{}}
}

// NewMapByInitial create a new map and store key and value from origin map
func NewMapByInitial[K comparable, V any](mmp map[K]V) *Map[K, V] {
	mp := NewMap[K, V]()
	if mmp == nil {
		return mp
	}
	for k, v := range mmp {
		mp.Store(k, v)
	}

	return mp
}

// Store sets the value for a key.
func (m *Map[K, V]) Store(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mp.Store(key, value)
}

// StoreAll sets all the key and values to map.
func (m *Map[K, V]) StoreAll(other *Map[K, V]) {
	if other == nil {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	other.Range(func(key K, value V) bool {
		m.mp.Store(key, value)
		return true
	})
}

// StoreAll sets all the key and values to map.
func (m *Map[K, V]) StoreAllOrigin(other map[K]V) {
	if other == nil {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	for key, value := range other {
		m.mp.Store(key, value)
	}
}

// Replace replaces the value for key if value compare condition
func (m *Map[K, V]) Replace(key K, oldValue, newValue V, equal base.EQL[V]) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	v, ok := m.Load(key)
	if !ok {
		return false
	}

	if equal(v, oldValue) {
		m.mp.Store(key, newValue)
		return true
	}

	return false
}

// ReplaceByCondition replaces the value for key if value compare condition
func (m *Map[K, V]) ReplaceByCondition(key K, c base.BiFunc[V, K, V]) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	v, ok := m.Load(key)
	if !ok {
		return false
	}

	r := c(key, v)
	m.mp.Store(key, r)
	return true
}

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.mp.Load(key)
	if !ok {
		return m.empty, ok
	}
	return v.(V), ok
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.mp.LoadOrStore(key, value)
	return v.(V), ok
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.mp.LoadAndDelete(key)
	if !ok {
		return m.empty, ok
	}
	return v.(V), ok
}

// Delete deletes the value for a key.
func (m *Map[K, V]) Delete(key K) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mp.Delete(key)
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently (including by f), Range may reflect any
// mapping for that key from any point during the Range call. Range does not
// block other methods on the receiver; even f itself may call any method on m.
//
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
func (m *Map[K, V]) Range(f base.BiFunc[bool, K, V]) {
	m.mp.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

// Exist return true if key exist
func (m *Map[K, V]) Exist(key K) bool {
	_, ok := m.mp.Load(key)
	return ok
}

// ExistValue return true if value exist
func (m *Map[K, V]) ExistValue(value V) (k K, exist bool) {
	de := reflectx.NewDeepEquals(value)
	m.Range(func(key K, val V) bool {
		if de.Matches(val) {
			exist = true
			k = key
			return false
		}
		return true
	})
	return
}

// ExistValue return true if value exist
func (m *Map[K, V]) ExistValueWithComparator(value V, equal base.EQL[V]) (k K, exist bool) {
	m.Range(func(key K, val V) bool {
		if equal(value, val) {
			exist = true
			k = key
			return false
		}
		return true
	})
	return
}

// ExistValue return true if value exist
func (m *Map[K, V]) ExistValueComparable(v base.Comparable[V]) (k K, exist bool) {
	m.Range(func(key K, val V) bool {
		if v.CompareTo(val) == 0 {
			exist = true
			k = key
			return false
		}
		return true
	})
	return
}

// IsEmpty return true if no keys
func (m *Map[K, V]) IsEmpty() (empty bool) {
	empty = true
	m.mp.Range(func(key, value any) bool {
		empty = false
		return false
	})
	return
}

// Size return count of size
func (m *Map[K, V]) Size() int {
	if m == nil {
		return 0
	}
	count := 0
	m.mp.Range(func(key, value any) bool {
		count++
		return true
	})
	return count
}

// ToMap convert key and value to origin map struct
func (m *Map[K, V]) ToMap() map[K]V {
	ret := make(map[K]V)
	m.Range(func(key K, value V) bool {
		ret[key] = value
		return true
	})
	return ret
}

// Values return all value as slice in map
func (m *Map[K, V]) Values() []V {
	ret := make([]V, 0)
	m.Range(func(key K, value V) bool {
		ret = append(ret, value)
		return true
	})
	return ret
}

// Keys return all key as slice in map
func (m *Map[K, V]) Keys() []K {
	ret := make([]K, 0)
	m.Range(func(key K, value V) bool {
		ret = append(ret, key)
		return true
	})
	return ret
}

// Clear remove all key and value
func (m *Map[K, V]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()

	keys := m.Keys()
	for _, k := range keys {
		m.mp.Delete(k)
	}
}

// Copy all keys and values to a new Map
func (m *Map[K, V]) Copy() *Map[K, V] {
	ret := NewMap[K, V]()
	m.Range(func(key K, value V) bool {
		ret.Store(key, value)
		return true
	})

	return ret
}

// MinValue to return min value in the map
func (m *Map[K, V]) MinValue(compare base.CMP[V]) (key K, v V) {
	return selectByCompareValue(m, func(o1, o2 V) int {
		return compare(o1, o2)
	})

}

// MaxValue to return max value in the map
func (m *Map[K, V]) MaxValue(compare base.CMP[V]) (key K, v V) {
	return selectByCompareValue(m, func(o1, o2 V) int {
		return compare(o2, o1)
	})

}

func selectByCompareValue[K comparable, V any](mp *Map[K, V], compare base.CMP[V]) (key K, v V) {
	var ret V
	i := 0
	mp.Range(func(k K, v V) bool {
		if i == 0 {
			ret = v
			key = k
		} else {
			if compare(ret, v) > 0 {
				ret = v
				key = k
			}
		}
		i++
		return true
	})
	return key, ret
}

// MinKey to return min key in the map
func (m *Map[K, V]) MinKey(compare base.CMP[K]) (key K, v V) {
	return selectByCompareKey(m, func(o1, o2 K) int {
		return compare(o1, o2)
	})

}

// MaxKey to return max key in the map
func (m *Map[K, V]) MaxKey(compare base.CMP[K]) (key K, v V) {
	return selectByCompareKey(m, func(o1, o2 K) int {
		return compare(o2, o1)
	})

}

func selectByCompareKey[K comparable, V any](mp *Map[K, V], compare base.CMP[K]) (key K, value V) {
	var ret K
	i := 0
	mp.Range(func(k K, v V) bool {
		if i == 0 {
			ret = k
			value = v
		} else {
			if compare(ret, k) > 0 {
				ret = k
				value = v
			}
		}
		i++
		return true
	})
	return ret, value
}
