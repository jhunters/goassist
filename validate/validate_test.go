package validate_test

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"testing"

	"github.com/jhunters/goassist/stringutil"
	"github.com/jhunters/goassist/validate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestValidateEmail(t *testing.T) {

	Convey("TestValidateEmail", t, func() {

		Convey("validate email success", func() {

			qqmail := "191988332@qq.com"
			isValid := validate.ValidateEmail(qqmail)
			So(isValid, ShouldBeTrue)

			hotmail := "happy_newyear@hotmail.com"
			isValid = validate.ValidateEmail(hotmail)
			So(isValid, ShouldBeTrue)

		})

		Convey("validate email fail", func() {
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

	})
}

func TestXxx(t *testing.T) {

	s := "第一百四十九章 尤梦清落败"

	reg, _ := regexp.Compile(`(第[\W]+章)\s+([\w|\W]+)`)
	b := reg.MatchString(s)
	fmt.Println(b)

	result := reg.FindAllStringSubmatch(s, -1)
	if len(result) == 1 && len(result[0]) == 2 {
		fmt.Println(result[0][0], result[0][1])
	}

}

func TestGetPrefaceAndName(t *testing.T) {

	reg, _ := regexp.Compile(`(第[\d]+章)\s+([\w|\W]+)`)

	c, err := ioutil.ReadFile(`D:\documents\medicine.txt`)
	if err != nil {
		t.Error(err)
		return
	}

	lines := strings.Split(string(c), "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if reg.MatchString(line) {
			if strings.Contains(line, "（") {
				line = stringutil.SubstringBefore(line, "（")
			}
			fmt.Println(line)
		}
	}
}
