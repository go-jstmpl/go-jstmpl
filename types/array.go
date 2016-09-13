package types

import (
	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/validations"
	schema "github.com/lestrrat/go-jsschema"
)

type Array struct {
	Schema      *schema.Schema `json:"-"`
	NativeType  string         `json:"-"`
	GoType      string
	ColumnName  string
	ColumnType  string
	Type        string
	Name        string
	key         string
	IsPrivate   bool
	Properties  []Schema
	Validations []validations.Validation `json:"-"`
	Item        Schema
	Items       *ItemSpec
}

func NewArray(ctx *Context, s, t *schema.Schema) *Array {
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
