package jstmpl

import (
	"net/url"

	"github.com/lestrrat/go-jshschema"
	"github.com/lestrrat/go-jsschema"
)

type Root struct {
	URL         *url.URL
	Definitions map[string]*Schema
	Links       LinkList
	Objects     []*Object
	Arrays      []*Array
	Strings     []*String
	Numbers     []*Number
	Integers    []*Integer
	Booleans    []*Boolean
}

type ByClassName []Schema

type Schema interface {
	Title() string
}

type Object struct {
	*schema.Schema
	Type       string
	Key        string
	IsPrivate  bool
	Properties []Schema
}

type Array struct {
	*schema.Schema
	Type      string
	Key       string
	IsPrivate bool
	Items     *ItemSpec
	Item      Schema
}

type String struct {
	*schema.Schema
	Type        string
	Key         string
	IsPrivate   bool
	Validations []Validation
}

type Validation interface {
	String() string
	Func() string
	Args() string
}

type Number struct {
	*schema.Schema
	Type        string
	Key         string
	IsPrivate   bool
	Validations []Validation
}

type Integer struct {
	*schema.Schema
	Type        string
	Key         string
	IsPrivate   bool
	Validations []Validation
}

type Boolean struct {
	*schema.Schema
	Type        string
	Key         string
	IsPrivate   bool
	Validations []Validation
}

type MinLength int

type MaxLength int

type ItemSpec struct {
	*schema.ItemSpec
	Schemas []Schema
}

type LinkList []*Link

type Link struct {
	hschema.Link
	Schema       Schema
	TargetSchema Schema
}

type Builder struct {
	Schema *schema.Schema
}
