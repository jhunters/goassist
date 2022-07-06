package mapx_test

import (
	"fmt"
	"testing"

	"github.com/jhunters/goassist/mapx"
)

func TestClone(t *testing.T) {

	m := make(map[string]string)
	m["hello"] = "world"
	m["name"] = "matthew"

	newM := mapx.Clone(m)
	for k, v := range newM {
		fmt.Println(k, v)
	}
}
