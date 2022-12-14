package concurrent

import "time"

// SafeChanClose to close chan for safty way. if channel is closed will retrun false
func SafeCloseChan[E any](c chan E) (ok bool) {
	ok = true
	defer func() {
		if v := recover(); v != nil {
			ok = false
		}
	}()
	close(c)
	return ok
}

// TrySend try to send value to channel within target time
func TrySendChan[E any](v E, c chan<- E, timeout time.Duration) (ok bool) {
	if timeout <= 0 {
		c <- v
		return true
	}
	tick := time.NewTicker(timeout)
	defer tick.Stop()
	select {
	case c <- v:
		return true
	case <-tick.C:
		return false
	}
}

// TryRecevie try to receive value from channel within target time
func TryRecevieChan[E any](c <-chan E, timeout time.Duration) (ok bool, v E) {
	if timeout <= 0 {
		v = <-c
		return true, v
	}
	tick := time.NewTicker(timeout)
	defer tick.Stop()
	select {
	case v = <-c:
		return true, v
	case <-tick.C:
		return false, v
	}
}
