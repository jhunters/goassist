/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-01-05 13:01:14
 */
package mapx

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
