package jstmpl

import (
	"fmt"
	"reflect"

	"github.com/lestrrat/go-jsschema"
)

func (a ByClassName) Len() int           { return len(a) }
func (a ByClassName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByClassName) Less(i, j int) bool { return a[i].ClassName() < a[j].ClassName() }

func NewObject(p string, s *schema.Schema) Object {
	return Object{
		Schema:     s,
		Type:       TitleToClassName(s.Title),
		Key:        p,
		IsPrivate:  false,
		PropName:   KeyToPropName(p),
		className:  TitleToClassName(s.Title),
		Properties: []Schema{},
	}
}

func (o Object) ClassName() string {
	return o.className
}

func NewArray(p string, s *schema.Schema) Array {
	return Array{
		Schema:    s,
		Type:      TitleToClassName(s.Title),
		Key:       p,
		IsPrivate: false,
		PropName:  KeyToPropName(p),
		className: TitleToClassName(s.Title),
		Items: &ItemSpec{
			ItemSpec: s.Items,
			Schemas:  make([]Schema, len(s.Items.Schemas)),
		},
	}
}

func NewString(p string, s *schema.Schema) String {
	v := []Validation{}
	if s.MinLength.Initialized {
		v = append(v, MinLength(s.MinLength.Val))
	}
	if s.MaxLength.Initialized {
		v = append(v, MinLength(s.MaxLength.Val))
	}
	return String{
		Schema:      s,
		Type:        "string",
		Key:         p,
		IsPrivate:   true,
		PropName:    KeyToPropName(p),
		className:   TitleToClassName(s.Title),
		Validations: v,
	}
}

func (o String) ClassName() string {
	return o.className
}
func NewNumber(p string, s *schema.Schema) Number {
	return Number{
		Schema:    s,
		Type:      "number",
		Key:       p,
		IsPrivate: true,
		PropName:  KeyToPropName(p),
		className: TitleToClassName(s.Title),
	}
}

func (o Number) ClassName() string {
	return o.className
}

func NewInteger(p string, s *schema.Schema) Integer {
	return Integer{
		Schema:    s,
		Type:      "number",
		Key:       p,
		IsPrivate: true,
		PropName:  KeyToPropName(p),
		className: TitleToClassName(s.Title),
	}
}

func (o Integer) ClassName() string {
	return o.className
}

func NewBoolean(p string, s *schema.Schema) Boolean {
	return Boolean{
		Schema:    s,
		Type:      "boolean",
		Key:       p,
		IsPrivate: true,
		PropName:  KeyToPropName(p),
		className: TitleToClassName(s.Title),
	}
}

func (o Boolean) ClassName() string {
	return o.className
}

func (o Array) ClassName() string {
	return o.className
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
