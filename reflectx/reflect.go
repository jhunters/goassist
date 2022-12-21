/*
 * Package reflectx to provides utility api for reflect operation
 */
package reflectx

import (
	"fmt"
	"reflect"
	"strings"
)

// ValueOf return value type of target t,
// if value is pointer type then the second return paramter is true
func ValueOf(t interface{}) (reflect.Value, bool) {
	v := reflect.ValueOf(t)
	isPtr := false
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		isPtr = true
	}
	return v, isPtr
}

// TypeOf return type of target t,
// if type is pointer type then the second return paramter is true
func TypeOf(t interface{}) (reflect.Type, bool) {
	v := reflect.TypeOf(t)
	isPtr := false
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		isPtr = true
	}
	return v, isPtr
}

// GetMethods returns suitable methods of type
func GetMethods(typ reflect.Type) map[string]*reflect.Method {
	methods := make(map[string]*reflect.Method)
	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		mname := method.Name
		// Method must be exported.
		if method.PkgPath != "" {
			continue
		}

		methods[mname] = &method
	}
	return methods
}

// CallMethodByName 根据方法名，入参数，反射调用对应方法
func CallMethodByName(rcvr interface{}, methodName string, params ...interface{}) ([]reflect.Value, error) {

	typ := reflect.TypeOf(rcvr)

	kind := typ.Kind()
	if kind == reflect.Pointer { // 处理指针类型
		kind = typ.Elem().Kind()
	}

	if kind != reflect.Struct && kind != reflect.Interface {
		return nil, fmt.Errorf("param 'rcvr' should be struct or interface type.")
	}

	// 根据名称查找方法
	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		mname := method.Name

		if strings.EqualFold(mname, methodName) {
			return callMethod(rcvr, &method, params...)
		}
	}

	// 如果方法未找到，返回nil
	return nil, fmt.Errorf("method name '%s' not found", methodName)
}

// CallMethod 根据方法反射类型，入参数，反射调用对应方法
func callMethod(rcvr interface{}, method *reflect.Method, params ...interface{}) ([]reflect.Value, error) {

	// 封装入参
	paramSize := len(params) + 1
	paramValues := make([]reflect.Value, paramSize)
	paramValues[0] = reflect.ValueOf(rcvr)
	i := 1
	for _, v := range params {
		paramValues[i] = reflect.ValueOf(v)
		i++
	}

	// 调用 反射的Call方法，进行调用
	returnValues := method.Func.Call(paramValues)

	return returnValues, nil
}

// SetSimpleValue set value to target obj and field name. return true if set success
// not supports array, slice, struct and map type
func SetValue(rcvr interface{}, fieldName string, value any) bool {
	v := reflect.ValueOf(rcvr)
	t := reflect.TypeOf(rcvr)
	if v.Kind() != reflect.Pointer || v.IsNil() {
		return false
	}

	if f, ok := t.Elem().FieldByName(fieldName); ok {
		fvalue := v.Elem().FieldByIndex(f.Index)

		if !fvalue.CanSet() {
			return false
		}

		switch fvalue.Kind() {
		case reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int64, reflect.Int, reflect.Uint8, reflect.Uint, reflect.Uint16,
			reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Bool, reflect.Array, reflect.Slice, reflect.Map:
			fvalue.Set(reflect.ValueOf(value))
		case reflect.String:
			fvalue.SetString(fmt.Sprintf("%s", value))
		case reflect.Chan:
			return fvalue.TrySend(reflect.ValueOf(value))
		}
		return true
	}

	return false
}
func GetValue[S any](rcvr interface{}, fieldName string) (S, bool) {
	v := reflect.ValueOf(rcvr)
	t := reflect.TypeOf(rcvr)
	var ret S
	if v.Kind() != reflect.Pointer || v.IsNil() {
		return ret, false
	}

	if f, ok := t.Elem().FieldByName(fieldName); ok {
		fvalue := v.Elem().FieldByIndex(f.Index)

		if !fvalue.CanInterface() {
			ret, ok := fvalue.Interface().(S)
			return ret, ok
		}
	}
	return ret, false

}
