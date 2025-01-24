package shortcodes

import (
	"reflect"
	"strings"
)

const (
	Space      = " "
	Underscore = "_"
	CommaSpace = ", "
)

func ReplaceSpacesForDelimiter(value, delimiter string) string {
	return strings.Replace(value, Space, delimiter, -1)
}

func IsStruct(i interface{}) bool {
	return reflect.ValueOf(i).Type().Kind() == reflect.Struct
}
