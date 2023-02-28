package validate_test

import (
	"testing"

	"github.com/jhunters/goassist/validate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestValidateEmail(t *testing.T) {

	Convey("TestValidateEmail", t, func() {

		Convey("validate email with correct format", func() {

			qqmail := "191988332@qq.com"
			isValid := validate.ValidateEmail(qqmail)
			So(isValid, ShouldBeTrue)

			hotmail := "happy_newyear@hotmail.com"
			isValid = validate.ValidateEmail(hotmail)
			So(isValid, ShouldBeTrue)

		})

		Convey("validate email rfc with correct format", func() {

			qqmail := "191988332@qq.com"
			isValid := validate.ValidateEmailRFC(qqmail)
			So(isValid, ShouldBeTrue)

			hotmail := "happy_newyear@hotmail.com"
			isValid = validate.ValidateEmailRFC(hotmail)
			So(isValid, ShouldBeTrue)

		})

		Convey("validate email with bad format", func() {
			noName := "@qq.com"
			isValid := validate.ValidateEmail(noName)
			So(isValid, ShouldBeFalse)

			noDomain := "191988332@.com"
			isValid = validate.ValidateEmail(noDomain)
			So(isValid, ShouldBeFalse)

			noDomain2 := "191988332@qq"
			isValid = validate.ValidateEmail(noDomain2)
			So(isValid, ShouldBeFalse)
		})

		Convey("validate email rfc with bad format", func() {
			noName := "@qq.com"
			isValid := validate.ValidateEmailRFC(noName)
			So(isValid, ShouldBeFalse)

			noDomain := "191988332@.com"
			isValid = validate.ValidateEmailRFC(noDomain)
			So(isValid, ShouldBeFalse)

			noDomain2 := "191988332@qq"
			isValid = validate.ValidateEmailRFC(noDomain2)
			So(isValid, ShouldBeFalse)
		})

	})
}
