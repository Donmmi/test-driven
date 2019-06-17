package main

import "reflect"

func Walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Slice, reflect.Array:
		numFields := val.Len()
		for i := 0; i < numFields; i++ {
			Walk(val.Index(i).Interface(), fn)
		}
	case reflect.Struct:
		numFields := val.NumField()
		for i := 0; i < numFields; i++ {
			Walk(val.Field(i).Interface(), fn)
		}
	case reflect.Map:
		numFields := len(val.MapKeys())
		keys := val.MapKeys()
		for i := 0; i < numFields; i++ {
			Walk(val.MapIndex(keys[i]).Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
