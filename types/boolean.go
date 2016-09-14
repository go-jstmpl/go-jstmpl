package types

import (
	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/validations"
	"github.com/lestrrat/go-jsschema"
)

type Boolean struct {
	Schema      *schema.Schema `json:"-"`
	NativeType  string         `json:"-"`
	ColumnName  string
	ColumnType  string
	Type        string
	Name        string
	key         string
	IsPrivate   bool                     `json:"-"`
	Validations []validations.Validation 
}

func NewBoolean(ctx *Context, s *schema.Schema) *Boolean {
	vs := []validations.Validation{}
	var cn, ct string

	if s.Extras["column"] != nil {
		cn, ct, _ = helpers.GetColumnData(s)
	} else {
		cn, ct, _ = helpers.GetColumnData(ctx.Raw)
	}

	return &Boolean{
		Schema:      s,
		NativeType:  "boolean",
		Type:        "boolean",
		ColumnName:  cn,
		ColumnType:  ct,
		Name:        helpers.SpaceToUpperCamelCase(s.Title),
		key:         ctx.Key,
		IsPrivate:   true,
		Validations: vs,
	}
}

func (o Boolean) Raw() *schema.Schema {
	return o.Schema
}

func (o Boolean) Title() string {
	return o.Schema.Title
}

func (o Boolean) Format() string {
	return string(o.Schema.Format)
}

func (o Boolean) Key() string {
	return o.key
}

func (o Boolean) Example() interface{} {
	e := o.Schema.Extras["example"]
	if e != nil {
		return e
	}
	return false
}
