package concurrent

// SafeChanClose to close chan for safty way. if channel is closed will retrun false
func SafeChanClose[E any](c chan E) (ok bool) {
	ok = true
	defer func() {
		if v := recover(); v != nil {
			ok = false
		}
	}()
	close(c)
	return ok
}
