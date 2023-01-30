package concurrent_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/jhunters/goassist/arrayutil"
	"github.com/jhunters/goassist/concurrent"
	. "github.com/smartystreets/goconvey/convey"
)

var var_s string = "hello"

var var_pojo Pojo = Pojo{"matt"}

type Pojo struct {
	Name string
}

func timedFunc[E any](t time.Duration, ret E) E {
	time.Sleep(t)
	return ret
}

func panicFunc() {
	panic("throws panic manually")
}

func panicFunc2[E any](ret E) E {
	panic("throws panic manually")
}

func TestAsyncGo(t *testing.T) {

	Convey("TestAsyncGo in time", t, func() {
		now := time.Now()
		ch := make(chan string, 1)
		f, err := concurrent.AsyncGo(func() {
			ch <- timedFunc(100*time.Millisecond, var_s)
		}, time.Second)

		took := time.Now().Sub(now)
		So(f, ShouldBeTrue)
		So(err, ShouldBeNil)
		So(<-ch, ShouldEqual, var_s)
		So(took, ShouldBeLessThanOrEqualTo, time.Second)
		So(took, ShouldBeGreaterThan, 100*time.Millisecond)
	})

	Convey("TestAsyncGo time out", t, func() {
		now := time.Now()
		ch := make(chan Pojo, 1)
		f, err := concurrent.AsyncGo(func() {
			ch <- timedFunc(2*time.Second, var_pojo)
		}, time.Second)

		took := time.Now().Sub(now)
		So(f, ShouldBeFalse)
		So(err, ShouldBeNil)
		So(took, ShouldBeGreaterThan, 900*time.Millisecond)
		So(took, ShouldBeLessThan, 2*time.Second)
		So(<-ch, ShouldResemble, var_pojo)
		took = time.Now().Sub(now)
		So(took, ShouldBeGreaterThan, 2*time.Second)
	})

}

func ExampleAsyncGo() {
	// run function in async way
	f, err := concurrent.AsyncGo(func() {
		time.Sleep(200 * time.Millisecond)
	}, time.Second)
	fmt.Println(f, err)

	// run function in async way and ocurres timeout
	f, err = concurrent.AsyncGo(func() {
		time.Sleep(2 * time.Second)
	}, time.Second)
	fmt.Println(f, err)

	// run function in async way and panic error
	f, err = concurrent.AsyncGo(func() {
		panic("surprise")
	}, time.Second)
	fmt.Println(f, err.Error())

	// Output:
	// true <nil>
	// false <nil>
	// false surprise
}

func TestAsyncCall(t *testing.T) {

	Convey("TestAsyncCall in time", t, func() {
		now := time.Now()
		f, err := concurrent.AsyncCall(func() string {
			return timedFunc(100*time.Millisecond, var_s)
		}, time.Second)

		took := time.Now().Sub(now)
		So(f, ShouldNotBeNil)
		So(err, ShouldBeNil)
		So(f(), ShouldEqual, var_s)
		So(took, ShouldBeLessThanOrEqualTo, time.Second)
		So(took, ShouldBeGreaterThan, 100*time.Millisecond)
	})

	Convey("TestAsyncCall time out", t, func() {
		now := time.Now()
		f, err := concurrent.AsyncCall(func() Pojo {
			return timedFunc(2*time.Second, var_pojo)
		}, time.Second)

		took := time.Now().Sub(now)
		So(f, ShouldNotBeNil)
		So(err, ShouldNotBeNil)
		So(took, ShouldBeGreaterThan, 1*time.Second)
		So(took, ShouldBeLessThan, 2*time.Second)
		So(f(), ShouldResemble, var_pojo)
		took = time.Now().Sub(now)
		So(took, ShouldBeGreaterThan, 2*time.Second)
	})

}

func ExampleAsyncCall() {

	get := func() (name, address string) {
		return "matt", "pudong"
	}

	// run call function in async way
	f, err := concurrent.AsyncCall(func() []string {
		time.Sleep(200 * time.Millisecond)
		name, address := get()
		return arrayutil.AsList(name, address)

	}, time.Second)
	fmt.Println(f(), err)

	// run call function in async way and ocurres timeout
	f, err = concurrent.AsyncCall(func() []string {
		time.Sleep(2 * time.Second)
		name, address := get()
		return arrayutil.AsList(name, address)

	}, time.Second)
	fmt.Println(f(), err)

	// run call function in async way without time wait
	f, err = concurrent.AsyncCall(func() []string {
		time.Sleep(2 * time.Second)
		name, address := get()
		return arrayutil.AsList(name, address)

	}, 0)
	fmt.Println(f(), err)

	// Output:
	// [matt pudong] <nil>
	// [matt pudong] AsyncCall execute timeout. expect 1s
	// [matt pudong] <nil>

}

func TestAsyncGoWithPanic(t *testing.T) {
	Convey("TestAsyncCallWithPanic in time", t, func() {
		b, err := concurrent.AsyncGo(func() {
			panicFunc()
		}, time.Second)
		So(b, ShouldBeFalse)
		So(err, ShouldNotBeNil)
	})
}
func TestAsyncCallWithPanic(t *testing.T) {
	Convey("TestAsyncCallWithPanic in time", t, func() {
		b, err := concurrent.AsyncCall(func() string {
			return panicFunc2(var_s)
		}, time.Second)
		So(b, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})
}
