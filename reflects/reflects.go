/*
 * @Author: Malin Xie
 * @Description: utility tools for reflect
 * @Date: 2022-01-21 11:43:49
 */
package reflects

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

// GetMethods returns suitable methods of typ
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
func CallMethodName(rcvr interface{}, methodName string, params ...interface{}) ([]reflect.Value, error) {

	typ := reflect.TypeOf(rcvr)

	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		mname := method.Name

		if strings.EqualFold(mname, methodName) {
			return CallMethod(rcvr, &method, params...)
		}
	}

	return nil, fmt.Errorf("method name '%s' not found", methodName)
}
