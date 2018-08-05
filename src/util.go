package main

import (
	"reflect"
)

//inArray checks if a value is in an array
func inArray(val interface{}, array interface{}) bool {
	if reflect.TypeOf(array).Kind() != reflect.Slice {
		return false
	}

	arrayData := reflect.ValueOf(array)
	for i := 0; i < arrayData.Len(); i++ {
		if reflect.DeepEqual(val, arrayData.Index(i).Interface()) {
			return true
		}
	}

	return false
}
