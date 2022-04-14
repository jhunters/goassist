/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-02-08 15:41:10
 */
package number_test

import (
	"fmt"
	"testing"

	"github.com/jhunters/goassist/number"
	. "github.com/smartystreets/goconvey/convey"
)

func TestInteger(t *testing.T) {

	Convey("TestInteger", t, func() {

		var v number.Integer = 100
		ProcessInteger(v)

	})
}

func ProcessInteger[E number.Number](e E) {

	ret := e.String()
	So("100", ShouldEqual, ret)

	fmt.Println(e.BinaryString())
}
