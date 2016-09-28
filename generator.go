package jstmpl

import (
	"io"
	"text/template"

	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/types"
)

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Process(out io.Writer, model *types.Root, tmpl []byte) error {
	t := template.Must(template.New("").Delims("/*", "*/").Funcs(map[string]interface{}{
		"notLast":            helpers.NotLast,
		"joinTypes":          helpers.JoinTypes,
		"serialize":          helpers.Serialize,
		"toStringLiteral":    helpers.ToStringLiteral,
		"convertTypeForGo":   helpers.ConvertTypeForGo,
		"convertTagsForGo":   helpers.ConvertTagsForGo,
		"convertArrayForGo":  helpers.ConvertArrayForGo,
		"convertTypeForJS":   helpers.ConvertTypeForJS,
		"linkTitle":          helpers.LinkTitle,
		"getKeyFromJSONPath": helpers.GetKeyFromJSONPath,
		"snakeCase":          helpers.SnakeCase,
		"lowerSnakeCase":     helpers.LowerSnakeCase,
		"upperSnakeCase":     helpers.UpperSnakeCase,
		"lowerCamelCase":     helpers.LowerCamelCase,
		"upperCamelCase":     helpers.UpperCamelCase,
		// Deprecated
		"spaceToUpperCamelCase": helpers.SpaceToUpperCamelCase,
		"snakeToUpperCamelCase": helpers.SnakeToUpperCamelCase,
		"snakeToLowerCamelCase": helpers.SnakeToLowerCamelCase,
		"toUpperFirst":          helpers.ToUpperFirst,
		"toLowerFirst":          helpers.ToLowerFirst,
	}).Parse(string(tmpl)))
	if err := t.Execute(out, model); err != nil {
		return err
	}
	return nil
}
