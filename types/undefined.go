package types

import (
	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/validations"
	"github.com/lestrrat/go-jsschema"
)

type Undefined struct {
	*schema.Schema
	NativeType  string
	Type        string
	Name        string
	key         string
	IsPrivate   bool
	Validations []validations.Validation
}

func NewUndefined(ctx *Context, s *schema.Schema) *Undefined {
	vs := []validations.Validation{}
	return &Undefined{
		Schema:      s,
		NativeType:  "undefined",
		Type:        "undefined",
		Name:        helpers.SpaceToUpperCamelCase(s.Title),
		key:         ctx.Key,
		IsPrivate:   true,
		Validations: vs,
	}
}

func (o Undefined) Raw() *schema.Schema {
	return o.Schema
}

func (o Undefined) Title() string {
	return o.Schema.Title
}

func (o Undefined) Key() string {
	return o.key
}

func (o Undefined) Example() interface{} {
	e := o.Schema.Extras["example"]
	if e != nil {
		return e
	}
	return ""
}
