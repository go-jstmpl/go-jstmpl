package helpers_test

import (
	"testing"

	"github.com/go-jstmpl/go-jstmpl/helpers"
)

func TestConvertToStringLiteral(t *testing.T) {
	type Case struct {
		Input    interface{}
		Expected string
		Message  string
	}
	cases := []Case{
		Case{
			Input:    1,
			Expected: "1",
			Message:  "int",
		},
		Case{
			Input:    1.0,
			Expected: "1",
			Message:  "float64",
		},
		Case{
			Input:    true,
			Expected: "true",
			Message:  "bool",
		},
		Case{
			Input:    "string",
			Expected: "\"string\"",
			Message:  "string",
		},
	}
	for _, c := range cases {
		actual := helpers.ToStringLiteral(c.Input)
		if actual != c.Expected {
			t.Errorf("%s: expected '%s', but actual '%s'", c.Message, c.Expected, actual)
		}
	}
}
