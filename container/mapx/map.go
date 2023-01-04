package mapx

import (
	"github.com/jhunters/goassist/base"
	"github.com/jhunters/goassist/maputil"
	"github.com/jhunters/goassist/reflectutil"
)

// Map is like a Go map[interface{}]interface{} but is provide more useful methods
type Map[K comparable, V any] struct {
	mp    map[K]V
	empty V
}

// NewMap create a new Map
func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{mp: make(map[K]V)}
}

// Put put key and value to map
func (m *Map[K, V]) Put(key K, value V) V {
	v := m.mp[key]
	m.mp[key] = value
	return v
}

// Put put key and value to map
func (m *Map[K, V]) Get(key K) (V, bool) {
	v, ok := m.mp[key]
	return v, ok
}

// IsEmpty return true if no keys
func (m *Map[K, V]) IsEmpty() (empty bool) {
	return m.mp == nil || len(m.mp) == 0
}

// Size return count of size
func (m *Map[K, V]) Size() int {
	if m.mp == nil {
		return 0
	}
	return len(m.mp)
}

// ToMap convert key and value to origin map struct
func (m *Map[K, V]) ToMap() map[K]V {
	if m.mp == nil {
		return nil
	}

	return maputil.Clone(m.mp)
}

// Range calls f sequentially for each key and value present in the map.
func (m *Map[K, V]) Range(f base.BiFunc[bool, K, V]) {
	if m.mp == nil {
		return
	}

	for k, v := range m.mp {
		ok := f(k, v)
		if !ok {
			break
		}
	}
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
	maputil.Clear(m.mp)
}

// Copy all keys and values to a new Map
func (m *Map[K, V]) Copy() *Map[K, V] {
	ret := NewMap[K, V]()
	m.Range(func(key K, value V) bool {
		ret.Put(key, value)
		return true
	})

	return ret
}

// Exist return true if key exist
func (m *Map[K, V]) Exist(key K) bool {
	_, ok := m.Get(key)
	return ok
}

// ExistValue return true if value exist
func (m *Map[K, V]) ExistValue(value V) (k K, exist bool) {
	de := reflectutil.NewDeepEquals(value)
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

// Exist return true if key exist
func (m *Map[K, V]) Remove(key K) bool {
	_, ok := m.mp[key]
	if ok {
		delete(m.mp, key)
	}
	return ok
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
