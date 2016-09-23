package types

import (
	"encoding/json"

	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/validations"
	schema "github.com/lestrrat/go-jsschema"
)

type Array struct {
	*schema.Schema
	NativeType  string `json:"-"`
	ColumnName  string
	ColumnType  string
	Type        string
	Name        string
	key         string
	Private     bool
	Properties  []Schema
	Reference   string
	Validations []validations.Validation
	Item        Schema
	Items       *ItemSpec
}

func NewArray(ctx *Context, s *schema.Schema) *Array {
	var cn, ct string

	if s.Extras["column"] != nil {
		cn, ct, _ = helpers.GetColumn(s)
	} else {
		cn, ct, _ = helpers.GetColumn(ctx.Raw)
	}

	return &Array{
		Schema:     s,
		NativeType: "array",
		ColumnName: cn,
		ColumnType: ct,
		Reference:  ctx.Raw.Reference,
		Type:       helpers.SpaceToUpperCamelCase(s.Title),
		Name:       helpers.SpaceToUpperCamelCase(s.Title),
		key:        ctx.Key,
		Private:    false,
		Items: &ItemSpec{
			ItemSpec: s.Items,
			Schemas:  make([]Schema, len(s.Items.Schemas)),
		},
	}
}

func (o Array) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"Title":       o.Title(),
		"Description": o.Description,
		"Type":        o.Type,
		"Name":        o.Name,
		"Required":    o.Required,
		"Validations": o.Validations,
		"Properties":  o.Properties,
		"ColumnName":  o.ColumnName,
		"ColumnType":  o.ColumnType,
		"Reference":   o.Reference,
	})
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
