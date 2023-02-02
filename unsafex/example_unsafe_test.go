package unsafex_test

import (
	"fmt"

	"github.com/jhunters/goassist/unsafex"
)

type Data struct {
	V1 int
	V2 int
}

func ExampleMappingToArray() {

	d1 := Data{1, 2}
	bytes := unsafex.MappingToArray(d1)
	fmt.Println(bytes)

	dr := []Data{{1, 1}, {2, 2}}
	bytes = unsafex.MappingToArray(dr)
	fmt.Println(bytes)

	// Output:
	// [1 0 0 0 0 0 0 0 2 0 0 0 0 0 0 0]
	// [128 161 9 0 192 0 0 0 2 0 0 0 0 0 0 0 2 0 0 0 0 0 0 0]
}

func ExampleArrayMapping() {

	bytes := []byte{1, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0}
	var d1 *Data = unsafex.ArrayMapping[Data](bytes)
	fmt.Println(d1.V1, d1.V2)

	dr := []Data{{1, 1}, {2, 2}}
	bytes = unsafex.MappingToArray(dr)
	v := unsafex.ArrayMapping[[]Data](bytes)
	fmt.Println(*v)

	// Output:
	// 1 2
	// [{1 1} {2 2}]
}
