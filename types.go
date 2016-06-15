package jstmpl

import (
	"fmt"
	"reflect"

	"github.com/lestrrat/go-jsschema"
)

func (a ByClassName) Len() int           { return len(a) }
func (a ByClassName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByClassName) Less(i, j int) bool { return a[i].Title() < a[j].Title() }

func NewObject(p string, s *schema.Schema) *Object {
	return &Object{
		Schema:     s,
		Key:        p,
		IsPrivate:  false,
		Properties: []Schema{},
	}
}

func (o Object) Title() string {
	return o.Schema.Title
}

func NewArray(p string, s *schema.Schema) *Array {
	return &Array{
		Schema:    s,
		Key:       p,
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
		Key:         p,
		IsPrivate:   true,
		Validations: v,
	}
}

func (o String) Title() string {
	return o.Schema.Title
}

func NewNumber(p string, s *schema.Schema) *Number {
	return &Number{
		Schema:    s,
		Type:      "number",
		Key:       p,
		IsPrivate: true,
	}
}

func (o Number) Title() string {
	return o.Schema.Title
}

func NewInteger(p string, s *schema.Schema) *Integer {
	return &Integer{
		Schema:    s,
		Type:      "number",
		Key:       p,
		IsPrivate: true,
	}
}

func (o Integer) Title() string {
	return o.Schema.Title
}

func NewBoolean(p string, s *schema.Schema) *Boolean {
	return &Boolean{
		Schema:    s,
		Type:      "boolean",
		Key:       p,
		IsPrivate: true,
	}
}

func (o Boolean) Title() string {
	return o.Schema.Title
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
