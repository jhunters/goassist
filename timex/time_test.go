package timex_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/jhunters/goassist/timex"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	bdate := time.Date(2015, 10, 31, 02, 20, 20, 0, time.Now().Local().Location())
	Convey("Test add days", t, func() {
		dayAfterbdate := time.Date(2015, 11, 01, 02, 20, 20, 0, time.Now().Local().Location())
		retDate := timex.AddDays(bdate, 1)
		So(retDate, ShouldEqual, dayAfterbdate)
	})
	Convey("Test add hours", t, func() {
		dayAfterbdate := time.Date(2015, 10, 31, 03, 20, 20, 0, time.Now().Local().Location())
		retDate := timex.AddHours(bdate, 1)
		So(retDate, ShouldEqual, dayAfterbdate)
	})
	Convey("Test add minutes", t, func() {
		dayAfterbdate := time.Date(2015, 10, 31, 02, 21, 20, 0, time.Now().Local().Location())
		retDate := timex.AddMinutes(bdate, 1)
		So(retDate, ShouldEqual, dayAfterbdate)
	})
}

func TestTruncate(t *testing.T) {
	Convey("TestTruncate", t, func() {
		bdate := time.Date(2015, 10, 31, 02, 20, 20, 0, time.Now().Local().Location())

		Convey("test truncate year", func() {
			retDate := timex.TruncateYear(bdate)
			So(retDate, ShouldEqual, time.Date(2015, 0, 0, 0, 0, 0, 0, bdate.Location()))
		})

		Convey("test truncate month", func() {
			retDate := timex.TruncateMonth(bdate)
			So(retDate, ShouldEqual, time.Date(2015, 10, 0, 0, 0, 0, 0, bdate.Location()))
		})

		Convey("test truncate day", func() {
			retDate := timex.TruncateDay(bdate)
			So(retDate, ShouldEqual, time.Date(2015, 10, 31, 0, 0, 0, 0, bdate.Location()))
		})

		Convey("test truncate hour", func() {
			retDate := timex.TruncateHour(bdate)
			So(retDate, ShouldEqual, time.Date(2015, 10, 31, 02, 0, 0, 0, bdate.Location()))
		})

		Convey("test truncate minute", func() {
			retDate := timex.TruncateMinute(bdate)
			So(retDate, ShouldEqual, time.Date(2015, 10, 31, 02, 20, 0, 0, bdate.Location()))
		})

		Convey("test truncate second", func() {
			retDate := timex.TruncateSecond(bdate)
			So(retDate, ShouldEqual, time.Date(2015, 10, 31, 02, 20, 20, 0, bdate.Location()))
		})

	})

}

func TestParseTime(t *testing.T) {
	Convey("TestParseTime", t, func() {

		Convey("test parse time by simple format", func() {
			bdate := time.Date(2015, 10, 31, 0, 0, 0, 0, time.UTC)
			tm, err := timex.ParseSimpleFormat("2015-10-31")
			So(err, ShouldBeNil)
			So(bdate, ShouldEqual, tm)
		})

		Convey("test parse time by format", func() {
			bdate := time.Date(2015, 10, 31, 02, 20, 20, 0, time.UTC)
			tm, err := timex.ParseFormat("2015-10-31 02:20:20")
			So(err, ShouldBeNil)
			So(bdate, ShouldEqual, tm)
		})

	})
}

func TestFormat(t *testing.T) {
	Convey("TestFormat", t, func() {

		Convey("test simple format", func() {
			bdate := time.Date(2015, 10, 31, 0, 0, 0, 0, time.UTC)
			sdate := timex.FormatSimple(bdate)
			So(sdate, ShouldEqual, "2015-10-31")
		})

		Convey("test format", func() {
			bdate := time.Date(2015, 10, 31, 02, 20, 20, 0, time.UTC)
			sdate := timex.Format(bdate)
			So(sdate, ShouldEqual, "2015-10-31 02:20:20")
		})

	})
}

func TestTimers(t *testing.T) {

	ti := time.NewTimer(1 * time.Second)

	fmt.Println(<-ti.C)

	ti.Stop()

}
