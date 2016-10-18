package helpers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func ToLiteralForJS(input interface{}) string {
	switch t := input.(type) {
	case bool:
		return fmt.Sprintf("%b", t)
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", t)
	case float32, float64:
		return fmt.Sprintf("%f", t)
	case string:
		return fmt.Sprintf("\"%s\"", t)
	}
	switch reflect.TypeOf(input).Kind() {
	case reflect.Slice:
		t := reflect.ValueOf(input)
		l := t.Len()
		es := make([]string, l)
		for i := 0; i < l; i++ {
			e := t.Index(i)
			es[i] = ToLiteralForGo(e.Interface())
		}
		return fmt.Sprintf("[%s]", strings.Join(es, ", "))
	}
	b, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}

func ConvertTypeForJS(s string) string {
	v, ok := map[string]string{
		"integer": "number",
	}[s]
	if !ok {
		return s
	}
	return v
}
