package helpers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	schema "github.com/lestrrat/go-jsschema"
)

func ToLiteralForGo(input interface{}) string {
	switch t := input.(type) {
	case bool:
		return fmt.Sprintf("%t", t)
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
		return fmt.Sprintf("[]%s{%s}", reflect.TypeOf(input).Elem().Kind().String(), strings.Join(es, ", "))
	}
	b, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}

func ConvertTypeForGo(s string) string {
	v, ok := map[string]string{
		"integer": "int64",
		"boolean": "bool",
		"number":  "float64",
	}[s]
	if !ok {
		return s
	}
	return v
}

func ConvertTypeInJSONForGo(s string) string {
	v, ok := map[string]string{
		"integer": "float64",
		"boolean": "bool",
		"number":  "float64",
	}[s]
	if !ok {
		return s
	}
	return v
}

func ConvertArrayForGo(m []string) string {
	s := "[]string{"
	for i, v := range m {
		if i != len(m)-1 {
			s += fmt.Sprintf("\"%s\",", UpperCamelCase(v))
		} else {
			s += fmt.Sprintf("\"%s\"", UpperCamelCase(v))
		}
	}
	s += "}"
	return s
}

func ConvertJSONTagForGo(tag string) string {
	if tag == "" || tag == "-" {
		return "json:\"-\""
	}
	return fmt.Sprintf("json:\"%s,omitempty\"", tag)
}

func ConvertXORMTagForGo(tag string) string {
	if tag == "" || tag == "-" {
		return "xorm:\"-\""
	}
	return fmt.Sprintf("xorm:\"%s\"", tag)
}

func GetTable(ts *schema.Schema) (tn string, err error) {
	if ts.Extras["table"] == nil {
		return
	}

	s, ok := ts.Extras["table"].(map[string]interface{})
	if !ok {
		err = fmt.Errorf("table %v is invalid type", ts.Extras["table"])
	}

	t, ok := s["name"].(string)
	if !ok {
		err = fmt.Errorf("table[name] %v is invalid type", s["name"])
	}

	tn = t
	return
}

func GetPrivate(ts *schema.Schema) (bool, error) {
	if ts.Extras["private"] == nil {
		return false, nil
	}

	c, ok := ts.Extras["private"].(bool)
	if !ok {
		return false, fmt.Errorf("private %v is invalid type", ts.Extras["private"])
	}
	return c, nil
}

func GetColumn(ts *schema.Schema) (cn, ct string, err error) {
	if ts.Extras["column"] == nil {
		return
	}

	c, ok := ts.Extras["column"].(map[string]interface{})
	if !ok {
		err = fmt.Errorf("column %v is invalid type", ts.Extras["column"])
	}

	n, ok := c["name"].(string)
	if !ok {
		err = fmt.Errorf("column[name] %v is invalid type", c["name"])
	}
	cn = n

	t, ok := c["db_type"].(string)
	if !ok {
		err = fmt.Errorf("column[db_type] %v is invalid type", c["db_type"])
	}
	ct = t
	return
}
