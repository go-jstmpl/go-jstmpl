package types

import (
	"encoding/json"

	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/validations"
	"github.com/lestrrat/go-jsschema"
)

type Object struct {
	*schema.Schema
	NativeType  string
	TableName   string
	ColumnName  string
	ColumnType  string
	Type        string
	Name        string
	key         string
	Private     bool
	Required    []string
	Reference   string
	Validations []validations.Validation
	Properties  []Schema
}

func NewObject(ctx *Context, s *schema.Schema) *Object {
	var pr bool
	if s.Extras["private"] != nil {
		pr, _ = helpers.GetPrivate(s)
	} else {
		pr, _ = helpers.GetPrivate(ctx.Raw)
	}

	var tn string
	if s.Extras["table"] != nil {
		tn, _ = helpers.GetTable(s)
	} else {
		tn, _ = helpers.GetTable(ctx.Raw)
	}

	return &Object{
		Schema:     s,
		NativeType: "object",
		TableName:  tn,
		Type:       helpers.UpperCamelCase(s.Title),
		Name:       helpers.UpperCamelCase(s.Title),
		key:        ctx.Key,
		Required:   ctx.Raw.Required,
		Private:    pr,
		Reference:  ctx.Raw.Reference,
		Properties: []Schema{},
	}
}

func (o Object) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"Title":       o.Title(),
		"Description": o.Description,
		"NativeType":  o.NativeType,
		"Type":        o.Type,
		"Name":        o.Name,
		"Required":    o.Required,
		"Validations": o.Validations,
		"Properties":  o.Properties,
		"TableName":   o.TableName,
		"ColumnName":  o.ColumnName,
		"ColumnType":  o.ColumnType,
		"Reference":   o.Reference,
	})
}

func (o Object) Raw() *schema.Schema {
	return o.Schema
}

func (o Object) Title() string {
	return o.Schema.Title
}

func (o Object) Format() string {
	return string(o.Schema.Format)
}

func (o Object) Key() string {
	return o.key
}

func (o Object) ReadOnly() bool {
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

func (o Object) Example(withoutReadOnly bool) interface{} {
	e := make(map[string]interface{})
	if withoutReadOnly {
		for _, s := range o.Properties {
			if !s.ReadOnly() {
				e[s.Key()] = s.Example(withoutReadOnly)
			}
		}
		return e
	}
	for _, s := range o.Properties {
		e[s.Key()] = s.Example(withoutReadOnly)
	}
	return e
}

func (o Object) ExampleString() string {
	return helpers.Serialize(o.Example(false))
}
