package utils

import (
	"reflect"

	"github.com/stoewer/go-strcase"
)

// GetType returns the type of the given variable
func GetTableName(variable interface{}) string {
	if t := reflect.TypeOf(variable); t.Kind() == reflect.Ptr {
		return strcase.SnakeCase(t.Elem().Name()) + "s"
	} else {
		return strcase.SnakeCase(t.Name()) + "s"
	}
}
