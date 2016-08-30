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
	URL         *url.URL
	Links       LinkList
	Definitions []Schema
	Properties  []Schema
	Objects     []*Object
	Arrays      []*Array
	Strings     []*String
	Numbers     []*Number
	Integers    []*Integer
	Booleans    []*Boolean
}

type ByTitle []Schema
type ByKey []Header

type Object struct {
	*schema.Schema
	Type        string
	key         string
	IsPrivate   bool
	Validations []Validation
	Properties  []Schema
}

type Array struct {
	*schema.Schema
	Type        string
	key         string
	IsPrivate   bool
	Validations []Validation
	Items       *ItemSpec
	Item        Schema
}

type String struct {
	*schema.Schema
	Type        string
	key         string
	IsPrivate   bool
	Validations []Validation
}

type Number struct {
	*schema.Schema
	Type        string
	key         string
	IsPrivate   bool
	Validations []Validation
}

type Integer struct {
	*schema.Schema
	Type        string
	key         string
	IsPrivate   bool
	Validations []Validation
}

type Boolean struct {
	*schema.Schema
	Type        string
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
