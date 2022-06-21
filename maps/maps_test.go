package maps_test

import (
	"fmt"
	"testing"

	"github.com/jhunters/goassist/maps"
)

func TestClone(t *testing.T) {

	m := make(map[string]string)
	m["hello"] = "world"
	m["name"] = "matthew"

	newM := maps.Clone(m)
	for k, v := range newM {
		fmt.Println(k, v)
	}
}
