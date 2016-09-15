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
		"toUpperFirst":          helpers.ToUpperFirst,
		"toLowerFirst":          helpers.ToLowerFirst,
		"notLast":               helpers.NotLast,
		"spaceToUpperCamelCase": helpers.SpaceToUpperCamelCase,
		"snakeToUpperCamelCase": helpers.SnakeToUpperCamelCase,
		"snakeToLowerCamelCase": helpers.SnakeToLowerCamelCase,
		"joinTypes":             helpers.JoinTypes,
		"serialize":             helpers.Serialize,
		"convertTypeForGo":      helpers.ConvertTypeForGo,
		"convertTagsForGo":      helpers.ConvertTagsForGo,
		"convertArrayForGo":     helpers.ConvertArrayForGo,
	}).Parse(string(tmpl)))
	if err := t.Execute(out, model); err != nil {
		return err
	}
	return nil
}
