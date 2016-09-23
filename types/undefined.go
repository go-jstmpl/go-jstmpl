package types

import (
	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/validations"
	"github.com/lestrrat/go-jsschema"
)

type Undefined struct {
	Schema      *schema.Schema `json:"-"`
	NativeType  string         `json:"-"`
	ColumnName  string
	ColumnType  string
	Type        string
	Name        string
	key         string
	Private     bool
	Validations []validations.Validation
}

func NewUndefined(ctx *Context, s *schema.Schema) *Undefined {
	vs := []validations.Validation{}
	var cn, ct string

	if s.Extras["column"] != nil {
		cn, ct, _ = helpers.GetColumn(s)
	} else {
		cn, ct, _ = helpers.GetColumn(ctx.Raw)
	}
	var pr bool
	if s.Extras["IsPrivate"] != nil {
		pr, _ = helpers.GetPrivate(s)
	} else {
		pr, _ = helpers.GetPrivate(ctx.Raw)
	}

	return &Undefined{
		Schema:      s,
		NativeType:  "undefined",
		Type:        "undefined",
		ColumnName:  cn,
		ColumnType:  ct,
		Name:        helpers.SpaceToUpperCamelCase(s.Title),
		key:         ctx.Key,
		Private:     pr,
		Validations: vs,
	}
}

func (o Undefined) Raw() *schema.Schema {
	return o.Schema
}

func (o Undefined) Title() string {
	return o.Schema.Title
}

func (o Undefined) Description() string {
	return o.Schema.Description
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
