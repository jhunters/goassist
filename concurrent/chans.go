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
	if c == nil {
		return false
	}
	defer func() {
		if v := recover(); v != nil {
			ok = false
		}
	}()
	if timeout <= 0 {
		c <- v
		return true
	}
	tout, cancel := timeoutF(timeout)
	defer cancel()
	select {
	case c <- v:
		return true
	case <-tout:
		return false
	}
}

// TryRecevie try to receive value from channel within target time
func TryRecevieChan[E any](c <-chan E, timeout time.Duration) (ok bool, v E) {
	if c == nil {
		return
	}
	if timeout <= 0 {
		v = <-c
		return true, v
	}
	tout, cancel := timeoutF(timeout)
	defer cancel()
	select {
	case v = <-c:
		return true, v
	case <-tout:
		return false, v
	}
}
