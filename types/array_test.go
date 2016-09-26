package types_test

import (
	"testing"

	"github.com/go-jstmpl/go-jstmpl/types"
	schema "github.com/lestrrat/go-jsschema"
)

func TestNewArray(t *testing.T) {
	s := types.NewArray(&types.Context{
		Raw:         &schema.Schema{},
		Validations: map[string]bool{},
	}, &schema.Schema{
		Title:       "title",
		Description: "description",
		Format:      "format",
		Items: &schema.ItemSpec{
			Schemas: schema.SchemaList{},
		},
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

func TestArrayPrivateField(t *testing.T) {
	type Case struct {
		Context  *types.Context
		Schema   *schema.Schema
		Expected bool
	}
	cases := []Case{
		{
			Context: &types.Context{
				Raw: &schema.Schema{},
			},
			Schema: &schema.Schema{
				Items: &schema.ItemSpec{
					Schemas: schema.SchemaList{},
				},
			},
			Expected: false,
		},
		{
			Context: &types.Context{
				Raw: &schema.Schema{},
			},
			Schema: &schema.Schema{
				Extras: map[string]interface{}{"private": true},
				Items: &schema.ItemSpec{
					Schemas: schema.SchemaList{},
				},
			},
			Expected: true,
		},
		{
			Context: &types.Context{
				Raw: &schema.Schema{
					Extras: map[string]interface{}{"private": true},
				},
			},
			Schema: &schema.Schema{
				Items: &schema.ItemSpec{
					Schemas: schema.SchemaList{},
				},
			},
			Expected: true,
		},
		{
			Context: &types.Context{
				Raw: &schema.Schema{
					Extras: map[string]interface{}{"private": true},
				},
			},
			Schema: &schema.Schema{
				Extras: map[string]interface{}{"private": false},
				Items: &schema.ItemSpec{
					Schemas: schema.SchemaList{},
				},
			},
			Expected: false,
		},
	}

	for _, c := range cases {
		s := types.NewArray(c.Context, c.Schema)
		if s.Private != c.Expected {
			t.Errorf("Title expected %t but actual %t", c.Expected, s.Private)
		}
	}
}
