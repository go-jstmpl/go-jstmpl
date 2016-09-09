package types

import (
	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/validations"
	"github.com/lestrrat/go-jsschema"
)

type Undefined struct {
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

func NewUndefined(ctx *Context, s *schema.Schema) *Undefined {
	vs := []validations.Validation{}
	gt, cn, ct, _ := helpers.GetExtraData(s)
	return &Undefined{
		Schema:      s,
		NativeType:  "undefined",
		Type:        "undefined",
		GoType:      gt,
		ColumnName:  cn,
		ColumnType:  ct,
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
