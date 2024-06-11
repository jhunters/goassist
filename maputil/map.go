/*
 * Package maputil to provides utility api for map container
 */
package maputil

import (
	"bytes"
	"encoding/json"
)

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

// JsonEncode 函数将一个类型为 map[E]V 的字典转换为 JSON 格式的字节数组
// 其中 E 必须是可比较的类型，V 可以是任意类型
// 参数 mapa 是待转换的字典
// 返回值为转换后的 JSON 格式的字节数组以及可能出现的错误
func JsonEncode[E comparable, V any](mapa map[E]V) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(mapa)
	return buf.Bytes(), err
}

// JsonDecode 函数接受一个字节切片作为参数，用于解码JSON格式的数据，
// 并将解码后的结果存储在一个以E类型为键，V类型为值的map中返回。
// 如果解码失败，则返回error类型的错误信息。
// 其中，E类型需要是可比较的，V类型则没有限制。
func JsonDecode[E comparable, V any](data []byte) (map[E]V, error) {
	buf := bytes.NewBuffer(data)
	ret := make(map[E]V)
	err := json.NewDecoder(buf).Decode(&ret)
	return ret, err
}
