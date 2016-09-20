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
	IsPrivate   bool
	Properties  []Schema
	Validations []validations.Validation
	Item        Schema
	Items       *ItemSpec
}

func NewArray(ctx *Context, s *schema.Schema) *Array {
	var cn, ct string

	if s.Extras["column"] != nil {
		cn, ct, _ = helpers.GetColumnData(s)
	} else {
		cn, ct, _ = helpers.GetColumnData(ctx.Raw)
	}

	return &Array{
		Schema:     s,
		NativeType: "array",
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

func (o Array) ReadOnly() bool {
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

func (o Array) Example(withoutReadOnly bool) interface{} {
	e := make([]interface{}, len(o.Items.Schemas))
	for i, s := range o.Items.Schemas {
		e[i] = s.Example(withoutReadOnly)
	}
	return e
}

func (o Array) ExampleString() string {
	return helpers.Serialize(o.Example(false))
}
