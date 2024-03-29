/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2021-12-24 13:55:47
 */
package arrayutil_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jhunters/goassist/arrayutil"
	"github.com/jhunters/goassist/conv"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	N1 = "xml"
	N2 = "matthew"
	N3 = "xiemalin"
)

var (
	persons = []Person{{"xml", 100}, {"matthew", 90}, {"xiemalin", 99}}

	persons2 = []Person{{"xml", 110}, {"matthew", 90}, {"xiemalin", 95}}

	sortedPersons = []Person{{"xml", 90}, {"matthew", 91}, {"xiemalin", 92}, {"xiemalin2", 93}, {"xiemalin3", 94}, {"xiemalin4", 95}, {"xiemalin5", 96}}

	strArray = []string{"xml", "matthew", "xiemalin", "xml"}

	strArray2 = []string{"xml", "hello", "world", "xml", "xiemalin", "xml"}

	sortedIntArray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
)

type Person struct {
	Name string
	Age  int8
}

func (p Person) Equal(other Person) bool {
	return p.Age == other.Age && p.Name == other.Name
}

func TestSort(t *testing.T) {
	Convey("Test sort struct array", t, func() {

		persons := arrayutil.Clone(persons)
		Convey("Test Person struct array desc order", func() {
			arrayutil.Sort(persons, func(e1, e2 Person) int {
				return int(e1.Age) - int(e2.Age)
			})
			So(persons[0].Name, ShouldEqual, "xml")
			So(len(persons), ShouldEqual, 3)
		})

		Convey("Test Person struct array asc order", func() {
			arrayutil.Sort(persons, func(e1, e2 Person) int {
				return int(e2.Age) - int(e1.Age)
			})
			So(persons[0].Name, ShouldEqual, "matthew")
			So(len(persons), ShouldEqual, 3)
		})
	})

	Convey("Test sort builtin type array", t, func() {
		strArray := []string{"xml", "matthew", "xiemalin"}
		arrayutil.Sort(strArray, func(e1, e2 string) int {
			return arrayutil.CompareTo(e1, e2)
		})
		So(strArray[0], ShouldEqual, "xml")
	})
}

func BenchmarkSort(b *testing.B) {
	persons := arrayutil.Clone(persons)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arrayutil.Sort(persons, func(e1, e2 Person) int {
			return int(e1.Age) - int(e2.Age)
		})
	}
}

func ExampleSort() {
	// order by person's age
	persons = []Person{{"xml", 100}, {"matthew", 90}, {"xiemalin", 99}}
	arrayutil.Sort(persons, func(e1, e2 Person) int {
		return int(e1.Age) - int(e2.Age)
	})
	fmt.Println(persons)
	//output:
	//[{xml 100} {xiemalin 99} {matthew 90}]
}

func TestSortOrdered(t *testing.T) {
	Convey("Test SortOrdered", t, func() {
		strArray := arrayutil.Clone(strArray)
		arrayutil.SortOrdered(strArray, false)
		So(strArray[0], ShouldEqual, "xml")

		arrayutil.SortOrdered(strArray, true)
		So(strArray[0], ShouldEqual, "matthew")
	})
}

func ExampleSortOrdered() {
	strArray := arrayutil.Clone(strArray)
	arrayutil.SortOrdered(strArray, false)
	fmt.Println(strArray)

	//output:
	//[xml xml xiemalin matthew]
}

func TestShuffle(t *testing.T) {
	Convey("Test Shuffle", t, func() {

		sortedIntArray := arrayutil.Clone(sortedIntArray)
		arrayutil.Shuffle(sortedIntArray)

		shuffled := false
		size := len(sortedIntArray)
		for i := 0; i < size; i++ {
			if sortedIntArray[i] != i {
				shuffled = true
			}

		}

		So(shuffled, ShouldBeTrue)
	})
}

func TestReverse(t *testing.T) {
	Convey("Test Shuffle", t, func() {

		sortedIntArray := arrayutil.Clone(sortedIntArray)
		arrayutil.Reverse(sortedIntArray)

		reversed := true
		size := len(sortedIntArray)
		j := size
		for i := 0; i < size; i++ {
			if sortedIntArray[i] != j {
				reversed = false
			}
			j--

		}

		So(reversed, ShouldBeTrue)
	})
}

func ExampleReverse() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arrayutil.Reverse(slice)
	fmt.Println(slice)

	// Output:
	// [10 9 8 7 6 5 4 3 2 1]
}

func TestBinarySearch(t *testing.T) {
	Convey("Test BinarySearch", t, func() {

		persons := arrayutil.Clone(sortedPersons)

		cmp := func(e1, e2 Person) int {
			return arrayutil.CompareTo(e1.Age, e2.Age)
		}

		Convey("Test BinarySearch exist", func() {
			key := Person{"matthew", 90}

			offset := arrayutil.BinarySearch(persons, key, cmp)
			So(offset, ShouldEqual, 0)

			key = Person{"matthew", 96}
			offset = arrayutil.BinarySearch(persons, key, cmp)
			So(offset, ShouldEqual, 6)

			key = Person{"matthew", 92}
			offset = arrayutil.BinarySearch(persons, key, cmp)
			So(offset, ShouldEqual, 2)
		})

		Convey("Test BinarySearch no exist", func() {

			key := Person{"matthew", 89}
			offset := arrayutil.BinarySearch(persons, key, cmp)
			So(offset < 0, ShouldBeTrue)

			key = Person{"matthew", 100}
			offset = arrayutil.BinarySearch(persons, key, cmp)
			So(offset < 0, ShouldBeTrue)
		})

	})
}

func ExampleBinarySearch() {
	sortedPersons = []Person{{"xml", 90}, {"matthew", 91}, {"xiemalin", 92}, {"xiemalin2", 93}, {"xiemalin3", 94}, {"xiemalin4", 95}, {"xiemalin5", 96}}
	key := Person{"matthew", 90}
	offset := arrayutil.BinarySearch(sortedPersons, key, func(e1, e2 Person) int {
		return arrayutil.CompareTo(e1.Age, e2.Age)
	})

	fmt.Println(offset, sortedPersons[offset].Age)
	//output:
	//0 90
}

func ExampleBinarySearchOrdered() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	offset := arrayutil.BinarySearchOrdered(slice, 2)

	fmt.Println(offset, slice[offset])
	//output:
	//1 2
}

func TestBinarySearchOrdered(t *testing.T) {
	Convey("Test BinarySearch", t, func() {
		sortedIntArray := arrayutil.Clone(sortedIntArray)

		cmp := func(e1, e2 int) int {
			return e1 - e2
		}
		Convey("Test BinarySearch exist", func() {

			offset := arrayutil.BinarySearchOrdered(sortedIntArray, 1)
			So(offset, ShouldEqual, 0)

			offset = arrayutil.BinarySearch(sortedIntArray, 10, cmp)
			So(offset, ShouldEqual, 9)

			offset = arrayutil.BinarySearchOrdered(sortedIntArray, 5)
			So(offset, ShouldEqual, 4)
		})

		Convey("Test BinarySearch no exist", func() {

			offset := arrayutil.BinarySearchOrdered(sortedIntArray, 0)
			So(offset < 0, ShouldBeTrue)

			offset = arrayutil.BinarySearchOrdered(sortedIntArray, 11)
			So(offset < 0, ShouldBeTrue)
		})

	})
}

func TestMax(t *testing.T) {
	Convey("Test Max", t, func() {

		persons := arrayutil.Clone(persons)
		Oldest, pos := arrayutil.Max(persons, func(e1, e2 Person) int {
			return int(e1.Age) - int(e2.Age)
		})
		So(Oldest.Age, ShouldEqual, 100)
		So(pos, ShouldEqual, 0)

		sortedPersons := arrayutil.Clone((sortedPersons))
		Oldest, pos = arrayutil.Max(sortedPersons, func(e1, e2 Person) int {
			return int(e1.Age) - int(e2.Age)
		})
		So(Oldest.Age, ShouldEqual, 96)
		So(pos, ShouldEqual, len(sortedPersons)-1)

		oneEle := []Person{{"hello", -100}}
		one, pos := arrayutil.Max(oneEle, func(e1, e2 Person) int {
			return int(e1.Age) - int(e2.Age)
		})
		So(one.Age, ShouldEqual, -100)
		So(pos, ShouldEqual, 0)
	})
}

func ExampleMax() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	v, pos := arrayutil.Max(slice, func(i, i2 int) int {
		return i - i2
	})
	fmt.Println(v, pos)

	// Output:
	// 10 9
}

func TestMaxOrdered(t *testing.T) {
	Convey("Test MaxOrdered", t, func() {

		strArray := arrayutil.Clone(strArray)
		biggestStr, pos := arrayutil.MaxOrdered(strArray)
		So(biggestStr, ShouldEqual, "xml")
		So(pos, ShouldEqual, 0)

		sortedIntArray := arrayutil.Clone(sortedIntArray)
		biggestInt, pos := arrayutil.MaxOrdered(sortedIntArray)
		So(biggestInt, ShouldEqual, 10)
		So(pos, ShouldEqual, len(sortedIntArray)-1)
	})
}

func ExampleMaxOrdered() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	v, pos := arrayutil.MaxOrdered(slice)
	fmt.Println(v, pos)

	// Output:
	// 10 9
}

func TestMin(t *testing.T) {
	Convey("Test Min", t, func() {

		persons := arrayutil.Clone(persons)
		Oldest, pos := arrayutil.Min(persons, func(e1, e2 Person) int {
			return int(e1.Age) - int(e2.Age)
		})
		So(Oldest.Age, ShouldEqual, 90)
		So(pos, ShouldEqual, 1)

		sortedPersons := arrayutil.Clone((sortedPersons))
		Oldest, pos = arrayutil.Min(sortedPersons, func(e1, e2 Person) int {
			return int(e1.Age) - int(e2.Age)
		})
		So(Oldest.Age, ShouldEqual, 90)
		So(pos, ShouldEqual, 0)

		oneEle := []Person{{"hello", -100}}
		one, pos := arrayutil.Min(oneEle, func(e1, e2 Person) int {
			return int(e1.Age) - int(e2.Age)
		})
		So(one.Age, ShouldEqual, -100)
		So(pos, ShouldEqual, 0)
	})
}

func ExampleMin() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	v, pos := arrayutil.Min(slice, func(i, i2 int) int { return i - i2 })
	fmt.Println(v, pos)

	// Output:
	// 1 0
}

func TestMinOrdered(t *testing.T) {
	Convey("Test MinOrdered", t, func() {

		strArray := arrayutil.Clone(strArray)
		biggestStr, pos := arrayutil.MinOrdered(strArray)
		So(biggestStr, ShouldEqual, "matthew")
		So(pos, ShouldEqual, 1)

		sortedIntArray := arrayutil.Clone(sortedIntArray)
		biggestInt, pos := arrayutil.MinOrdered(sortedIntArray)
		So(biggestInt, ShouldEqual, 1)
		So(pos, ShouldEqual, 0)
	})
}

func ExampleMinOrdered() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	v, pos := arrayutil.MinOrdered(slice)
	fmt.Println(v, pos)

	// Output:
	// 1 0
}

func TestRelaceAll(t *testing.T) {
	Convey("Test RelaceOrderedAll", t, func() {

		persons := arrayutil.Clone(persons)
		manReplaceArray := arrayutil.Clone(persons)
		manReplaceArray[1] = Person{"xiemalin", 100}
		oldPerson := Person{"matthew", 90}
		newPerson := Person{"xiemalin", 100}
		arrayutil.ReplaceAll(persons, oldPerson, newPerson, func(e1, e2 Person) bool {
			return e1.Equal(e2)
		})

		So(persons, ShouldResemble, manReplaceArray)

		// replace not exist
		notexistPerson := Person{"michael", 10}
		arrayutil.ReplaceAll(persons, notexistPerson, newPerson, func(e1, e2 Person) bool {
			return e1.Equal(e2)
		})
		So(persons, ShouldResemble, manReplaceArray)
	})
}

func ExampReplaceAll() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arrayutil.ReplaceAll(slice, 5, 10, func(i, i2 int) bool { return i == i2 })
	fmt.Println(slice)

	// Output:
	// [1 2 3 4 10 6 7 8 9 10]
}

func TestRelaceOrderedAll(t *testing.T) {
	Convey("Test RelaceOrderedAll", t, func() {

		strArray := arrayutil.Clone(strArray)
		manReplaceArray := arrayutil.Clone(strArray)
		manReplaceArray[1] = "xiemalin"
		arrayutil.ReplaceOrderedAll(strArray, "matthew", "xiemalin")

		So(len(strArray), ShouldEqual, 4)
		So(strArray, ShouldResemble, manReplaceArray)

		// replace not exist
		arrayutil.ReplaceOrderedAll(strArray, "matthew", "xml")

	})
}

func ExampleReplaceAll() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arrayutil.ReplaceAll(slice, 5, 10, func(i, i2 int) bool { return i == i2 })
	fmt.Println(slice)

	// Output:
	// [1 2 3 4 10 6 7 8 9 10]
}

func ExampleReplaceOrderedAll() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arrayutil.ReplaceOrderedAll(slice, 5, 10)
	fmt.Println(slice)

	// Output:
	// [1 2 3 4 10 6 7 8 9 10]
}

func TestCreateAndFill(t *testing.T) {
	Convey("Test CreateAndFill", t, func() {
		strArray := arrayutil.CreateAndFill(10, "name")
		So(len(strArray), ShouldEqual, 10)
		So(strArray[0], ShouldEqual, "name")
		So(strArray[4], ShouldEqual, "name")
		So(strArray[9], ShouldEqual, "name")

		Convey("Test CreateAndFill with zero size", func() {
			strArray := arrayutil.CreateAndFill(0, "name")
			So(len(strArray), ShouldEqual, 0)
		})

	})
}

func ExampleCreateAndFill() {
	strArray := arrayutil.CreateAndFill(10, 1)
	fmt.Println(strArray)

	// Output:
	// [1 1 1 1 1 1 1 1 1 1]
}

func TestIndexOfSubArray(t *testing.T) {
	Convey("Test IndexOfSubArray", t, func() {
		Convey("Test IndexOfSubArray with not found", func() {

			persons := arrayutil.Clone(persons)
			subPersons := []Person{}

			pos := arrayutil.IndexOfSubArray(persons, subPersons, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, -1)

			// the length of sub array is large than source
			subPersons2 := append(persons, Person{})
			pos = arrayutil.IndexOfSubArray(persons, subPersons2, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, -1)
		})

		Convey("Test IndexOfSubArray start with sub array", func() {
			// with same array
			persons := arrayutil.Clone(persons)
			pos := arrayutil.IndexOfSubArray(persons, persons, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, 0)

			// with sub pos at 0
			persons = arrayutil.Clone(persons)
			sub := arrayutil.Clone(persons)

			persons = append(persons, sub...)
			pos = arrayutil.IndexOfSubArray(persons, persons, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, 0)

		})

		Convey("Test IndexOfSubArray with has sub array", func() {
			// with sub pos at some position
			newPersons := arrayutil.Clone(sortedPersons)
			newPersons = append(newPersons, persons...)
			pos := arrayutil.IndexOfSubArray(newPersons, persons, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, len(sortedPersons))
		})

		Convey("Test IndexOfSubArray with has two same sub array", func() {
			// with sub pos at some position
			newPersons := arrayutil.Clone(sortedPersons)
			newPersons = append(newPersons, persons...)
			newPersons = append(newPersons, persons...)
			pos := arrayutil.IndexOfSubArray(newPersons, persons, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, len(sortedPersons))
		})

	})
}

func ExampleIndexOfSubArray() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pos := arrayutil.IndexOfSubArray(slice, []int{2, 3, 4}, func(i, i2 int) bool { return i == i2 })
	fmt.Println(pos)

	// Output:
	// 1
}

func TestIndexOfSubOrderedArray(t *testing.T) {
	Convey("Test IndexOfSubOrderedArray", t, func() {
		strArray := arrayutil.Clone(strArray)
		index := arrayutil.IndexOfSubOrderedArray(strArray, []string{"xml"})
		So(index, ShouldEqual, 0)
	})
}

func ExampleIndexOfSubOrderedArray() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pos := arrayutil.IndexOfSubOrderedArray(slice, []int{2, 3, 4})
	fmt.Println(pos)

	// Output:
	// 1
}

func TestLastIndexOfSubOrderedArray(t *testing.T) {
	Convey("Test LastIndexOfSubOrderedArray", t, func() {
		strArray := arrayutil.Clone(strArray)
		index := arrayutil.LastIndexOfSubOrderedArray(strArray, []string{"xml"})
		So(index, ShouldEqual, 3)
	})
}

func ExampleLastIndexOfSubOrderedArray() {
	slice := []int{1, 2, 3, 4, 5, 6, 2, 3, 4, 10}
	pos := arrayutil.LastIndexOfSubOrderedArray(slice, []int{2, 3, 4})
	fmt.Println(pos)

	// Output:
	// 6
}

func TestLastIndexOfSubArray(t *testing.T) {
	Convey("Test LastIndexOfSubArray", t, func() {
		Convey("Test LastIndexOfSubArray with not found", func() {

			persons := arrayutil.Clone(persons)
			subPersons := []Person{}

			pos := arrayutil.LastIndexOfSubArray(persons, subPersons, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, -1)

			// the length of sub array is large than source
			subPersons2 := append(persons, Person{})
			pos = arrayutil.LastIndexOfSubArray(persons, subPersons2, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, -1)
		})

		Convey("Test LastIndexOfSubArray start with sub array", func() {
			// with same array
			persons := arrayutil.Clone(persons)
			pos := arrayutil.LastIndexOfSubArray(persons, persons, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, 0)

			// with sub pos at 0
			persons = arrayutil.Clone(persons)
			sub := arrayutil.Clone(persons)

			persons = append(persons, sub...)
			pos = arrayutil.LastIndexOfSubArray(persons, persons, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, 0)

		})

		Convey("Test LastIndexOfSubArray with has sub array", func() {
			// with sub pos at some position
			newPersons := arrayutil.Clone(sortedPersons)
			newPersons = append(newPersons, persons...)
			pos := arrayutil.LastIndexOfSubArray(newPersons, persons, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, len(sortedPersons))
		})

		Convey("Test LastIndexOfSubArray with has two same sub array", func() {
			// with sub pos at some position
			newPersons := arrayutil.Clone(sortedPersons)
			newPersons = append(newPersons, persons...)
			newPersons = append(newPersons, persons...)
			pos := arrayutil.LastIndexOfSubArray(newPersons, persons, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, len(sortedPersons)+len(persons))
		})

	})
}

func ExampleLastIndexOfSubArray() {
	slice := []int{1, 2, 3, 4, 5, 6, 2, 3, 4, 10}
	pos := arrayutil.LastIndexOfSubArray(slice, []int{2, 3, 4}, func(i, i2 int) bool { return i == i2 })
	fmt.Println(pos)

	// Output:
	// 6
}

func TestDisJoint(t *testing.T) {
	Convey("Test DisJoint", t, func() {

		hasNoSame := arrayutil.Disjoint(persons, persons2, func(e1, e2 Person) bool {
			return e1.Equal(e2)
		})
		So(hasNoSame, ShouldBeFalse)

		otherPersons := []Person{}
		hasNoSame = arrayutil.Disjoint(persons, otherPersons, func(e1, e2 Person) bool {
			return e1.Equal(e2)
		})
		So(hasNoSame, ShouldBeTrue)

		hasNoSame = arrayutil.Disjoint(otherPersons, persons, func(e1, e2 Person) bool {
			return e1.Equal(e2)
		})
		So(hasNoSame, ShouldBeTrue)

	})
}

func ExampleDisjoint() {
	slice := []int{1, 2, 3, 4, 5, 6, 2, 3, 4, 10}
	subSlice := []int{2, 4, 6}
	exist := arrayutil.Disjoint(slice, subSlice, func(i, i2 int) bool { return i == i2 })
	fmt.Println(exist)

	subSlice = []int{11, 14, 16}
	exist = arrayutil.Disjoint(slice, subSlice, func(i, i2 int) bool { return i == i2 })
	fmt.Println(exist)

	// Output:
	// false
	// true
}

func TestDisJointOrdered(t *testing.T) {
	Convey("Test DisjointOrdered", t, func() {
		hasNoSame := arrayutil.DisjointOrdered(strArray, strArray)
		So(hasNoSame, ShouldBeFalse)

	})
}

func ExampleDisjointOrdered() {
	slice := []int{1, 2, 3, 4, 5, 6, 2, 3, 4, 10}
	subSlice := []int{2, 4, 6}
	exist := arrayutil.DisjointOrdered(slice, subSlice)
	fmt.Println(exist)

	subSlice = []int{11, 14, 16}
	exist = arrayutil.DisjointOrdered(slice, subSlice)
	fmt.Println(exist)

	// Output:
	// false
	// true
}

func TestRotate(t *testing.T) {
	Convey("Test Rotate", t, func() {

		expected := []string{"s", "t", "a", "n", "k"}
		arr := []string{"t", "a", "n", "k", "s"}
		arrayutil.Rotate(arr, 1)
		So(arr, ShouldResemble, expected)
		arr = []string{"t", "a", "n", "k", "s"}
		arrayutil.Rotate(arr, -4)
		So(arr, ShouldResemble, expected)

		expected2 := []string{"a", "n", "k", "s", "t"}
		arr = []string{"t", "a", "n", "k", "s"}
		arrayutil.Rotate(arr, 4)
		So(arr, ShouldResemble, expected2)
		arr = []string{"t", "a", "n", "k", "s"}
		arrayutil.Rotate(arr, -1)
		So(arr, ShouldResemble, expected2)
	})
}

func ExampleRotate() {
	slice := []string{"t", "a", "n", "k", "s"}
	arrayutil.Rotate(slice, 1)
	fmt.Println(slice)

	// Output:
	// [s t a n k]
}

func TestContainsOrdered(t *testing.T) {
	Convey("Test ContainsOrdered", t, func() {
		strArray := arrayutil.Clone(strArray)

		contains := arrayutil.ContainsOrdered(strArray, "xml")
		So(contains, ShouldBeTrue)

		contains = arrayutil.ContainsOrdered(strArray, "notexist")
		So(contains, ShouldBeFalse)
	})
}

func ExampleContains() {
	slice := []string{"t", "a", "n", "k", "s"}
	exist := arrayutil.Contains(slice, "a", func(s1, s2 string) bool { return strings.EqualFold(s1, s2) })
	fmt.Println(exist)

	// Output:
	// true
}

func ExampleContainsOrdered() {
	slice := []string{"t", "a", "n", "k", "s"}
	exist := arrayutil.ContainsOrdered(slice, "a")
	fmt.Println(exist)

	// Output:
	// true
}

func TestEqualWithOrdered(t *testing.T) {
	Convey("Test EqualWithOrdered", t, func() {
		strArray := arrayutil.Clone(strArray)
		emptyArray := []string{}

		equals := arrayutil.EqualWithOrdered(strArray, emptyArray)
		So(equals, ShouldBeFalse)

		equals = arrayutil.EqualWithOrdered(strArray, strArray)
		So(equals, ShouldBeTrue)
	})
}

func ExampleEqualWithOrdered() {
	slice := []string{"t", "a", "n", "k", "s"}
	equals := arrayutil.EqualWithOrdered(slice, slice)
	fmt.Println(equals)

	// Output:
	// true
}

func TestIntersectionOrdered(t *testing.T) {
	Convey("Test IntersectionOrdered", t, func() {
		result := arrayutil.IntersectionOrdered(strArray, strArray2)
		So(len(result), ShouldEqual, 3) // xml xml xiemalin

		arrayutil.SortOrdered(result, true)
		expected := []string{"xiemalin", "xml", "xml"}
		So(result, ShouldResemble, expected)
	})
}

func ExampleIntersectionOrdered() {
	slice := []string{"t", "a", "n", "k", "s"}
	slice2 := []string{"a", "b", "s"}
	slice3 := arrayutil.IntersectionOrdered(slice, slice2)
	arrayutil.SortOrdered(slice3, true)
	fmt.Println(slice3)

	// Output:
	// [a s]
}

func TestUnionOrdered(t *testing.T) {
	Convey("Test UnionOrdered", t, func() {
		result := arrayutil.UnionOrdered(strArray, strArray2)
		So(len(result), ShouldEqual, 7) // hello world matthew xml xml xml xiemalin

		arrayutil.SortOrdered(result, true)

		expected := []string{"hello", "matthew", "world", "xiemalin", "xml", "xml", "xml"}
		So(result, ShouldResemble, expected)

	})
}

func ExampleUnionOrdered() {
	slice := []string{"t", "a", "n", "k", "s"}
	slice2 := []string{"a", "b", "s"}
	slice3 := arrayutil.UnionOrdered(slice, slice2)
	arrayutil.SortOrdered(slice3, true)
	fmt.Println(slice3)

	// Output:
	// [a b k n s t]
}

func TestDisjunctionOrdered(t *testing.T) {
	Convey("Test DisjunctionOrdered", t, func() {
		result := arrayutil.DisjunctionOrdered(strArray, strArray2)
		So(len(result), ShouldEqual, 4) // hello world matthew xml xml xml xiemalin

		arrayutil.SortOrdered(result, true)

		expected := []string{"hello", "matthew", "world", "xml"}
		So(result, ShouldResemble, expected)

	})
}

func ExampleDisjunctionOrdered() {
	slice := []string{"t", "a", "n", "k", "s"}
	slice2 := []string{"a", "b", "s"}
	slice3 := arrayutil.DisjunctionOrdered(slice, slice2)
	arrayutil.SortOrdered(slice3, true)
	fmt.Println(slice3)

	// Output:
	// [b k n t]
}

func TestSubstract(t *testing.T) {
	Convey("TestSubstract", t, func() {
		arr1 := arrayutil.AsList(1, 2, 3, 4, 5, 6)
		arr2 := arrayutil.AsList(1, 3, 5)
		arr3 := arrayutil.Substract(arr1, arr2, func(s1, s2 int) bool {
			return s1 == s2
		})
		So(len(arr3), ShouldEqual, 3)
		So(arr3, ShouldResemble, []int{2, 4, 6})
	})
}

func ExampleSubstract() {
	slice := []string{"t", "a", "n", "k", "s"}
	slice2 := []string{"a", "b", "s"}
	slice3 := arrayutil.Substract(slice, slice2, func(s1, s2 string) bool { return strings.EqualFold(s1, s2) })
	fmt.Println(slice3)

	// Output:
	// [t n k]
}

func TestSubtractOrdered(t *testing.T) {
	Convey("Test SubtractOrdered", t, func() {
		result := arrayutil.SubstractOrdered(strArray, strArray2)
		So(len(result), ShouldEqual, 1) // hello world matthew xml xml xml xiemalin

		arrayutil.SortOrdered(result, true)

		expected := []string{"matthew"}
		So(result, ShouldResemble, expected)

	})
}

func ExampleSubstractOrdered() {
	slice := []string{"t", "a", "n", "k", "s"}
	slice2 := []string{"a", "b", "s"}
	slice3 := arrayutil.SubstractOrdered(slice, slice2)
	fmt.Println(slice3)

	// Output:
	// [t n k]
}

func TestFilter(t *testing.T) {
	Convey("Test Filter", t, func() {

		Convey("Test Filter struct type", func() {

			result := arrayutil.Filter(persons, func(person Person) bool {
				return person.Age < 100
			})

			So(len(result), ShouldEqual, 1) // hello world matthew xml xml xml xiemalin
			So(result[0].Name, ShouldEqual, N1)
		})

		Convey("Test Filter struct type without match filter condition", func() {

			result := arrayutil.Filter(persons, func(person Person) bool {
				return person.Age > 100
			})

			So(len(result), ShouldEqual, 3) // hello world matthew xml xml xml xiemalin
			So(result[0].Name, ShouldEqual, N1)
		})

		Convey("Test Filter ordered type", func() {

			result := arrayutil.Filter(strArray, func(s string) bool {
				return s != N1
			})

			So(len(result), ShouldEqual, 2) // hello world matthew xml xml xml xiemalin
			So(result[0], ShouldEqual, N1)
		})

	})
}

func ExampleFilter() {
	slice := []string{"t", "a", "n", "k", "s"}
	sl2 := arrayutil.Filter(slice, func(s1 string) bool { return s1 == "a" || s1 == "k" })
	fmt.Println(sl2)

	// Output:
	// [t n s]
}

func TestInsert(t *testing.T) {
	Convey("TestInsert", t, func() {
		sortedIntArray := arrayutil.Clone(sortedIntArray)
		arr := arrayutil.Insert(sortedIntArray, 0, 0)
		So(arr, ShouldResemble, append([]int{0}, sortedIntArray...))

		arr = arrayutil.Insert(sortedIntArray, 5, 0)
		So(arr, ShouldResemble, []int{1, 2, 3, 4, 5, 0, 6, 7, 8, 9, 10})

		arr = arrayutil.Insert(sortedIntArray, 10, 0)
		So(arr, ShouldResemble, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0})

		// invalid index
		arr = arrayutil.Insert(sortedIntArray, -1, 0)
		So(arr, ShouldResemble, sortedIntArray)

		arr = arrayutil.Insert(sortedIntArray, len(sortedIntArray)+1, 0)
		So(arr, ShouldResemble, sortedIntArray)
	})
}

func ExampleInsert() {

	sortedIntArray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// invalid index
	arr := arrayutil.Insert(sortedIntArray, -1, 0)
	fmt.Println(arr)

	arr = arrayutil.Insert(sortedIntArray, len(sortedIntArray)+1, 0)
	fmt.Println(arr)

	// insert by index
	arr = arrayutil.Insert(sortedIntArray, 0, 0) // insert at head
	fmt.Println(arr)

	arr = arrayutil.Insert(sortedIntArray, 5, 0) // insert
	fmt.Println(arr)

	arr = arrayutil.Insert(sortedIntArray, len(sortedIntArray), 0) // insert at tail
	fmt.Println(arr)

	// Output:
	// [1 2 3 4 5 6 7 8 9 10]
	// [1 2 3 4 5 6 7 8 9 10]
	// [0 1 2 3 4 5 6 7 8 9 10]
	// [1 2 3 4 5 0 6 7 8 9 10]
	// [1 2 3 4 5 6 7 8 9 10 0]

}

func TestAsList(t *testing.T) {
	Convey("TestAsList", t, func() {
		// test input nil  will return nil
		arr := arrayutil.AsList[*string]()
		So(arr, ShouldBeNil)

		arr = arrayutil.AsList(conv.ToPtr("hello"), conv.ToPtr("world"))
		So(len(arr), ShouldEqual, 2)
		So(*arr[0], ShouldEqual, "hello")

	})
}

func ExampleAsList() {
	// empty list
	arr := arrayutil.AsList[string]()
	fmt.Println(arr)

	arr = arrayutil.AsList("hello", "world")
	fmt.Println(arr)

	// Output:
	// []
	// [hello world]
}

func TestContainsAny(t *testing.T) {
	Convey("TestContainsAny", t, func() {
		// contains
		arr1 := arrayutil.AsList(1, 3, 5, 7, 9)
		arr2 := arrayutil.AsList(3, 4, 6)
		exist := arrayutil.ContainsAny(arr1, arr2, func(i1, i2 int) bool { return i1 == i2 })
		So(exist, ShouldBeTrue)

		// not contains
		arr2 = arrayutil.AsList(2, 4, 6)
		exist = arrayutil.ContainsAny(arr1, arr2, func(i1, i2 int) bool { return i1 == i2 })
		So(exist, ShouldBeFalse)
	})
}

func ExampleContainsAny() {
	arr1 := arrayutil.AsList(1, 3, 5, 7, 9)
	arr2 := arrayutil.AsList(3, 4, 6)
	exist := arrayutil.ContainsAny(arr1, arr2, func(i1, i2 int) bool { return i1 == i2 })
	fmt.Println(exist)

	// Output:
	// true

}

func TestContainsAnyOrdered(t *testing.T) {
	Convey("TestContainsAny", t, func() {
		// contains
		arr1 := arrayutil.AsList(1, 3, 5, 7, 9)
		arr2 := arrayutil.AsList(3, 4, 6)
		exist := arrayutil.ContainsAnyOrdered(arr1, arr2)
		So(exist, ShouldBeTrue)

		// not contains
		arr2 = arrayutil.AsList(2, 4, 6)
		exist = arrayutil.ContainsAnyOrdered(arr1, arr2)
		So(exist, ShouldBeFalse)
	})
}

func ExampleContainsAnyOrdered() {
	arr1 := arrayutil.AsList(1, 3, 5, 7, 9)
	arr2 := arrayutil.AsList(3, 4, 6)
	exist := arrayutil.ContainsAnyOrdered(arr1, arr2)
	fmt.Println(exist)

	// Output:
	// true

}

func TestRemoves(t *testing.T) {

	Convey("TestRemoves", t, func() {
		arr1 := arrayutil.AsList(3, 5, 7, 9, 3, 5, 1)
		Convey("test remove", func() {
			arr, b := arrayutil.Remove(arr1, 1, func(i1, i2 int) bool { return i1 == i2 })
			So(arrayutil.ContainsOrdered(arr, 1), ShouldBeFalse)
			So(b, ShouldBeTrue)

			arr, b = arrayutil.Remove(arr1, 3, func(i1, i2 int) bool { return i1 == i2 })
			So(arrayutil.ContainsOrdered(arr, 3), ShouldBeTrue)
			So(b, ShouldBeTrue)

			arr, b = arrayutil.RemoveAll(arr1, 5, func(i1, i2 int) bool { return i1 == i2 })
			So(arrayutil.ContainsOrdered(arr, 5), ShouldBeFalse)
			So(b, ShouldBeTrue)
		})

		Convey("test remove ordered", func() {
			arr, b := arrayutil.RemoveOrdered(arr1, 1)
			So(arrayutil.ContainsOrdered(arr, 1), ShouldBeFalse)
			So(b, ShouldBeTrue)

			arr, b = arrayutil.RemoveOrdered(arr1, 3)
			So(arrayutil.ContainsOrdered(arr, 3), ShouldBeTrue)
			So(b, ShouldBeTrue)

			arr, b = arrayutil.RemoveAllOrdered(arr1, 5)
			So(arrayutil.ContainsOrdered(arr, 5), ShouldBeFalse)
			So(b, ShouldBeTrue)
		})

		Convey("test remove index", func() {
			arr := arrayutil.RemoveIndex(arr1, 3)
			So(len(arr), ShouldEqual, 6)
			So(arr, ShouldResemble, []int{3, 5, 7, 3, 5, 1})
		})
	})
}

func ExampleRemove() {
	arr1 := arrayutil.AsList(3, 5, 7, 9, 3, 5, 1)
	arr, b := arrayutil.Remove(arr1, 1, func(i1, i2 int) bool { return i1 == i2 })
	fmt.Println(arr, b)

	arr1 = arrayutil.AsList(3, 5, 7, 9, 3, 5, 1)
	arr, b = arrayutil.Remove(arr1, 5, func(i1, i2 int) bool { return i1 == i2 })
	fmt.Println(arr, b)

	// Output:
	// [3 5 7 9 3 5] true
	// [3 7 9 3 5 1] true
}
func ExampleRemoveAll() {
	arr1 := arrayutil.AsList(3, 5, 7, 9, 3, 5, 1)
	arr, b := arrayutil.RemoveAll(arr1, 1, func(i1, i2 int) bool { return i1 == i2 })
	fmt.Println(arr, b)

	arr1 = arrayutil.AsList(3, 5, 7, 9, 3, 5, 1)
	arr, b = arrayutil.RemoveAll(arr1, 5, func(i1, i2 int) bool { return i1 == i2 })
	fmt.Println(arr, b)

	// Output:
	// [3 5 7 9 3 5] true
	// [3 7 9 3 1] true
}

func TestJoin(t *testing.T) {
	Convey("TestJoin", t, func() {

		intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
		str := arrayutil.Join(intSlice, "-")
		So(str, ShouldEqual, "1-2-3-4-5-6-7-8")

		floatSlice := []float32{1.1, 2.2, 5.5}
		str = arrayutil.Join(floatSlice, ",")
		So(str, ShouldEqual, "1.1,2.2,5.5")

		strSlice := []string{"hello", "world", "!"}
		str = arrayutil.Join(strSlice, " ")
		So(str, ShouldEqual, "hello world !")
	})

}

func ExampleJoin() {
	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	str := arrayutil.Join(intSlice, "-")
	fmt.Println(str)

	// Output:
	// 1-2-3-4-5-6-7-8
}

func TestSwap(t *testing.T) {
	Convey("TestSwap", t, func() {
		Convey("TestSwap1", func() {
			arr := []int{1, 2, 3, 4}
			arrayutil.Swap(arr, 0, 2, 1)
			So(arr, ShouldResemble, []int{3, 2, 1, 4})
		})
		Convey("TestSwap2", func() {
			arr := []int{1, 2, 3, 4}
			arrayutil.Swap(arr, 0, 0, 1)
			So(arr, ShouldResemble, []int{1, 2, 3, 4})
		})
		Convey("TestSwap3", func() {
			arr := []int{1, 2, 3, 4}
			arrayutil.Swap(arr, 2, 0, 2)
			So(arr, ShouldResemble, []int{3, 4, 1, 2})
		})
		Convey("TestSwap4", func() {
			arr := []int{1, 2, 3, 4}
			arrayutil.Swap(arr, -3, 2, 2)
			So(arr, ShouldResemble, []int{3, 4, 1, 2})
		})
		Convey("TestSwap5", func() {
			arr := []int{1, 2, 3, 4}
			arrayutil.Swap(arr, 0, 3, 3)
			So(arr, ShouldResemble, []int{4, 2, 3, 1})
		})
	})

}

func ExampleSwap() {
	arr := []int{1, 2, 3, 4}
	arrayutil.Swap(arr, 0, 2, 1)
	fmt.Println(arr)

	arr = []int{1, 2, 3, 4}
	arrayutil.Swap(arr, 0, 0, 1)
	fmt.Println(arr)

	// Output:
	// [3 2 1 4]
	// [1 2 3 4]
}

func TestAddall(t *testing.T) {
	Convey("TestAddall", t, func() {
		intArray := []int{1, 2}
		array := arrayutil.Addall(intArray, 3, 4)
		So(array, ShouldResemble, []int{1, 2, 3, 4})

		intArray = []int{1, 2}
		array = arrayutil.Addall(intArray)
		So(array, ShouldResemble, []int{1, 2})

		intArray = make([]int, 0)
		array = arrayutil.Addall(intArray, 1, 2)
		So(array, ShouldResemble, []int{1, 2})
	})

}
