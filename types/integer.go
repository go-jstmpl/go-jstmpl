package types

import (
	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/validations"
	"github.com/lestrrat/go-jsschema"
)

type Integer struct {
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

func NewInteger(ctx *Context, s *schema.Schema) *Integer {
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
	return &Integer{
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

func (o Integer) Raw() *schema.Schema {
	return o.Schema
}

func (o Integer) Title() string {
	return o.Schema.Title
}

func (o Integer) Format() string {
	return string(o.Schema.Format)
}

func (o Integer) Key() string {
	return o.key
}

func (o Integer) Example() interface{} {
	e := o.Schema.Extras["example"]
	if e != nil {
		return e
	}
	return 0
}
