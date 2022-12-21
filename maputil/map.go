/*
 * Package maputil to provides utility api for map container
 */
package maputil

// Clone copy all key and value to a new map
func Clone[E comparable, V any](mapa map[E]V) map[E]V {
	size := len(mapa)
	ret := make(map[E]V, size)
	for k, v := range mapa {
		ret[k] = v
	}
	return ret
}

// AddAll merge the target mapb into mapa
func AddAll[E comparable, V any](mapa, mapb map[E]V) map[E]V {
	ret := Clone(mapa)
	for k, v := range mapb {
		_, exist := ret[k]
		if !exist {
			ret[k] = v
		}
	}
	return ret
}

// Clear remove all keys and values in map
func Clear[E comparable, V any](mapa map[E]V) {
	if mapa == nil {
		return
	}

	for k := range mapa {
		delete(mapa, k)
	}
}
