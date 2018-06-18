package util

import "reflect"

func Unique(array []reflect.Type) []reflect.Type {
	// Use map to record duplicates as we find them.
	encountered := map[interface{}]bool{}
	var result []reflect.Type

	for _, v := range array {
		if encountered[v] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[v] = true
			// Append to result slice.
			result = append(result, v)
		}
	}
	// Return the new slice.
	return result
}


