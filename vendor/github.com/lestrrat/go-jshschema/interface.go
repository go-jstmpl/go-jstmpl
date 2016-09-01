package hschema

import "github.com/lestrrat/go-jsschema"

type HyperSchema struct {
	*schema.Schema
	PathStart string   `json:"pathStart,omitempty"`
	Links     LinkList `json:"links,omitempty"`
	Media     Media    `json:"media,omitempty"`
}

type LinkList []*Link
type Link struct {
	parent       *HyperSchema
	Href         string         `json:"href"`
	Rel          string         `json:"rel"`
	Title        string         `json:"title,omitempty"`
	TargetSchema *schema.Schema `json:"targetSchema,omitempty"`
	MediaType    string         `json:"mediaType,omitempty"`
	Method       string         `json:"method,omitempty"`
	EncType      string         `json:"encType,omitempty"`
	Schema       *schema.Schema `json:"schema,omitempty"`
	Extras      map[string]interface{}
}

type Media struct {
	Type           string `json:"type"`
	BinaryEncoding string `json:"binaryEncoding"`
}
