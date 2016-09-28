package types

import (
	"encoding/json"

	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/go-jstmpl/go-jstmpl/validations"
	"github.com/lestrrat/go-jsschema"
)

type String struct {
	*schema.Schema
	NativeType  string `json:"-"`
	ColumnName  string
	ColumnType  string
	Type        string
	Name        string
	key         string
	Private     bool
	Validations []validations.Validation
}

func NewString(ctx *Context, s *schema.Schema) *String {
	vs := []validations.Validation{}
	if v, err := validations.NewFormatValidation(s); err == nil {
		ctx.AddValidation(v)
		vs = append(vs, v)
	}
	if v, err := validations.NewMinLengthValidation(s); err == nil {
		ctx.AddValidation(v)
		vs = append(vs, v)
	}
	if v, err := validations.NewMaxLengthValidation(s); err == nil {
		ctx.AddValidation(v)
		vs = append(vs, v)
	}
	if v, err := validations.NewPatternValidation(s); err == nil {
		ctx.AddValidation(v)
		vs = append(vs, v)
	}

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

	return &String{
		Schema:      s,
		NativeType:  "string",
		Type:        "string",
		ColumnName:  cn,
		ColumnType:  ct,
		Name:        helpers.UpperCamelCase(s.Title),
		key:         ctx.Key,
		Private:     pr,
		Validations: vs,
	}
}

func (o String) MarshalJSON() ([]byte, error) {
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

func (o String) Raw() *schema.Schema {
	return o.Schema
}

func (o String) Title() string {
	return o.Schema.Title
}

func (o String) Format() string {
	return string(o.Schema.Format)
}

func (o String) Key() string {
	return o.key
}

func (o String) ReadOnly() bool {
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

func (o String) Example(withoutReadOnly bool) interface{} {
	v := o.Schema.Extras["example"]
	if v == nil {
		return ""
	}
	return v
}

func (o String) ExampleString() string {
	return helpers.Serialize(o.Example(false))
}
