package jstmpl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"
	"text/template"

	"github.com/lestrrat/go-jsschema"
)

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Process(out io.Writer, model *Root, tmpl []byte) error {
	t := template.Must(template.New("").Delims("/*", "*/").Funcs(map[string]interface{}{
		"notLast":               notLast,
		"spaceToUpperCamelCase": spaceToUpperCamelCase,
		"snakeToLowerCamelCase": snakeToLowerCamelCase,
		"joinTypes":             joinTypes,
		"serialize":             serialize,
		"convertTypeToGo":       convertTypeToGo,
	}).Parse(string(tmpl)))
	if err := t.Execute(out, model); err != nil {
		return err
	}
	return nil
}

var (
	rspace = regexp.MustCompile(`\s+`)
	rsnake = regexp.MustCompile(`_`)
)

func notLast(i, len int) bool {
	return i != len-1
}

func spaceToUpperCamelCase(s string) string {
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

func snakeToLowerCamelCase(s string) string {
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

func joinTypes(ts schema.PrimitiveTypes, sep string) string {
	var strs []string
	for _, t := range ts {
		strs = append(strs, t.String())
	}
	return strings.Join(strs, sep)
}

func serialize(v interface{}) string {
	j, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprint(v)
	}
	return string(j)
}

func convertTypeToGo(ts schema.PrimitiveTypes) string {
	if ts.Len() >= 2 {
		return ""
	}

	conv := map[string]string{
		"integer": "int",
		"boolean": "bool",
	}

	for _, t := range ts {
		if conv[t.String()] == "" {
			return t.String()
		}
	}
	return conv[ts[0].String()]
}
