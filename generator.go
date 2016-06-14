package jstmpl

import (
	"io"
	"text/template"
)

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Process(out io.Writer, model *TSModel, tmpl []byte) error {
	t := template.Must(template.New("jstmpl").Delims("/*", "*/").Funcs(map[string]interface{}{
		"keyToPropName": KeyToPropName,
		"notLast": func(i, len int) bool {
			return i != len-1
		},
	}).Parse(string(tmpl)))
	if err := t.Execute(out, model); err != nil {
		return err
	}
	return nil
}
