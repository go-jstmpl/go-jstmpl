package types

import "github.com/lestrrat/go-jsschema"

type Schema interface {
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

type ItemSpec struct {
	*schema.ItemSpec
	Schemas []Schema
}
