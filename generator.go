package jstmpl

import (
	"bytes"
	"go/format"
	"os"
	"text/template"

	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/types"
)

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Process(model *types.Root, tmpl []byte, ext string) ([]byte, error) {
	out := bytes.NewBuffer([]byte{})
	t := template.Must(template.New("").Funcs(map[string]interface{}{
		"getEnv":                 os.Getenv,
		"notLast":                helpers.NotLast,
		"joinTypes":              helpers.JoinTypes,
		"serialize":              helpers.Serialize,
		"toStringLiteral":        helpers.ToStringLiteral,
		"convertTypeForGo":       helpers.ConvertTypeForGo,
		"convertTypeInJSONForGo": helpers.ConvertTypeInJSONForGo,
		"jsonTagForGo":           helpers.ConvertJSONTagForGo,
		"xormTagForGo":           helpers.ConvertXORMTagForGo,
		"convertArrayForGo":      helpers.ConvertArrayForGo,
		"convertTypeForJS":       helpers.ConvertTypeForJS,
		"linkTitle":              helpers.LinkTitle,
		"getKeyFromJSONPath":     helpers.GetKeyFromJSONPath,
		"snakeCase":              helpers.SnakeCase,
		"lowerSnakeCase":         helpers.LowerSnakeCase,
		"upperSnakeCase":         helpers.UpperSnakeCase,
		"lowerCamelCase":         helpers.LowerCamelCase,
		"upperCamelCase":         helpers.UpperCamelCase,
	}).Parse(string(tmpl)))
	if err := t.Execute(out, model); err != nil {
		return nil, err
	}

	// Format for each language
	b := out.Bytes()
	switch ext {
	case ".go":
		var err error
		b, err = format.Source(b)
		if err != nil {
			return nil, err
		}
	}

	return b, nil
}
