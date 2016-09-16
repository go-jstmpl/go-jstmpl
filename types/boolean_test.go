package types_test

import (
	"testing"

	"github.com/go-jstmpl/go-jstmpl/types"
	schema "github.com/lestrrat/go-jsschema"
)

func TestNewBoolean(t *testing.T) {
	s := types.NewBoolean(&types.Context{
		Raw:         &schema.Schema{},
		Validations: map[string]bool{},
	}, &schema.Schema{
		Title:       "title",
		Description: "description",
		Format:      "format",
	})
	if s.Title() != "title" {
		t.Fatalf("Title expected %s but actual %s", "title", s.Title())
	}
	if s.Description != "description" {
		t.Fatalf("Description expected %s but actual %s", "description", s.Description)
	}
	if s.Format() != "format" {
		t.Fatalf("Format expected %s but actual %s", "format", s.Format())
	}
}
