package utils

import "reflect"

// GetType returns the type of the given variable
func GetType(variable interface{}) string {
	if t := reflect.TypeOf(variable); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
