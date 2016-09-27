package helpers_test

import "testing"
import "github.com/go-jstmpl/go-jstmpl/helpers"

type TestCaseConvertTypeForJs struct {
	Type   string
	Expect string
}

func TestConvertTypeForJS(t *testing.T) {
	cases := []TestCaseConvertTypeForJs{
		{
			Type:   "array",
			Expect: "array",
		},
		{
			Type:   "boolean",
			Expect: "boolean",
		},
		{
			Type:   "number",
			Expect: "number",
		},
		{
			Type:   "object",
			Expect: "object",
		},
		{
			Type:   "string",
			Expect: "string",
		},
		{
			Type:   "integer",
			Expect: "number",
		},
		{
			Type:   "Foo",
			Expect: "Foo",
		},
	}
	for _, c := range cases {
		r := helpers.ConvertTypeForJS(c.Type)
		if r != c.Expect {
			t.Errorf("%s was expected to be %s, but actual %s.", c.Type, c.Expect, r)
		}
	}
}
