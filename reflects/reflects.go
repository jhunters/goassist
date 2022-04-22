/*
 * @Author: Malin Xie
 * @Description: utility tools for reflect
 * @Date: 2022-01-21 11:43:49
 */
package reflects

import (
	"fmt"
	"reflect"
	"strconv"
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

// CallMethod
func CallMethod(rcvr interface{}, method *reflect.Method, params ...interface{}) ([]reflect.Value, error) {

	paramSize := len(params) + 1
	paramValues := make([]reflect.Value, paramSize)
	paramValues[0] = reflect.ValueOf(rcvr)
	i := 1
	for _, v := range params {
		paramValues[i] = reflect.ValueOf(v)
		i++
	}

	returnValues := method.Func.Call(paramValues)

	return returnValues, nil
}

// CallMethodName
func CallMethodByName(rcvr interface{}, methodName string, params ...interface{}) ([]reflect.Value, error) {

	typ := reflect.TypeOf(rcvr)
	fmt.Println(typ.NumMethod())
	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		mname := method.Name

		if strings.EqualFold(mname, methodName) {
			return CallMethod(rcvr, &method, params...)
		}
	}

	return nil, fmt.Errorf("method name '%s' not found", methodName)
}

// SetSimpleValue set value to target obj and field name. return true if set success
// not supports array, slice, struct and map type
func SetSimpleValue(rcvr interface{}, fieldName string, value string) bool {
	v := reflect.ValueOf(rcvr)
	t := reflect.TypeOf(rcvr)

	if f, ok := t.FieldByName(fieldName); ok {
		fvalue := v.FieldByIndex(f.Index)

		if !fvalue.CanSet() {
			return false
		}

		switch fvalue.Kind() {
		case reflect.Int32:
			v, _ := strconv.Atoi(value)
			fvalue.Set(reflect.ValueOf(int32(v)))
		case reflect.Int16:
			v, _ := strconv.Atoi(value)
			fvalue.Set(reflect.ValueOf(int16(v)))
		case reflect.Int8:
			v, _ := strconv.Atoi(value)
			fvalue.Set(reflect.ValueOf(int8(v)))
		case reflect.Int64:
			v, _ := strconv.Atoi(value)
			fvalue.Set(reflect.ValueOf(int64(v)))
		case reflect.Int:
			v, _ := strconv.Atoi(value)
			fvalue.Set(reflect.ValueOf(int(v)))
		case reflect.Uint8:
			v, _ := strconv.Atoi(value)
			fvalue.Set(reflect.ValueOf(uint8(v)))
		case reflect.Uint:
			v, _ := strconv.Atoi(value)
			fvalue.Set(reflect.ValueOf(uint(v)))
		case reflect.Uint16:
			v, _ := strconv.Atoi(value)
			fvalue.Set(reflect.ValueOf(uint16(v)))
		case reflect.Uint32:
			v, _ := strconv.Atoi(value)
			fvalue.Set(reflect.ValueOf(uint32(v)))
		case reflect.Uint64:
			v, _ := strconv.Atoi(value)
			fvalue.Set(reflect.ValueOf(uint64(v)))
		case reflect.Float32:
			v, _ := strconv.ParseFloat(value, 32)
			fvalue.Set(reflect.ValueOf(float32(v)))
		case reflect.Float64:
			v, _ := strconv.ParseFloat(value, 64)
			fvalue.Set(reflect.ValueOf(v))
		case reflect.String:
			fvalue.SetString(value)
		case reflect.Bool:
			if strings.Compare(value, "1") == 0 || strings.Compare(strings.ToLower(value), "true") == 0 {
				fvalue.SetBool(true)
			} else {
				fvalue.SetBool(false)
			}
		}
		return true
	}

	return false
}
