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

	var pr bool
	if s.Extras["private"] != nil {
		pr, _ = helpers.GetPrivate(s)
	} else {
		pr, _ = helpers.GetPrivate(ctx.Raw)
	}

	var cn, ct string
	if s.Extras["column"] != nil {
		cn, ct, _ = helpers.GetColumn(s)
	} else {
		cn, ct, _ = helpers.GetColumn(ctx.Raw)
	}

	return &Undefined{
		Schema:      s,
		NativeType:  "undefined",
		Type:        "undefined",
		ColumnName:  cn,
		ColumnType:  ct,
		Name:        helpers.UpperCamelCase(s.Title),
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

func (o Undefined) ReadOnly() bool {
	v := o.Schema.Extras["readOnly"]
	if v == nil {
		return false
	}
	r, ok := v.(bool)
	if !ok {
		return false
	}
	return r
}

func (o Undefined) Example(withoutReadOnly bool) interface{} {
	v := o.Schema.Extras["example"]
	if v == nil {
		return ""
	}
	return v
}

func (o Undefined) ExampleString() string {
	return helpers.Serialize(o.Example(false))
}
