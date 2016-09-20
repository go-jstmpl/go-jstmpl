package types_test

import (
	"testing"

	"github.com/go-jstmpl/go-jstmpl/types"
	schema "github.com/lestrrat/go-jsschema"
)

func TestNewNumber(t *testing.T) {
	type Expected struct {
		Title       string
		Description string
		Format      string
		ReadOnly    bool
	}
	type Case struct {
		Schema   *schema.Schema
		Expected Expected
	}
	cases := []Case{
		Case{
			&schema.Schema{},
			Expected{
				Title:       "",
				Description: "",
				Format:      "",
				ReadOnly:    false,
			},
		},
		Case{
			&schema.Schema{
				Title:       "example title",
				Description: "example description",
				Format:      "example format",
				Extras: map[string]interface{}{
					"readOnly": true,
				},
			},
			Expected{
				Title:       "example title",
				Description: "example description",
				Format:      "example format",
				ReadOnly:    true,
			},
		},
	}
	for _, c := range cases {
		s := types.NewNumber(&types.Context{
			Raw:         &schema.Schema{},
			Validations: map[string]bool{},
		}, c.Schema)
		if s.Title() != c.Expected.Title {
			t.Errorf("Title expected %s but actual %s", c.Expected.Title, s.Title())
		}
		if s.Description != c.Expected.Description {
			t.Errorf("Description expected %s but actual %s", c.Expected.Description, s.Description)
		}
		if s.Format() != c.Expected.Format {
			t.Errorf("Format expected %s but actual %s", c.Expected.Format, s.Format())
		}
		if s.ReadOnly() != c.Expected.ReadOnly {
			t.Errorf("ReadOnly expected %b but actual %b", c.Expected.ReadOnly, s.ReadOnly())
		}
	}
}
