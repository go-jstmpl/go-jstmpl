package jstmpl

import (
	"net/url"

	"github.com/lestrrat/go-jshschema"
	"github.com/lestrrat/go-jsschema"
)

type Schema interface {
	Raw() *schema.Schema
	Title() string
	Key() string
	Example() interface{}
}

type Root struct {
	*hschema.HyperSchema
	URL                 *url.URL
	Links               LinkList
	Definitions         []Schema
	Properties          []Schema
	Objects             []*Object
	Arrays              []*Array
	Strings             []*String
	Numbers             []*Number
	Integers            []*Integer
	Booleans            []*Boolean
	RequiredValidations []string
}

type Object struct {
	*schema.Schema
	NativeType  string
	Type        string
	Name        string
	key         string
	IsPrivate   bool
	Validations []Validation
	Properties  []Schema
}

type Array struct {
	*schema.Schema
	NativeType  string
	Type        string
	Name        string
	key         string
	IsPrivate   bool
	Validations []Validation
	Item        Schema
	Items       *ItemSpec
}

type String struct {
	*schema.Schema
	NativeType  string
	Type        string
	Name        string
	key         string
	IsPrivate   bool
	Validations []Validation
}

type Number struct {
	*schema.Schema
	NativeType  string
	Type        string
	Name        string
	key         string
	IsPrivate   bool
	Validations []Validation
}

type Integer struct {
	*schema.Schema
	NativeType  string
	Type        string
	Name        string
	key         string
	IsPrivate   bool
	Validations []Validation
}

type Boolean struct {
	*schema.Schema
	NativeType  string
	Type        string
	Name        string
	key         string
	IsPrivate   bool
	Validations []Validation
}

type ItemSpec struct {
	*schema.ItemSpec
	Schemas []Schema
}

type LinkList []*Link

type Link struct {
	hschema.Link
	URL          *url.URL
	Schema       Schema
	TargetSchema Schema
}

type Header struct {
	Key, Value string
}

type Builder struct {
	Schema *schema.Schema
}
