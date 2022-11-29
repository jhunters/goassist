/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-01-19 20:38:11
 */
package base_test

import (
	"fmt"
	"testing"

	"github.com/jhunters/goassist/base"
)

type Student struct {
	base.Base
	Name string
	Age  int
}

func TestBase(t *testing.T) {

	student := Student{}
	student.Name = "abc"

	ns := student.Clone()
	s, ok := ns.(Student)
	student.Name = "hello"
	fmt.Println(ok, s.Name)
}
