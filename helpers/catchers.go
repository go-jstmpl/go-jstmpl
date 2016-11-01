package helpers

import (
	"fmt"
	"reflect"
	"runtime"
)

func CatchError(fn func(interface{}) (interface{}, error)) func(interface{}) interface{} {
	return func(in interface{}) interface{} {
		out, err := fn(in)
		if err != nil {
			name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
			fmt.Printf("Call '%s' with an argument '%s' returns error '%s'\n", name, in, err)
			return ""
		}
		return out
	}
}

func CatchErrorForString(fn func(string) (string, error)) func(string) string {
	return func(in string) string {
		out, err := fn(in)
		if err != nil {
			name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
			fmt.Printf("Call '%s' with an argument '%s' returns error '%s'\n", name, in, err)
			return ""
		}
		return out
	}
}
