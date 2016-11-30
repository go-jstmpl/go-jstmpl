package helpers_test

import (
	"testing"

	"github.com/go-jstmpl/go-jstmpl/helpers"
)

func TestToLiteralForJS(t *testing.T) {
	type Case struct {
		Description string
		Input       interface{}
		Expected    string
	}

	cases := []Case{
		{
			Description: "bool",
			Input:       true,
			Expected:    "true",
		},
		{
			Description: "int",
			Input:       int(1),
			Expected:    "1",
		},
		{
			Description: "int8",
			Input:       int8(1),
			Expected:    "1",
		},
		{
			Description: "int16",
			Input:       int16(1),
			Expected:    "1",
		},
		{
			Description: "int32",
			Input:       int32(1),
			Expected:    "1",
		},
		{
			Description: "64",
			Input:       int64(1),
			Expected:    "1",
		},
		{
			Description: "float32",
			Input:       float32(1.2),
			Expected:    "1.2",
		},
		{
			Description: "float64",
			Input:       float64(1.2),
			Expected:    "1.2",
		},
		{
			Description: "string",
			Input:       "foo",
			Expected:    "\"foo\"",
		},
	}

	for _, c := range cases {
		actual := helpers.ToLiteralForJS(c.Input)
		if actual != c.Expected {
			t.Errorf("Test with %s: expected %s, but actual %s", c.Description, c.Expected, actual)
		}
	}
}

func TestConvertTypeForJS(t *testing.T) {
	type Case struct {
		Input  string
		Expect string
	}

	cases := []Case{
		{
			Input:  "array",
			Expect: "array",
		},
		{
			Input:  "boolean",
			Expect: "boolean",
		},
		{
			Input:  "number",
			Expect: "number",
		},
		{
			Input:  "object",
			Expect: "object",
		},
		{
			Input:  "string",
			Expect: "string",
		},
		{
			Input:  "integer",
			Expect: "number",
		},
		{
			Input:  "Foo",
			Expect: "Foo",
		},
	}
	for _, c := range cases {
		r := helpers.ConvertTypeForJS(c.Input)
		if r != c.Expect {
			t.Errorf("%s was expected to be %s, but actual %s.", c.Input, c.Expect, r)
		}
	}
}
