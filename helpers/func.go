package helpers

import (
	"reflect"
	"strings"
)

func RemoveWhiteSpace(text string) string {
	return strings.ReplaceAll(text, " ", "")
}

func IsEmpty(text interface{}) (r bool) {
	if reflect.TypeOf(text).String() == "int" {
		return reflect.ValueOf(text).Int() == 0
	} else if reflect.TypeOf(text).String() == "string" {
		r = RemoveWhiteSpace(reflect.ValueOf(text).String()) == ""
	}

	return
}
