package concurrent

import "log"

func SafeChanClose[E any](c chan E) (ok bool) {
	ok = true
	defer func() {
		if v := recover(); v != nil {
			log.Println(v)
			ok = false
		}
	}()
	close(c)
	return ok
}
