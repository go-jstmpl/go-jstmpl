package types_test

import (
	"testing"

	"github.com/go-jstmpl/go-jstmpl/types"
	hschema "github.com/lestrrat/go-jshschema"
	schema "github.com/lestrrat/go-jsschema"
)

func TestNewRoot(t *testing.T) {
	s, err := types.NewRoot(&hschema.HyperSchema{
		Schema: &schema.Schema{
			Title:       "title",
			Description: "description",
		},
	})
	if err != nil {
		t.Fatalf("fail to NewRoot with %s", err)
	}
	if s.Title() != "title" {
		t.Fatalf("Title expected %s but actual %s", "title", s.Title())
	}
	if s.Description != "description" {
		t.Fatalf("Description expected %s but actual %s", "description", s.Description)
	}
}
