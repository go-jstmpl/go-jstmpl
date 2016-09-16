package types

import (
	"net/url"

	hschema "github.com/lestrrat/go-jshschema"
)

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

func NewRoot(hs *hschema.HyperSchema) (*Root, error) {
	var u *url.URL
	if hs.Schema.Extras["href"] != nil {
		var err error
		u, err = url.Parse(hs.Schema.Extras["href"].(string))
		if err != nil {
			return nil, err
		}
	}
	return &Root{
		HyperSchema: hs,
		URL:         u,
		Links:       make(LinkList, len(hs.Links)),
	}, nil
}

func (r Root) Title() string {
	return r.HyperSchema.Title
}
