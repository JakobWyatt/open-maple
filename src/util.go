package main

import (
	"reflect"
)

//simple mathematical functions to pass to other functions
type binaryFunc func(float64, float64) float64

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	return a / b
}

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
