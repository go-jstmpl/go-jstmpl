package types

import (
	"net/url"

	hschema "github.com/lestrrat/go-jshschema"
	"github.com/lestrrat/go-jsschema"
)

type Schema interface {
	Raw() *schema.Schema
	Title() string
	Key() string
	Example() interface{}
}

type SchemasByKey []Schema

func (a SchemasByKey) Len() int           { return len(a) }
func (a SchemasByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SchemasByKey) Less(i, j int) bool { return a[i].Key() < a[j].Key() }

type ByTitle []Schema

func (a ByTitle) Len() int           { return len(a) }
func (a ByTitle) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTitle) Less(i, j int) bool { return a[i].Title() < a[j].Title() }

type Root struct {
	HyperSchema         *hschema.HyperSchema `json:"-"`
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

type ItemSpec struct {
	*schema.ItemSpec
	Schemas []Schema
}
