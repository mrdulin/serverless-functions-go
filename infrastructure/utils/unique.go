package utils

import (
	"reflect"
)

func Unique(slice []interface{}) []interface{} {
	var result []interface{}
	for _, x := range slice {
		f := false
		for _, y := range result {
			if reflect.DeepEqual(x, y) {
				f = true
				break
			}
		}
		if !f {
			result = append(result, x)
		}
	}
	return result
}

func getField(s interface{}, field string) interface{} {
	r := reflect.ValueOf(s)
	f := reflect.Indirect(r).FieldByName(field)
	return f
}

func UniqueV2(slice interface{}) []interface{} {
	var ret = make([]interface{}, 0)
	val := reflect.ValueOf(slice)
	for i := 0; i < val.Len(); i++ {
		f := false
		for j := 0; j < len(ret); j++ {
			if reflect.DeepEqual(ret[j], val.Index(i).Interface()) {
				f = true
				break
			}
		}
		if !f {
			ret = append(ret, val.Index(i).Interface())
		}
	}
	return ret
}

// Unique best
// https://play.golang.org/p/ckNErHx81oQ
