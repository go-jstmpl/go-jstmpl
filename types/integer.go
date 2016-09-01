package types

import (
	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/validations"
	"github.com/lestrrat/go-jsschema"
)

type Integer struct {
	*schema.Schema
	NativeType  string
	Type        string
	Name        string
	key         string
	IsPrivate   bool
	Validations []validations.Validation
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
	return &Integer{
		Schema:      s,
		NativeType:  "number",
		Type:        "number",
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
