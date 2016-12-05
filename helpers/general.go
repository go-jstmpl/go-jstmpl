package helpers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	schema "github.com/lestrrat/go-jsschema"
)

var (
	rspace = regexp.MustCompile(`\s+`)
	rsnake = regexp.MustCompile(`_`)
)

func NotLast(i, len int) bool {
	return i != len-1
}

func GetKeyFromJSONPath(url string) string {
	s := strings.Split(url, "/")
	return UpperCamelCase(s[len(s)-1])
}

func LinkTitle(title, method, suffix string) string {
	return method + UpperCamelCase(title) + suffix
}

func JoinTypes(ts schema.PrimitiveTypes, sep string) string {
	var strs []string
	for _, t := range ts {
		strs = append(strs, t.String())
	}
	return strings.Join(strs, sep)
}

func Serialize(v interface{}) string {
	j, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprint(v)
	}
	return string(j)
}

func ToStringLiteral(v interface{}) string {
	s, ok := v.(string)
	if ok {
		return fmt.Sprintf("\"%s\"", s)
	}
	return fmt.Sprint(v)
}

func Slice(s ...interface{}) []interface{} {
	return s
}

func InSlice(e, s interface{}, condition func(interface{}, reflect.Value) bool) bool {
	switch reflect.TypeOf(s).Kind() {
	case reflect.Slice:
		t := reflect.ValueOf(s)
		if condition(e, t) {
			return true
		}
		return false
	default:
		return false
	}
}

func In(e interface{}, s interface{}) bool {
	cond := func(e interface{}, t reflect.Value) bool {
		for i := 0; i < t.Len(); i++ {
			v := t.Index(i)
			if v.Interface() == e {
				return true
			}
		}
		return false
	}
	return InSlice(e, s, cond)
}

func InMapKeys(e interface{}, s interface{}) bool {
	cond := func(e interface{}, t reflect.Value) bool {
		for i := 0; i < t.Len(); i++ {
			v := t.Index(i)
			m, ok := v.Interface().(map[string]interface{})
			if !ok {
				continue
			}
			for k := range m {
				if k == e {
					return true
				}
			}
		}
		return false
	}
	return InSlice(e, s, cond)
}
