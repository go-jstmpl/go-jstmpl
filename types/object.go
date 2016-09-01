package types

import (
	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/validations"
	"github.com/lestrrat/go-jsschema"
)

type Object struct {
	*schema.Schema
	NativeType  string
	Type        string
	Name        string
	key         string
	IsPrivate   bool
	Validations []validations.Validation
	Properties  []Schema
}

func NewObject(ctx *Context, s *schema.Schema) *Object {
	return &Object{
		Schema:     s,
		NativeType: "object",
		Type:       helpers.SpaceToUpperCamelCase(s.Title),
		Name:       helpers.SpaceToUpperCamelCase(s.Title),
		key:        ctx.Key,
		IsPrivate:  false,
		Properties: []Schema{},
	}
}

func (o Object) Raw() *schema.Schema {
	return o.Schema
}

func (o Object) Title() string {
	return o.Schema.Title
}

func (o Object) Format() string {
	return string(o.Schema.Format)
}

func (o Object) Key() string {
	return o.key
}

func (o Object) Example() interface{} {
	e := make(map[string]interface{})
	for _, s := range o.Properties {
		e[s.Key()] = s.Example()
	}
	return e
}
