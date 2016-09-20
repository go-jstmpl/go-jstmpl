package types

import (
	"encoding/json"

	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/validations"
	"github.com/lestrrat/go-jsschema"
)

type Object struct {
	*schema.Schema
	NativeType  string `json:"-"`
	TableName   string
	ColumnName  string
	ColumnType  string
	Type        string
	Name        string
	key         string
	IsPrivate   bool `json:"-"`
	Required    []string
	Validations []validations.Validation
	Properties  []Schema
}

func NewObject(ctx *Context, s *schema.Schema) *Object {
	var tn string
	if s.Extras["table"] != nil {
		tn, _ = helpers.GetTableData(s)
	} else {
		tn, _ = helpers.GetTableData(ctx.Raw)
	}
	return &Object{
		Schema:     s,
		NativeType: "object",
		TableName:  tn,
		Type:       helpers.SpaceToUpperCamelCase(s.Title),
		Name:       helpers.SpaceToUpperCamelCase(s.Title),
		key:        ctx.Key,
		Required:   ctx.Raw.Required,
		IsPrivate:  false,
		Properties: []Schema{},
	}
}

func (o Object) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"Title":       o.Title(),
		"Description": o.Description,
		"Type":        o.Type,
		"Name":        o.Name,
		"Required":    o.Required,
		"Validations": o.Validations,
		"Properties":  o.Properties,
		"TableName":   o.TableName,
		"ColumnName":  o.ColumnName,
		"ColumnType":  o.ColumnType,
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
