/*
 * Package arrayutil provides utility api for array operation
 */
package arrayutil

import (
	"math/rand"
	"sort"

	"github.com/jhunters/goassist/base"
	"github.com/jhunters/goassist/generic"
	"github.com/jhunters/goassist/maputil"
	"github.com/jhunters/goassist/mathutil"
)

const (
	SHUFFLE_THRESHOLD = 5
)

// Sort sort array object, sort order type is decided by cmp function.
// example code:
//
//	type Student struct {
//		Name string
//	}
//
//	students := []Student{{"xml"}, {"matthew"}, {"matt"}, {"xiemalin"}}
//	Sort(students, func(e1, e2 Student) int {
//		return strings.Compare(e1.Name, e2.Name)
//	})
func Sort[E any](data []E, cmp base.CMP[E]) {
	sortobject := sortable[E]{data: data, cmp: cmp}
	sort.Sort(sortobject)
}

// Sort sort array object by order type. for more details pls visit generic.Ordered.
// asc if true means by ascending order
// example code:
// strArray := []string{"xml", "matthew"， "matt", "xiemalin"}
// SortOrdered(strArray, true)
func SortOrdered[E generic.Ordered](data []E, asc bool) {
	sortobject := sortable[E]{data: data, cmp: func(e1, e2 E) int {
		ord := -1
		if !asc {
			ord *= -1
		}
		return ord * CompareTo(e1, e2)
	}}
	sort.Sort(sortobject)
}

type sortable[E any] struct {
	data []E
	cmp  base.CMP[E]
}

func (s sortable[E]) Len() int      { return len(s.data) }
func (s sortable[E]) Swap(i, j int) { s.data[i], s.data[j] = s.data[j], s.data[i] }
func (s sortable[E]) Less(i, j int) bool {
	return s.cmp(s.data[i], s.data[j]) >= 0
}

// Shuffle randomly permutes the specified list using a default source of
// randomness.
func Shuffle[E any](data []E) {
	r := rand.New(rand.NewSource(int64(len(data))))
	ShuffleRandom(data, r)
}

// ShuffleRandom randomly permute the specified array using the specified source of
//
//	randomness.
func ShuffleRandom[E any](data []E, r *rand.Rand) {
	size := len(data)
	for i := 0; i < size; i++ {
		j := r.Intn(size)
		data[i], data[j] = data[j], data[i]
	}

}

// Shuffle read element from slice
func ShuffleRead[E any](data []E, reader func(e E)) {
	size := len(data)
	eleOrder := rand.Perm(size)
	for _, k := range eleOrder {
		reader(data[k])
	}
}

// Reverse reverses the order of the elements in the specified
func Reverse[E any](data []E) {
	size := len(data)
	mid := size >> 1
	j := size - 1
	for i := 0; i < mid; i++ {
		data[i], data[j] = data[j], data[i]
		j--
	}
}

// BinarySearch searches the specified array for the specified value using the
// binary search algorithm. return index of the search key
// note that target array must be ordered
func BinarySearch[E any](data []E, key E, cmp base.CMP[E]) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := (low + high) >> 1
		midVal := data[mid]

		r := cmp(midVal, key)

		if r < 0 {
			low = mid + 1
		} else if r > 0 {
			high = mid - 1
		} else {
			return mid // key found
		}
	}
	return -(low + 1) // key not found.

}

// BinarySearch searches the specified array for the specified value using the
// binary search algorithm. return index of the search key
func BinarySearchOrdered[E generic.Ordered](data []E, key E) int {
	return BinarySearch(data, key, CompareTo[E])
}

// Contains returns <tt>true</tt> if this array contains the specified element.
func Contains[E any](data []E, key E, equal base.EQL[E]) bool {
	size := len(data)
	if size == 0 {
		return false
	}
	for i := 0; i < size; i++ {
		if equal(data[i], key) {
			return true
		}
	}

	return false
}

// Contains returns <tt>true</tt> if this array contains the specified element.
func ContainsOrdered[E generic.Ordered](data []E, key E) bool {
	return Contains(data, key, Equals[E])
}

// ContainsAny returns if any elements from other exist in this array
func ContainsAny[E any](data, other []E, equal base.EQL[E]) bool {
	size := len(other)
	if size == 0 {
		return false
	}
	for i := 0; i < size; i++ {
		if Contains(data, other[i], equal) {
			return true
		}
	}

	return false
}

// ContainsAny returns if any elements from other exist in this array
func ContainsAnyOrdered[E generic.Ordered](data, other []E) bool {
	return ContainsAny(data, other, Equals[E])
}

// Remove removes the first same element value of the key from this array
func Remove[E any](data []E, key E, equal base.EQL[E]) ([]E, bool) {
	return removeContional(data, key, equal, false)
}

// RemoveOrdered  removes the first same element value of the key from this array
func RemoveOrdered[E generic.Ordered](data []E, key E) ([]E, bool) {
	return removeContional(data, key, Equals[E], false)
}

// Remove removes the all same element value of the key from this array
func RemoveAll[E any](data []E, key E, equal base.EQL[E]) ([]E, bool) {
	return removeContional(data, key, equal, true)
}

// Remove Removes the all same element value of the key from this array
func RemoveAllOrdered[E generic.Ordered](data []E, key E) ([]E, bool) {
	return removeContional(data, key, Equals[E], true)
}

func removeContional[E any](data []E, key E, equal base.EQL[E], all bool) ([]E, bool) {
	size := len(data)
	if size == 0 {
		return data, false
	}
	ret := make([]E, 0, size)
	for i := 0; i < size; i++ {
		if equal(data[i], key) {
			if !all {
				ret = append(ret, data[i+1:]...)
				return ret, true
			}
		} else {
			ret = append(ret, data[i])
		}
	}

	return ret, true
}

// RemoveIndex remove element by index
func RemoveIndex[E any](data []E, i int) []E {
	size := len(data)
	if i >= size || i < 0 { // out of index, do nothing
		return data
	}

	if i < size-1 {
		copy(data[i:], data[i+1:])
	}
	return data[:len(data)-1]
}

// Min Returns the minimum element and position of the given array
func Min[E any](data []E, cmp base.CMP[E]) (E, int) {
	size := len(data)
	if size == 1 {
		return data[0], 0
	}
	ret := data[0]
	pos := 0
	for i := 1; i < size; i++ {
		if cmp(ret, data[i]) > 0 {
			ret = data[i]
			pos = i
		}
	}

	return ret, pos
}

// MinOrdered Returns the minimum element and position of the given array
func MinOrdered[E generic.Ordered](data []E) (E, int) {
	return Min(data, func(e1, e2 E) int {
		return CompareTo(e1, e2)
	})
}

// Max Returns the maximum element and position of the given array
func Max[E any](data []E, cmp base.CMP[E]) (E, int) {
	size := len(data)
	if size == 1 {
		return data[0], 0
	}
	ret := data[0]
	pos := 0
	for i := 1; i < size; i++ {
		if cmp(ret, data[i]) < 0 {
			ret = data[i]
			pos = i
		}
	}
	return ret, pos
}

// MaxOrdered Returns the maximum element and position of the given array
func MaxOrdered[E generic.Ordered](data []E) (E, int) {
	return Max(data, CompareTo[E])
}

// ReplaceAll Replaces all occurrences of one specified value in a array with another
func ReplaceAll[E any](data []E, oldVal, newVal E, euqal base.EQL[E]) {
	if data == nil {
		return
	}
	size := len(data)
	for i := 0; i < size; i++ {
		if euqal(data[i], oldVal) {
			data[i] = newVal
		}
	}
}

// ReplaceOrderedAll Replaces all occurrences of one specified value in a array with another
func ReplaceOrderedAll[E generic.Ordered](data []E, oldVal, newVal E) {
	ReplaceAll(data, oldVal, newVal, Equals[E])
}

// EqualWith to theck all elements of the two array are same
func EqualWith[E any](data, other []E, euqal base.EQL[E]) bool {
	s1, s2 := len(data), len(other)
	if s1 != s2 {
		return false
	}

	for i := 0; i < s1; i++ {
		if !euqal(data[i], other[i]) {
			return false
		}
	}
	return true
}

// EqualWithOrdered to theck all elements of the two array are same
func EqualWithOrdered[E generic.Ordered](data, other []E) bool {
	return EqualWith(data, other, Equals[E])

}

// Filter to filter target array by specified tester.
func Filter[E any](data []E, evaluate base.Evaluate[E]) []E {
	ret := make([]E, 0)
	for _, v := range data {
		if !evaluate(v) {
			ret = append(ret, v)
		}
	}

	return ret
}

// IndexOfSubArrayReturns the starting position of the first occurrence of the specified
//
//	target array within the specified source array
func IndexOfSubArray[E any](data, sub []E, euqal base.EQL[E]) int {
	s1, s2 := len(data), len(sub)
	if s2 == 0 {
		return -1
	}

	if s2 > s1 {
		return -1
	}

	checkSize := s2
	beginPos := 0
	for beginPos+checkSize <= s1 {
		if EqualWith(data[beginPos:beginPos+checkSize], sub, euqal) {
			return beginPos
		}
		beginPos++
	}
	return -1
}

// IndexOfSubOrderedArray the starting position of the first occurrence of the specified
//
//	target array within the specified source array
func IndexOfSubOrderedArray[E generic.Ordered](data, sub []E) int {
	return IndexOfSubArray(data, sub, Equals[E])
}

// LastIndexOfSubArray the last starting position of the first occurrence of the specified
//
//	target array within the specified source array
func LastIndexOfSubArray[E any](data, sub []E, euqal base.EQL[E]) int {
	s1, s2 := len(data), len(sub)
	if s2 == 0 {
		return -1
	}

	if s2 > s1 {
		return -1
	}

	checkSize := s2
	beginPos := s1 - checkSize
	for beginPos >= 0 {
		if EqualWith(data[beginPos:beginPos+checkSize], sub, euqal) {
			return beginPos
		}
		beginPos--
	}

	return -1

}

// LastIndexOfSubArray the last starting position of the first occurrence of the specified
//
//	target array within the specified source array
func LastIndexOfSubOrderedArray[E generic.Ordered](data, sub []E) int {
	return LastIndexOfSubArray(data, sub, Equals[E])
}

// Disjoint Returns true if the two specified collections have no
// elements in common.
func Disjoint[E any](data []E, other []E, euqal base.EQL[E]) bool {
	s1, s2 := len(data), len(other)
	if s1 == 0 || s2 == 0 {
		return true
	}

	for i := 0; i < s1; i++ {
		if Contains(other, data[i], euqal) {
			return false
		}
	}

	return true
}

// Disjoint Returns true if the two specified collections have no
// elements in common.
func DisjointOrdered[E generic.Ordered](data []E, other []E) bool {
	return Disjoint(data, other, Equals[E])
}

// Rotate Rotates the elements in the specified array by the specified distance.
// For example, suppose a string array arr := []string{"t", "a", "n", "k", "s"}.
// After invoking arrays.Rotate(arr, 1) (or arrays.Rotate(arr, -4)),  output is [s, t, a, n, k].
func Rotate[E any](data []E, distance int) {

	size := len(data)
	if size == 0 {
		return
	}
	mid := -distance % size
	if mid < 0 {
		mid += size
	}
	if mid == 0 {
		return
	}

	Reverse(data[0:mid])
	Reverse(data[mid:size])
	Reverse(data)

}

func getFreq[E generic.Ordered](key E, mapa map[E]int) int {
	v, exist := mapa[key]
	if !exist {
		return 0
	}
	return v
}

// UnionOrdered returns a array containing the union
// of the given array.
func UnionOrdered[E generic.Ordered](data, other []E) []E {
	ret := make([]E, 0)

	mapa := getCardinalityMap(data)
	mapb := getCardinalityMap(other)

	merged := maputil.AddAll(mapa, mapb)
	for k := range merged {
		i := 0
		for m := mathutil.Max(int(getFreq(k, mapa)), int(getFreq(k, mapb))); i < m; i++ {
			ret = append(ret, k)
		}
	}

	return ret
}

// IntersectionOrdered returns a array containing the intersection
// of the given array.
func IntersectionOrdered[E generic.Ordered](data, other []E) []E {
	ret := make([]E, 0)

	mapa := getCardinalityMap(data)
	mapb := getCardinalityMap(other)

	merged := maputil.AddAll(mapa, mapb)
	for k := range merged {
		i := 0
		for m := mathutil.Min(int(getFreq(k, mapa)), int(getFreq(k, mapb))); i < m; i++ {
			ret = append(ret, k)
		}
	}

	return ret
}

// DisjunctionOrdered returns a array containing the exclusive disjunction
// (symmetric difference) of the given array
func DisjunctionOrdered[E generic.Ordered](data, other []E) []E {
	ret := make([]E, 0)

	mapa := getCardinalityMap(data)
	mapb := getCardinalityMap(other)

	merged := maputil.AddAll(mapa, mapb)
	for k := range merged {
		i := 0
		m := mathutil.Max(int(getFreq(k, mapa)), int(getFreq(k, mapb))) - mathutil.Min(int(getFreq(k, mapa)), int(getFreq(k, mapb)))
		for ; i < m; i++ {
			ret = append(ret, k)
		}
	}

	return ret
}

// Subtract returns a new array containing data - other.
func Substract[E any](data, other []E, equal base.EQL[E]) []E {
	ret := Clone(data)

	for _, e := range other {
		ret, _ = Remove(ret, e, equal)
	}
	return ret
}

// SubtractOrdered returns a new array containing data - other.
func SubstractOrdered[E generic.Ordered](data, other []E) []E {
	ret := Clone(data)
	for _, v := range other {
		ret, _ = RemoveOrdered(ret, v)
	}

	return ret
}

func getCardinalityMap[E generic.Ordered](data []E) map[E]int {

	ret := make(map[E]int)
	for _, e := range data {
		_, exist := ret[e]
		if exist {
			ret[e] = ret[e] + 1
		} else {
			ret[e] = 1
		}

	}
	return ret
}

// CompareTo Compares this object with the specified object for order.  Returns a
//
//	negative integer, zero, or a positive integer as this object is less
//	 than, equal to, or greater than the specified object.
func CompareTo[E generic.Ordered](e1, e2 E) int {
	if e1 > e2 {
		return 1
	} else if e1 < e2 {
		return -1
	}
	return 0
}

// Clone Copies all of the elements and return a new array
func Clone[E any](data []E) []E {
	size := len(data)
	ret := make([]E, size, cap(data))
	copy(ret, data)
	return ret
}

// Insert target value in array, if index < 0 or > len(data) will do noting
func Insert[E any](data []E, index int, v E) []E {
	if index < 0 || index > len(data) {
		return data
	}
	if index == len(data) {
		return append(data, v)
	}
	var empty E
	ret := append(data, empty)
	copy(ret[index+1:], ret[index:])
	ret[index] = v
	return ret
}

// Equals to compare two value is equal
func Equals[E generic.Ordered](e1, e2 E) bool {
	return e1 == e2
}

// CreateAndFill create array with the same element t
func CreateAndFill[E any](size int, defaultElementValue E) []E {
	ret := make([]E, size)
	for i := 0; i < size; i++ {
		ret[i] = defaultElementValue
	}
	return ret
}

// AsList convert to array
func AsList[E any](e ...E) []E {
	if e == nil {
		return nil
	}
	ret := make([]E, len(e))
	copy(ret, e)
	return ret
}
