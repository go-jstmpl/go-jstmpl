package jstmpl

import (
	"bytes"
	"go/format"
	"os"
	"text/template"

	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/types"
	"github.com/pkg/errors"
)

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

type FormatError struct {
	err error
}

func NewFormatError(e error) FormatError {
	return FormatError{
		err: errors.Wrap(e, "fail to format Go"),
	}
}

func (err FormatError) Error() string {
	return err.err.Error()
}

func (g *Generator) Process(model *types.Root, tmpl []byte, ext string) ([]byte, error) {
	out := bytes.NewBuffer([]byte{})
	t := template.Must(template.New("").Funcs(map[string]interface{}{
		"getEnv":                 os.Getenv,
		"add":                    helpers.Add,
		"sub":                    helpers.Sub,
		"slice":                  helpers.Slice,
		"in":                     helpers.In,
		"notLast":                helpers.NotLast,
		"joinTypes":              helpers.JoinTypes,
		"serialize":              helpers.Serialize,
		"toStringLiteral":        helpers.ToStringLiteral,
		"buildURLToken":          helpers.BuildURLToken,
		"toPathLikeGorilla":      helpers.CatchErrorForString(helpers.ToPathLikeGorilla),
		"toPathLikeSinatra":      helpers.CatchErrorForString(helpers.ToPathLikeSinatra),
		"toLiteralForGo":         helpers.ToLiteralForGo,
		"convertTypeForGo":       helpers.ConvertTypeForGo,
		"convertTypeInJSONForGo": helpers.ConvertTypeInJSONForGo,
		"jsonTagForGo":           helpers.ConvertJSONTagForGo,
		"xormTagForGo":           helpers.ConvertXORMTagForGo,
		"convertArrayForGo":      helpers.ConvertArrayForGo,
		"toLiteralForJS":         helpers.ToLiteralForJS,
		"convertTypeForJS":       helpers.ConvertTypeForJS,
		"getKeyFromJSONPath":     helpers.GetKeyFromJSONPath,
		"snakeCase":              helpers.SnakeCase,
		"lowerSnakeCase":         helpers.LowerSnakeCase,
		"upperSnakeCase":         helpers.UpperSnakeCase,
		"lowerCamelCase":         helpers.LowerCamelCase,
		"upperCamelCase":         helpers.UpperCamelCase,
		// Deprecated: use printf
		"linkTitle": helpers.LinkTitle,
	}).Parse(string(tmpl)))
	if err := t.Execute(out, model); err != nil {
		return nil, errors.Wrap(err, "fail to parse template")
	}

	// Format for each language
	b := out.Bytes()
	switch ext {
	case ".go":
		var err error
		formatted, err := format.Source(b)
		if err != nil {
			return b, NewFormatError(err)
		}
		b = formatted
	}

	return b, nil
}
