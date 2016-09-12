package types

import (
	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/validations"
	"github.com/lestrrat/go-jsschema"
)

type Undefined struct {
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

func NewUndefined(ctx *Context, s, t *schema.Schema) *Undefined {
	vs := []validations.Validation{}
	var gt, cn, ct string
	if s.Extras["go_type"] != nil {
		gt, _ = helpers.GetGoTypeData(s)
	} else {
		gt, _ = helpers.GetGoTypeData(t)
	}

	if s.Extras["column"] != nil {
		cn, ct, _ = helpers.GetColumnData(s)
	} else {
		cn, ct, _ = helpers.GetColumnData(t)
	}
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
