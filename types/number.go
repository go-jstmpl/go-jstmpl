package types

import (
	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/validations"
	"github.com/lestrrat/go-jsschema"
)

type Number struct {
	Schema      *schema.Schema `json:"-"`
	NativeType  string         `json:"-"`
	GoType      string
	ColumnName  string
	ColumnType  string
	Type        string
	Name        string
	key         string
	IsPrivate   bool                     `json:"-"`
	Validations []validations.Validation `json:"-"`
}

func NewNumber(ctx *Context, s *schema.Schema) *Number {
	vs := []validations.Validation{}
	if v, err := validations.NewMaximumValidation(s); err == nil {
		ctx.AddValidation(v)
		vs = append(vs, v)
	}
	if v, err := validations.NewMinimumValidation(s); err == nil {
		ctx.AddValidation(v)
		vs = append(vs, v)
	}
	var gt, cn, ct string
	if s.Extras["go_type"] != nil {
		gt, _ = helpers.GetGoTypeData(s)
	} else {
		gt, _ = helpers.GetGoTypeData(ctx.Raw)
	}

	if s.Extras["column"] != nil {
		cn, ct, _ = helpers.GetColumnData(s)
	} else {
		cn, ct, _ = helpers.GetColumnData(ctx.Raw)
	}
	return &Number{
		Schema:      s,
		NativeType:  "number",
		Type:        "number",
		GoType:      gt,
		ColumnName:  cn,
		ColumnType:  ct,
		Name:        helpers.SpaceToUpperCamelCase(s.Title),
		key:         ctx.Key,
		IsPrivate:   true,
		Validations: vs,
	}
}

func (o Number) Raw() *schema.Schema {
	return o.Schema
}

func (o Number) Title() string {
	return o.Schema.Title
}

func (o Number) Format() string {
	return string(o.Schema.Format)
}

func (o Number) Key() string {
	return o.key
}

func (o Number) Example() interface{} {
	e := o.Schema.Extras["example"]
	if e != nil {
		return e
	}
	return 0
}
