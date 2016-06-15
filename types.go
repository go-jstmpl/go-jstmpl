package jstmpl

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/lestrrat/go-jsschema"
)

func (a ByClassName) Len() int           { return len(a) }
func (a ByClassName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByClassName) Less(i, j int) bool { return a[i].Title() < a[j].Title() }

func NewObject(p string, s *schema.Schema) *Object {
	return &Object{
		Schema:     s,
		Type:       "object",
		key:        p,
		IsPrivate:  false,
		Properties: []Schema{},
	}
}

func (o Object) Title() string {
	return o.Schema.Title
}

func (o Object) Key() string {
	return o.key
}

func (o Object) Example() string {
	e := make(map[string]interface{})
	for _, s := range o.Properties {
		e[s.Key()] = s.Example()
	}
	j, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		log.Printf("cannot marshal example JSON: %s\n", err)
		return ""
	}
	return string(j)
}

func NewArray(p string, s *schema.Schema) *Array {
	return &Array{
		Schema:    s,
		Type:      "array",
		key:       p,
		IsPrivate: false,
		Items: &ItemSpec{
			ItemSpec: s.Items,
			Schemas:  make([]Schema, len(s.Items.Schemas)),
		},
	}
}

func (o Array) Title() string {
	return o.Schema.Title
}

func (o Array) Key() string {
	return o.key
}

func (o Array) Example() string {
	e := []interface{}{}
	for i, s := range o.Items.Schemas {
		e[i] = s.Example()
	}
	j, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		log.Printf("cannot marshal example JSON: %s\n", err)
		return ""
	}
	return string(j)
}

func NewString(p string, s *schema.Schema) *String {
	v := []Validation{}
	if s.MinLength.Initialized {
		v = append(v, MinLength(s.MinLength.Val))
	}
	if s.MaxLength.Initialized {
		v = append(v, MinLength(s.MaxLength.Val))
	}
	return &String{
		Schema:      s,
		Type:        "string",
		key:         p,
		IsPrivate:   true,
		Validations: v,
	}
}

func (o String) Title() string {
	return o.Schema.Title
}

func (o String) Key() string {
	return o.key
}

func (o String) Example() string {
	e := o.Schema.Extras["example"]
	if e != nil {
		return e.(string)
	}
	return ""
}

func NewNumber(p string, s *schema.Schema) *Number {
	return &Number{
		Schema:    s,
		Type:      "number",
		key:       p,
		IsPrivate: true,
	}
}

func (o Number) Title() string {
	return o.Schema.Title
}

func (o Number) Key() string {
	return o.key
}

func (o Number) Example() string {
	e := o.Schema.Extras["example"]
	if e != nil {
		return e.(string)
	}
	return "0"
}

func NewInteger(p string, s *schema.Schema) *Integer {
	return &Integer{
		Schema:    s,
		Type:      "number",
		key:       p,
		IsPrivate: true,
	}
}

func (o Integer) Title() string {
	return o.Schema.Title
}

func (o Integer) Key() string {
	return o.key
}

func (o Integer) Example() string {
	e := o.Schema.Extras["example"]
	if e != nil {
		return e.(string)
	}
	return "0"
}

func NewBoolean(p string, s *schema.Schema) *Boolean {
	return &Boolean{
		Schema:    s,
		Type:      "boolean",
		key:       p,
		IsPrivate: true,
	}
}

func (o Boolean) Title() string {
	return o.Schema.Title
}

func (o Boolean) Key() string {
	return o.key
}

func (o Boolean) Example() string {
	e := o.Schema.Extras["example"]
	if e != nil {
		return e.(string)
	}
	return "false"
}

func (v MinLength) String() string {
	return fmt.Sprintf("%s(%d)", reflect.TypeOf(v).Name(), v)
}

func (v MinLength) Func() string {
	return reflect.TypeOf(v).Name()
}

func (v MinLength) Args() string {
	return fmt.Sprintf("%d", v)
}

func (v MaxLength) String() string {
	return fmt.Sprintf("%s(%d)", reflect.TypeOf(v).Name(), v)
}

func (v MaxLength) Func() string {
	return reflect.TypeOf(v).Name()
}

func (v MaxLength) Args() string {
	return fmt.Sprintf("%d", v)
}
