package concurrent

import (
	"fmt"
	"time"
)

// AsyncGo execute target function by goroutine. if panic happened will wrap error object and return false
// if just time out will return ok(false), err(nil)
func AsyncGo(f func(), timeout time.Duration) (ok bool, err error) {
	tick := time.NewTicker(timeout)
	defer tick.Stop()
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
	case <-tick.C:
		ok = false
	}
	return
}

// AsyncCall execute target function by goroutine and has a generic returned parameter. if panic happened will wrap error object and return future(nil)
// if just time out will return future(func), err(nil)
func AsyncCall[E any](f func() E, timeout time.Duration) (future func() E, err error) {
	tick := time.NewTicker(timeout)
	defer tick.Stop()
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
	case <-tick.C:
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
