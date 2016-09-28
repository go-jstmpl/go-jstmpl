package helpers_test

import (
	"testing"

	"github.com/go-jstmpl/go-jstmpl/helpers"
)

func TestLowerSnakeCase(t *testing.T) {
	type Case struct {
		Input    string
		Expected string
		Message  string
	}
	cases := []Case{
		Case{
			Input:    "lower spaced string",
			Expected: "lower_spaced_string",
			Message:  "lower spaced string",
		},
		Case{
			Input:    "Upper Spaced String",
			Expected: "upper_spaced_string",
			Message:  "upper spaced string",
		},
		Case{
			Input:    "lower_snake_case",
			Expected: "lower_snake_case",
			Message:  "lower snake case",
		},
		Case{
			Input:    "Upper_Snake_Case",
			Expected: "upper_snake_case",
			Message:  "upper snake case",
		},
		Case{
			Input:    "lower-hyphens",
			Expected: "lower_hyphens",
			Message:  "lower hyphens",
		},
		Case{
			Input:    "Upper-Hyphens",
			Expected: "upper_hyphens",
			Message:  "upper hyphens",
		},
		Case{
			Input:    "lowerCamelCase",
			Expected: "lower_camel_case",
			Message:  "lower camel case",
		},
		Case{
			Input:    "UpperCamelCase",
			Expected: "upper_camel_case",
			Message:  "upper camel case",
		},
		Case{
			Input:    "_mixed veryUgly_Case_",
			Expected: "_mixed_very_ugly_case_",
			Message:  "mixed very ugly case",
		},
	}
	for _, c := range cases {
		actual := helpers.LowerSnakeCase(c.Input)
		if actual != c.Expected {
			t.Errorf("%s: expected '%s', but actual '%s'", c.Message, c.Expected, actual)
		}
	}
}

func TestUpperSnakeCase(t *testing.T) {
	type Case struct {
		Input    string
		Expected string
		Message  string
	}
	cases := []Case{
		Case{
			Input:    "lower spaced string",
			Expected: "LOWER_SPACED_STRING",
			Message:  "lower spaced string",
		},
		Case{
			Input:    "Upper Spaced String",
			Expected: "UPPER_SPACED_STRING",
			Message:  "upper spaced string",
		},
		Case{
			Input:    "lower_snake_case",
			Expected: "LOWER_SNAKE_CASE",
			Message:  "lower snake case",
		},
		Case{
			Input:    "Upper_Snake_Case",
			Expected: "UPPER_SNAKE_CASE",
			Message:  "upper snake case",
		},
		Case{
			Input:    "lower-hyphens",
			Expected: "LOWER_HYPHENS",
			Message:  "lower hyphens",
		},
		Case{
			Input:    "Upper-Hyphens",
			Expected: "UPPER_HYPHENS",
			Message:  "upper hyphens",
		},
		Case{
			Input:    "lowerCamelCase",
			Expected: "LOWER_CAMEL_CASE",
			Message:  "lower camel case",
		},
		Case{
			Input:    "UpperCamelCase",
			Expected: "UPPER_CAMEL_CASE",
			Message:  "upper camel case",
		},
		Case{
			Input:    "_mixed veryUgly_Case_",
			Expected: "_MIXED_VERY_UGLY_CASE_",
			Message:  "mixed very ugly case",
		},
	}
	for _, c := range cases {
		actual := helpers.UpperSnakeCase(c.Input)
		if actual != c.Expected {
			t.Errorf("%s: expected '%s', but actual '%s'", c.Message, c.Expected, actual)
		}
	}
}

func TestLowerCamelCase(t *testing.T) {
	type Case struct {
		Input    string
		Expected string
		Message  string
	}
	cases := []Case{
		Case{
			Input:    "lower spaced string",
			Expected: "lowerSpacedString",
			Message:  "lower spaced string",
		},
		Case{
			Input:    "Upper Spaced String",
			Expected: "upperSpacedString",
			Message:  "upper spaced string",
		},
		Case{
			Input:    "lower_snake_case",
			Expected: "lowerSnakeCase",
			Message:  "lower snake case",
		},
		Case{
			Input:    "Upper_Snake_Case",
			Expected: "upperSnakeCase",
			Message:  "upper snake case",
		},
		Case{
			Input:    "lower-hyphens",
			Expected: "lowerHyphens",
			Message:  "lower hyphens",
		},
		Case{
			Input:    "Upper-Hyphens",
			Expected: "upperHyphens",
			Message:  "upper hyphens",
		},
		Case{
			Input:    "lowerCamelCase",
			Expected: "lowerCamelCase",
			Message:  "lower camel case",
		},
		Case{
			Input:    "UpperCamelCase",
			Expected: "upperCamelCase",
			Message:  "upper camel case",
		},
		Case{
			Input:    "_mixed veryUgly_Case_",
			Expected: "mixedVeryUglyCase",
			Message:  "mixed very ugly case",
		},
	}
	for _, c := range cases {
		actual := helpers.LowerCamelCase(c.Input)
		if actual != c.Expected {
			t.Errorf("%s: expected '%s', but actual '%s'", c.Message, c.Expected, actual)
		}
	}
}

func TestUpperCamelCase(t *testing.T) {
	type Case struct {
		Input    string
		Expected string
		Message  string
	}
	cases := []Case{
		Case{
			Input:    "lower spaced string",
			Expected: "LowerSpacedString",
			Message:  "lower spaced string",
		},
		Case{
			Input:    "Upper Spaced String",
			Expected: "UpperSpacedString",
			Message:  "upper spaced string",
		},
		Case{
			Input:    "lower_snake_case",
			Expected: "LowerSnakeCase",
			Message:  "lower snake case",
		},
		Case{
			Input:    "Upper_Snake_Case",
			Expected: "UpperSnakeCase",
			Message:  "upper snake case",
		},
		Case{
			Input:    "lower-hyphens",
			Expected: "LowerHyphens",
			Message:  "lower hyphens",
		},
		Case{
			Input:    "Upper-Hyphens",
			Expected: "UpperHyphens",
			Message:  "upper hyphens",
		},
		Case{
			Input:    "lowerCamelCase",
			Expected: "LowerCamelCase",
			Message:  "lower camel case",
		},
		Case{
			Input:    "UpperCamelCase",
			Expected: "UpperCamelCase",
			Message:  "upper camel case",
		},
		Case{
			Input:    "_mixed veryUgly_Case_",
			Expected: "MixedVeryUglyCase",
			Message:  "mixed very ugly case",
		},
	}
	for _, c := range cases {
		actual := helpers.UpperCamelCase(c.Input)
		if actual != c.Expected {
			t.Errorf("%s: expected '%s', but actual '%s'", c.Message, c.Expected, actual)
		}
	}
}
