package types

import (
	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/validations"
	schema "github.com/lestrrat/go-jsschema"
)

type Array struct {
	*schema.Schema
	NativeType  string
	Type        string
	Name        string
	key         string
	IsPrivate   bool
	Validations []validations.Validation
	Item        Schema
	Items       *ItemSpec
}

func NewArray(ctx *Context, s *schema.Schema) *Array {
	return &Array{
		Schema:     s,
		NativeType: "array",
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
