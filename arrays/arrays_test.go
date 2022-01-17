/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2021-12-24 13:55:47
 */
package arrays_test

import (
	"fmt"
	"testing"

	"github.com/jhunters/goassist/arrays"

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

		persons := arrays.Clone(persons)
		Convey("Test Person struct array desc order", func() {
			arrays.Sort(persons, func(e1, e2 Person) int {
				return int(e1.Age) - int(e2.Age)
			})
			So(persons[0].Name, ShouldEqual, "xml")
			So(len(persons), ShouldEqual, 3)
		})

		Convey("Test Person struct array asc order", func() {
			arrays.Sort(persons, func(e1, e2 Person) int {
				return int(e2.Age) - int(e1.Age)
			})
			So(persons[0].Name, ShouldEqual, "matthew")
			So(len(persons), ShouldEqual, 3)
		})
	})

	Convey("Test sort builtin type array", t, func() {
		strArray := []string{"xml", "matthew", "xiemalin"}
		arrays.Sort(strArray, func(e1, e2 string) int {
			return arrays.CompareTo(e1, e2)
		})
		So(strArray[0], ShouldEqual, "xml")
	})
}

func ExampleSort() {
	// order by person's age
	persons = []Person{{"xml", 100}, {"matthew", 90}, {"xiemalin", 99}}
	arrays.Sort(persons, func(e1, e2 Person) int {
		return int(e1.Age) - int(e2.Age)
	})
	fmt.Println(persons)
	//output:
	//[{xml 100} {xiemalin 99} {matthew 90}]
}

func TestSortOrdered(t *testing.T) {
	Convey("Test SortOrdered", t, func() {
		strArray := arrays.Clone(strArray)
		arrays.SortOrdered(strArray, false)
		So(strArray[0], ShouldEqual, "xml")

		arrays.SortOrdered(strArray, true)
		So(strArray[0], ShouldEqual, "matthew")
	})
}

func ExampleSortOrdered() {
	strArray := arrays.Clone(strArray)
	arrays.SortOrdered(strArray, false)
	fmt.Println(strArray)

	//output:
	//[xml xml xiemalin matthew]
}

func TestShuffle(t *testing.T) {
	Convey("Test Shuffle", t, func() {

		sortedIntArray := arrays.Clone(sortedIntArray)
		arrays.Shuffle(sortedIntArray)

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

		sortedIntArray := arrays.Clone(sortedIntArray)
		arrays.Reverse(sortedIntArray)

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

func TestBinarySearch(t *testing.T) {
	Convey("Test BinarySearch", t, func() {

		persons := arrays.Clone(sortedPersons)

		cmp := func(e1, e2 Person) int {
			return arrays.CompareTo(e1.Age, e2.Age)
		}

		Convey("Test BinarySearch exist", func() {
			key := Person{"matthew", 90}

			offset := arrays.BinarySearch(persons, key, cmp)
			So(offset, ShouldEqual, 0)

			key = Person{"matthew", 96}
			offset = arrays.BinarySearch(persons, key, cmp)
			So(offset, ShouldEqual, 6)

			key = Person{"matthew", 92}
			offset = arrays.BinarySearch(persons, key, cmp)
			So(offset, ShouldEqual, 2)
		})

		Convey("Test BinarySearch no exist", func() {

			key := Person{"matthew", 89}
			offset := arrays.BinarySearch(persons, key, cmp)
			So(offset < 0, ShouldBeTrue)

			key = Person{"matthew", 100}
			offset = arrays.BinarySearch(persons, key, cmp)
			So(offset < 0, ShouldBeTrue)
		})

	})
}

func ExampleBinarySearch() {
	sortedPersons = []Person{{"xml", 90}, {"matthew", 91}, {"xiemalin", 92}, {"xiemalin2", 93}, {"xiemalin3", 94}, {"xiemalin4", 95}, {"xiemalin5", 96}}
	key := Person{"matthew", 90}
	offset := arrays.BinarySearch(sortedPersons, key, func(e1, e2 Person) int {
		return arrays.CompareTo(e1.Age, e2.Age)
	})

	fmt.Println(offset, sortedPersons[offset].Age)
	//output:
	//0 90
}

func TestBinarySearchOrdered(t *testing.T) {
	Convey("Test BinarySearch", t, func() {
		sortedIntArray := arrays.Clone(sortedIntArray)

		cmp := func(e1, e2 int) int {
			return e1 - e2
		}
		Convey("Test BinarySearch exist", func() {

			offset := arrays.BinarySearchOrdered(sortedIntArray, 1)
			So(offset, ShouldEqual, 0)

			offset = arrays.BinarySearch(sortedIntArray, 10, cmp)
			So(offset, ShouldEqual, 9)

			offset = arrays.BinarySearchOrdered(sortedIntArray, 5)
			So(offset, ShouldEqual, 4)
		})

		Convey("Test BinarySearch no exist", func() {

			offset := arrays.BinarySearchOrdered(sortedIntArray, 0)
			So(offset < 0, ShouldBeTrue)

			offset = arrays.BinarySearchOrdered(sortedIntArray, 11)
			So(offset < 0, ShouldBeTrue)
		})

	})
}

func TestMax(t *testing.T) {
	Convey("Test Max", t, func() {

		persons := arrays.Clone(persons)
		Oldest, pos := arrays.Max(persons, func(e1, e2 Person) int {
			return int(e1.Age) - int(e2.Age)
		})
		So(Oldest.Age, ShouldEqual, 100)
		So(pos, ShouldEqual, 0)

		sortedPersons := arrays.Clone((sortedPersons))
		Oldest, pos = arrays.Max(sortedPersons, func(e1, e2 Person) int {
			return int(e1.Age) - int(e2.Age)
		})
		So(Oldest.Age, ShouldEqual, 96)
		So(pos, ShouldEqual, len(sortedPersons)-1)

		oneEle := []Person{{"hello", -100}}
		one, pos := arrays.Max(oneEle, func(e1, e2 Person) int {
			return int(e1.Age) - int(e2.Age)
		})
		So(one.Age, ShouldEqual, -100)
		So(pos, ShouldEqual, 0)
	})
}

func TestMaxOrdered(t *testing.T) {
	Convey("Test MaxOrdered", t, func() {

		strArray := arrays.Clone(strArray)
		biggestStr, pos := arrays.MaxOrdered(strArray)
		So(biggestStr, ShouldEqual, "xml")
		So(pos, ShouldEqual, 0)

		sortedIntArray := arrays.Clone(sortedIntArray)
		biggestInt, pos := arrays.MaxOrdered(sortedIntArray)
		So(biggestInt, ShouldEqual, 10)
		So(pos, ShouldEqual, len(sortedIntArray)-1)
	})
}

func TestMin(t *testing.T) {
	Convey("Test Min", t, func() {

		persons := arrays.Clone(persons)
		Oldest, pos := arrays.Min(persons, func(e1, e2 Person) int {
			return int(e1.Age) - int(e2.Age)
		})
		So(Oldest.Age, ShouldEqual, 90)
		So(pos, ShouldEqual, 1)

		sortedPersons := arrays.Clone((sortedPersons))
		Oldest, pos = arrays.Min(sortedPersons, func(e1, e2 Person) int {
			return int(e1.Age) - int(e2.Age)
		})
		So(Oldest.Age, ShouldEqual, 90)
		So(pos, ShouldEqual, 0)

		oneEle := []Person{{"hello", -100}}
		one, pos := arrays.Min(oneEle, func(e1, e2 Person) int {
			return int(e1.Age) - int(e2.Age)
		})
		So(one.Age, ShouldEqual, -100)
		So(pos, ShouldEqual, 0)
	})
}

func TestMinOrdered(t *testing.T) {
	Convey("Test MinOrdered", t, func() {

		strArray := arrays.Clone(strArray)
		biggestStr, pos := arrays.MinOrdered(strArray)
		So(biggestStr, ShouldEqual, "matthew")
		So(pos, ShouldEqual, 1)

		sortedIntArray := arrays.Clone(sortedIntArray)
		biggestInt, pos := arrays.MinOrdered(sortedIntArray)
		So(biggestInt, ShouldEqual, 1)
		So(pos, ShouldEqual, 0)
	})
}

func TestRelaceAll(t *testing.T) {
	Convey("Test RelaceOrderedAll", t, func() {

		persons := arrays.Clone(persons)
		manReplaceArray := arrays.Clone(persons)
		manReplaceArray[1] = Person{"xiemalin", 100}
		oldPerson := Person{"matthew", 90}
		newPerson := Person{"xiemalin", 100}
		arrays.ReplaceAll(persons, oldPerson, newPerson, func(e1, e2 Person) bool {
			return e1.Equal(e2)
		})

		So(persons, ShouldResemble, manReplaceArray)

		// replace not exist
		notexistPerson := Person{"michael", 10}
		arrays.ReplaceAll(persons, notexistPerson, newPerson, func(e1, e2 Person) bool {
			return e1.Equal(e2)
		})
		So(persons, ShouldResemble, manReplaceArray)
	})
}

func TestRelaceOrderedAll(t *testing.T) {
	Convey("Test RelaceOrderedAll", t, func() {

		strArray := arrays.Clone(strArray)
		manReplaceArray := arrays.Clone(strArray)
		manReplaceArray[1] = "xiemalin"
		arrays.ReplaceOrderedAll(strArray, "matthew", "xiemalin")

		So(len(strArray), ShouldEqual, 4)
		So(strArray, ShouldResemble, manReplaceArray)

		// replace not exist
		arrays.ReplaceOrderedAll(strArray, "matthew", "xml")

	})
}

func TestCreateAndFill(t *testing.T) {
	Convey("Test CreateAndFill", t, func() {
		strArray := arrays.CreateAndFill(10, "name")
		So(len(strArray), ShouldEqual, 10)
		So(strArray[0], ShouldEqual, "name")
		So(strArray[4], ShouldEqual, "name")
		So(strArray[9], ShouldEqual, "name")

		Convey("Test CreateAndFill with zero size", func() {
			strArray := arrays.CreateAndFill(0, "name")
			So(len(strArray), ShouldEqual, 0)
		})

	})
}

func TestIndexOfSubArray(t *testing.T) {
	Convey("Test IndexOfSubArray", t, func() {
		Convey("Test IndexOfSubArray with not found", func() {

			persons := arrays.Clone(persons)
			subPersons := []Person{}

			pos := arrays.IndexOfSubArray(persons, subPersons, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, -1)

			// the length of sub array is large than source
			subPersons2 := append(persons, Person{})
			pos = arrays.IndexOfSubArray(persons, subPersons2, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, -1)
		})

		Convey("Test IndexOfSubArray start with sub array", func() {
			// with same array
			persons := arrays.Clone(persons)
			pos := arrays.IndexOfSubArray(persons, persons, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, 0)

			// with sub pos at 0
			persons = arrays.Clone(persons)
			sub := arrays.Clone(persons)

			persons = append(persons, sub...)
			pos = arrays.IndexOfSubArray(persons, persons, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, 0)

		})

		Convey("Test IndexOfSubArray with has sub array", func() {
			// with sub pos at some position
			newPersons := arrays.Clone(sortedPersons)
			newPersons = append(newPersons, persons...)
			pos := arrays.IndexOfSubArray(newPersons, persons, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, len(sortedPersons))
		})

		Convey("Test IndexOfSubArray with has two same sub array", func() {
			// with sub pos at some position
			newPersons := arrays.Clone(sortedPersons)
			newPersons = append(newPersons, persons...)
			newPersons = append(newPersons, persons...)
			pos := arrays.IndexOfSubArray(newPersons, persons, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, len(sortedPersons))
		})

	})
}

func TestIndexOfSubOrderedArray(t *testing.T) {
	Convey("Test IndexOfSubOrderedArray", t, func() {
		strArray := arrays.Clone(strArray)
		index := arrays.IndexOfSubOrderedArray(strArray, []string{"xml"})
		So(index, ShouldEqual, 0)
	})
}

func TestLastIndexOfSubOrderedArray(t *testing.T) {
	Convey("Test LastIndexOfSubOrderedArray", t, func() {
		strArray := arrays.Clone(strArray)
		index := arrays.LastIndexOfSubOrderedArray(strArray, []string{"xml"})
		So(index, ShouldEqual, 3)
	})
}

func TestLastIndexOfSubArray(t *testing.T) {
	Convey("Test LastIndexOfSubArray", t, func() {
		Convey("Test LastIndexOfSubArray with not found", func() {

			persons := arrays.Clone(persons)
			subPersons := []Person{}

			pos := arrays.LastIndexOfSubArray(persons, subPersons, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, -1)

			// the length of sub array is large than source
			subPersons2 := append(persons, Person{})
			pos = arrays.LastIndexOfSubArray(persons, subPersons2, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, -1)
		})

		Convey("Test LastIndexOfSubArray start with sub array", func() {
			// with same array
			persons := arrays.Clone(persons)
			pos := arrays.LastIndexOfSubArray(persons, persons, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, 0)

			// with sub pos at 0
			persons = arrays.Clone(persons)
			sub := arrays.Clone(persons)

			persons = append(persons, sub...)
			pos = arrays.LastIndexOfSubArray(persons, persons, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, 0)

		})

		Convey("Test LastIndexOfSubArray with has sub array", func() {
			// with sub pos at some position
			newPersons := arrays.Clone(sortedPersons)
			newPersons = append(newPersons, persons...)
			pos := arrays.LastIndexOfSubArray(newPersons, persons, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, len(sortedPersons))
		})

		Convey("Test LastIndexOfSubArray with has two same sub array", func() {
			// with sub pos at some position
			newPersons := arrays.Clone(sortedPersons)
			newPersons = append(newPersons, persons...)
			newPersons = append(newPersons, persons...)
			pos := arrays.LastIndexOfSubArray(newPersons, persons, func(e1, e2 Person) bool {
				return e1.Equal(e2)
			})
			So(pos, ShouldEqual, len(sortedPersons)+len(persons))
		})

	})
}

func TestDisJoint(t *testing.T) {
	Convey("Test DisJoint", t, func() {

		hasNoSame := arrays.Disjoint(persons, persons2, func(e1, e2 Person) bool {
			return e1.Equal(e2)
		})
		So(hasNoSame, ShouldBeFalse)

		otherPersons := []Person{}
		hasNoSame = arrays.Disjoint(persons, otherPersons, func(e1, e2 Person) bool {
			return e1.Equal(e2)
		})
		So(hasNoSame, ShouldBeTrue)

		hasNoSame = arrays.Disjoint(otherPersons, persons, func(e1, e2 Person) bool {
			return e1.Equal(e2)
		})
		So(hasNoSame, ShouldBeTrue)

	})
}

func TestDisJointOrdered(t *testing.T) {
	Convey("Test DisjointOrdered", t, func() {
		hasNoSame := arrays.DisjointOrdered(strArray, strArray)
		So(hasNoSame, ShouldBeFalse)

	})
}

func TestRotate(t *testing.T) {
	Convey("Test Rotate", t, func() {

		expected := []string{"s", "t", "a", "n", "k"}
		arr := []string{"t", "a", "n", "k", "s"}
		arrays.Rotate(arr, 1)
		So(arr, ShouldResemble, expected)
		arr = []string{"t", "a", "n", "k", "s"}
		arrays.Rotate(arr, -4)
		So(arr, ShouldResemble, expected)

		expected2 := []string{"a", "n", "k", "s", "t"}
		arr = []string{"t", "a", "n", "k", "s"}
		arrays.Rotate(arr, 4)
		So(arr, ShouldResemble, expected2)
		arr = []string{"t", "a", "n", "k", "s"}
		arrays.Rotate(arr, -1)
		So(arr, ShouldResemble, expected2)
	})
}

func TestContainsOrdered(t *testing.T) {
	Convey("Test ContainsOrdered", t, func() {
		strArray := arrays.Clone(strArray)

		contains := arrays.ContainsOrdered(strArray, "xml")
		So(contains, ShouldBeTrue)

		contains = arrays.ContainsOrdered(strArray, "notexist")
		So(contains, ShouldBeFalse)
	})
}

func TestEqualWithOrdered(t *testing.T) {
	Convey("Test EqualWithOrdered", t, func() {
		strArray := arrays.Clone(strArray)
		emptyArray := []string{}

		equals := arrays.EqualWithOrdered(strArray, emptyArray)
		So(equals, ShouldBeFalse)

		equals = arrays.EqualWithOrdered(strArray, strArray)
		So(equals, ShouldBeTrue)
	})
}

func TestIntersectionOrdered(t *testing.T) {
	Convey("Test IntersectionOrdered", t, func() {
		result := arrays.IntersectionOrdered(strArray, strArray2)
		So(len(result), ShouldEqual, 3) // xml xml xiemalin

		arrays.SortOrdered(result, true)
		expected := []string{"xiemalin", "xml", "xml"}
		So(result, ShouldResemble, expected)
	})
}

func TestUnionOrdered(t *testing.T) {
	Convey("Test UnionOrdered", t, func() {
		result := arrays.UnionOrdered(strArray, strArray2)
		So(len(result), ShouldEqual, 7) // hello world matthew xml xml xml xiemalin

		arrays.SortOrdered(result, true)

		expected := []string{"hello", "matthew", "world", "xiemalin", "xml", "xml", "xml"}
		So(result, ShouldResemble, expected)

	})
}

func TestDisjunctionOrdered(t *testing.T) {
	Convey("Test DisjunctionOrdered", t, func() {
		result := arrays.DisjunctionOrdered(strArray, strArray2)
		So(len(result), ShouldEqual, 4) // hello world matthew xml xml xml xiemalin

		arrays.SortOrdered(result, true)

		expected := []string{"hello", "matthew", "world", "xml"}
		So(result, ShouldResemble, expected)

	})
}

func TestSubtractOrdered(t *testing.T) {
	Convey("Test SubtractOrdered", t, func() {
		result := arrays.SubtractOrdered(strArray, strArray2)
		So(len(result), ShouldEqual, 1) // hello world matthew xml xml xml xiemalin

		arrays.SortOrdered(result, true)

		expected := []string{"matthew"}
		So(result, ShouldResemble, expected)

	})
}

func TestFilter(t *testing.T) {
	Convey("Test Filter", t, func() {

		Convey("Test Filter struct type", func() {

			result := arrays.Filter(persons, func(person Person) bool {
				return person.Age < 100
			})

			So(len(result), ShouldEqual, 1) // hello world matthew xml xml xml xiemalin
			So(result[0].Name, ShouldEqual, N1)
		})

		Convey("Test Filter struct type without match filter condition", func() {

			result := arrays.Filter(persons, func(person Person) bool {
				return person.Age > 100
			})

			So(len(result), ShouldEqual, 3) // hello world matthew xml xml xml xiemalin
			So(result[0].Name, ShouldEqual, N1)
		})

		Convey("Test Filter ordered type", func() {

			result := arrays.Filter(strArray, func(s string) bool {
				return s != N1
			})

			So(len(result), ShouldEqual, 2) // hello world matthew xml xml xml xiemalin
			So(result[0], ShouldEqual, N1)
		})

	})
}
