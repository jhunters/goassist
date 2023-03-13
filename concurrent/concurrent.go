package concurrent

import (
	"fmt"
	"time"

	"github.com/jhunters/goassist/base"
	"github.com/jhunters/timewheel"
)

const (
	TIME_OUT_TW     = 1
	TIME_OUT_TICKER = 2
)

var (
	tw, _   = createTimeWheel(60, 50*time.Millisecond)
	EmptyFn = func() {}

	TimeOutType = TIME_OUT_TW
)

func ReSetTimeWheel(slotNum uint16, interval time.Duration) error {
	if tw != nil {
		tw.Stop()
	}
	tw, err := createTimeWheel(slotNum, interval)
	tw.Start()
	return err
}

func timeoutTW(timeout time.Duration) (<-chan time.Time, func()) {
	task, tout := newTask()
	tw.AddTask(timeout, task)
	return tout, EmptyFn
}

func timeoutTicker(timeout time.Duration) (<-chan time.Time, func()) {
	tick := time.NewTicker(timeout)
	return tick.C, func() {
		tick.Stop()
	}
}

func timeoutF(timeout time.Duration) (<-chan time.Time, func()) {
	if TimeOutType == TIME_OUT_TW {
		return timeoutTW(timeout)
	}
	return timeoutTicker(timeout)
}

// AsyncGo execute target function by goroutine. if panic happened will wrap error object and return false
// if just time out will return ok(false), err(nil)
func AsyncGo(f base.Call, timeout time.Duration) (ok bool, err error) {
	if timeout <= 0 { // no timeout need
		ok = true
		return
	}
	tout, cancel := timeoutF(timeout)
	defer cancel()

	return AsyncGoWithEvent(f, tout)

}

// AsyncGo execute target function by goroutine. if panic happened will wrap error object and return false
// if just time out will return ok(false), err(nil)
func AsyncGoWithEvent[E any](f base.Call, toevent <-chan E) (ok bool, err error) {
	ch := make(chan error, 1)
	go func(ch chan<- error) {
		defer panicCatch(ch)
		f() // do process
		close(ch)
	}(ch)

	select {
	case e := <-ch:
		ok = (e == nil)
		err = e
	case <-toevent:
		ok = false
	}
	return

}

// AsyncCall execute target function by goroutine and has a generic returned parameter. if panic happened will wrap error object and return future(nil)
// if just time out will return future(func), err(nil)
func AsyncCall[E any](f base.Supplier[E], timeout time.Duration) (future base.Supplier[E], err error) {

	if timeout <= 0 { // no timeout need
		future = func() E {
			return f()
		}
		return
	}

	tout, cancel := timeoutF(timeout)
	defer cancel()
	return AsyncCallWithEvent(f, tout)
}

// AsyncCall execute target function by goroutine and has a generic returned parameter. if panic happened will wrap error object and return future(nil)
// if just time out will return future(func), err(nil)
func AsyncCallWithEvent[E, T any](f base.Supplier[E], toevent <-chan T) (future base.Supplier[E], err error) {
	ret := make(chan E, 1)
	future = func() E {
		return <-ret
	}
	ch := make(chan error, 1)
	go func(ch chan<- error) {
		defer panicCatch(ch)
		e := f()
		ret <- e
		close(ch)
	}(ch)

	select {
	case e := <-ch:
		err = e
		if e != nil {
			future = nil
		}
	case <-toevent:
		err = fmt.Errorf("AsyncCall execute timeout")
	}
	return
}

func panicCatch(ch chan<- error) {
	if v := recover(); v != nil {
		e, ok := v.(error)
		if ok {
			ch <- e
		} else {
			ch <- fmt.Errorf("%v", v)
		}
	}
}

func createTimeWheel(slotNum uint16, interval time.Duration) (*timewheel.TimeWheel, error) {
	tw, err := timewheel.New(interval, slotNum)
	tw.Start()
	return tw, err
}

func newTask() (timewheel.Task, <-chan time.Time) {
	ch := make(chan time.Time, 1)
	tt := timewheel.Task{
		TimeoutCallback: func(task timewheel.Task) { // call back function on time out
			ch <- time.Now()
		}}
	return tt, ch
}
