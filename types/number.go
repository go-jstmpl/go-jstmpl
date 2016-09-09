package types

import (
	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/validations"
	"github.com/lestrrat/go-jsschema"
)

type Number struct {
	Schema      *schema.Schema           `json:"-"`
	NativeType  string                   `json:"-"`
	GoType      string                   `json:",omitempty"`
	ColumnName  string                   `json:",omitempty"`
	ColumnType  string                   `json:",omitempty"`
	Type        string                   `json:",omitempty"`
	Name        string                   `json:",omitempty"`
	key         string                   `json:",omitempty"`
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
	gt, cn, ct, _ := helpers.GetExtraData(s)
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
