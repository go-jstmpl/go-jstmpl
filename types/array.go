package types

import (
	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/validations"
	schema "github.com/lestrrat/go-jsschema"
)

type Array struct {
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
	Item        Schema                   `json:",omitempty"`
	Items       *ItemSpec                `json:",omitempty"`
}

func NewArray(ctx *Context, s *schema.Schema) *Array {
	gt, cn, ct, _ := helpers.GetExtraData(s)
	return &Array{
		Schema:     s,
		NativeType: "array",
		GoType:     gt,
		ColumnName: cn,
		ColumnType: ct,
		Type:       helpers.SpaceToUpperCamelCase(s.Title),
		Name:       helpers.SpaceToUpperCamelCase(s.Title),
		key:        ctx.Key,
		IsPrivate:  false,
		Items: &ItemSpec{
			ItemSpec: s.Items,
			Schemas:  make([]Schema, len(s.Items.Schemas)),
		},
	}
}

func (o Array) Raw() *schema.Schema {
	return o.Schema
}

func (o Array) Title() string {
	return o.Schema.Title
}

func (o Array) Format() string {
	return string(o.Schema.Format)
}

func (o Array) Key() string {
	return o.key
}

func (o Array) Example() interface{} {
	e := make([]interface{}, len(o.Items.Schemas))
	for i, s := range o.Items.Schemas {
		e[i] = s.Example()
	}
	return e
}
