package types

import (
	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/validations"
	"github.com/lestrrat/go-jsschema"
)

type String struct {
	*schema.Schema
	NativeType  string
	Type        string
	Name        string
	key         string
	IsPrivate   bool
	Validations []validations.Validation
}

func NewString(ctx *Context, s *schema.Schema) *String {
	vs := []validations.Validation{}
	if v, err := validations.NewFormatValidation(s); err == nil {
		ctx.AddValidation(v)
		vs = append(vs, v)
	}
	if v, err := validations.NewMinLengthValidation(s); err == nil {
		ctx.AddValidation(v)
		vs = append(vs, v)
	}
	if v, err := validations.NewMaxLengthValidation(s); err == nil {
		ctx.AddValidation(v)
		vs = append(vs, v)
	}
	if v, err := validations.NewPatternValidation(s); err == nil {
		ctx.AddValidation(v)
		vs = append(vs, v)
	}
	return &String{
		Schema:      s,
		NativeType:  "string",
		Type:        "string",
		Name:        helpers.SpaceToUpperCamelCase(s.Title),
		key:         ctx.Key,
		IsPrivate:   true,
		Validations: vs,
	}
}

func (o String) Raw() *schema.Schema {
	return o.Schema
}

func (o String) Title() string {
	return o.Schema.Title
}

func (o String) Format() string {
	return string(o.Schema.Format)
}

func (o String) Key() string {
	return o.key
}

func (o String) Example() interface{} {
	e := o.Schema.Extras["example"]
	if e != nil {
		return e
	}
	return ""
}
