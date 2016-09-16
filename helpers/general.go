package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func HandlerName(title, method string) string {
	return method + SpaceToUpperCamelCase(title)
}

func SpaceToUpperCamelCase(s string) string {
	if s == "" {
		return ""
	}
	buf := bytes.Buffer{}
	for _, p := range rspace.Split(s, -1) {
		buf.WriteString(strings.ToUpper(p[:1]))
		buf.WriteString(p[1:])
	}
	return buf.String()
}

func SnakeToUpperCamelCase(s string) string {
	if s == "" {
		return ""
	}
	buf := bytes.Buffer{}
	for _, p := range rsnake.Split(s, -1) {
		buf.WriteString(strings.ToUpper(p[:1]))
		buf.WriteString(p[1:])
	}
	return buf.String()
}

func SnakeToLowerCamelCase(s string) string {
	if s == "" {
		return ""
	}
	buf := bytes.Buffer{}
	for i, p := range rsnake.Split(s, -1) {
		if i == 0 {
			buf.WriteString(p)
			continue
		}
		buf.WriteString(strings.ToUpper(p[:1]))
		buf.WriteString(p[1:])
	}
	return buf.String()
}

func ToUpperFirst(s string) string {
	buf := bytes.Buffer{}
	buf.WriteString(strings.ToUpper(s[:1]))
	buf.WriteString(s[1:])
	return buf.String()
}

func ToLowerFirst(s string) string {
	buf := bytes.Buffer{}
	buf.WriteString(strings.ToLower(s[:1]))
	buf.WriteString(s[1:])
	return buf.String()
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
