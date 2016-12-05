package helpers_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jstmpl/helpers"
)

func TestToStringLiteral(t *testing.T) {
	type Case struct {
		description string
		input       interface{}
		expected    string
	}
	cases := []Case{
		Case{
			description: "int",
			input:       1,
			expected:    "1",
		},
		Case{
			description: "float64",
			input:       1.0,
			expected:    "1",
		},
		Case{
			description: "bool",
			input:       true,
			expected:    "true",
		},
		Case{
			description: "string",
			input:       "string",
			expected:    "\"string\"",
		},
	}
	for _, c := range cases {
		actual := helpers.ToStringLiteral(c.input)
		if actual != c.expected {
			t.Errorf("%s: expected '%s', but actual '%s'", c.description, c.expected, actual)
		}
	}
}

func TestSlice(t *testing.T) {
	type Case struct {
		description string
		input       []interface{}
		expected    []interface{}
	}

	cases := []Case{
		{
			description: "string",
			input:       []interface{}{"foo", "bar", "baz"},
			expected:    []interface{}{"foo", "bar", "baz"},
		},
		{
			description: "int",
			input:       []interface{}{0, 1, 2},
			expected:    []interface{}{0, 1, 2},
		},
		{
			description: "bool",
			input:       []interface{}{false, true},
			expected:    []interface{}{false, true},
		},
		{
			description: "mixed",
			input:       []interface{}{"foo", 1, true},
			expected:    []interface{}{"foo", 1, true},
		},
	}

	for _, c := range cases {
		actual := helpers.Slice(c.input...)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("%s: expected %+v, but actual %+v", c.description, c.expected, actual)
		}
	}
}

func TestIn(t *testing.T) {
	type Input struct {
		element interface{}
		slice   []interface{}
	}
	type Case struct {
		description string
		input       Input
		expected    bool
	}

	cases := []Case{
		{
			description: "existed string in strings",
			input: Input{
				element: "bar",
				slice:   []interface{}{"foo", "bar", "baz"},
			},
			expected: true,
		},
		{
			description: "non-existed string in strings",
			input: Input{
				element: "qux",
				slice:   []interface{}{"foo", "bar", "baz"},
			},
			expected: false,
		},
		{
			description: "existed int in ints",
			input: Input{
				element: 1,
				slice:   []interface{}{0, 1, 2},
			},
			expected: true,
		},
		{
			description: "non-existed int in ints",
			input: Input{
				element: 3,
				slice:   []interface{}{0, 1, 2},
			},
			expected: false,
		},
		{
			description: "existed bool in bools",
			input: Input{
				element: false,
				slice:   []interface{}{false},
			},
			expected: true,
		},
		{
			description: "non-existed bool in bools",
			input: Input{
				element: true,
				slice:   []interface{}{false},
			},
			expected: false,
		},
		{
			description: "existed string in mixed",
			input: Input{
				element: "foo",
				slice:   []interface{}{"foo", 1, true},
			},
			expected: true,
		},
		{
			description: "existed int in mixed",
			input: Input{
				element: 1,
				slice:   []interface{}{"foo", 1, true},
			},
			expected: true,
		},
		{
			description: "existed bool in mixed",
			input: Input{
				element: true,
				slice:   []interface{}{"foo", 1, true},
			},
			expected: true,
		},
		{
			description: "non-existed string in mixed",
			input: Input{
				element: "baz",
				slice:   []interface{}{"foo", 1, true},
			},
			expected: false,
		},
		{
			description: "non-existed int in mixed",
			input: Input{
				element: 0,
				slice:   []interface{}{"foo", 1, true},
			},
			expected: false,
		},
		{
			description: "non-existed bool in mixed",
			input: Input{
				element: false,
				slice:   []interface{}{"foo", 1, true},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		actual := helpers.In(c.input.element, c.input.slice)
		if actual != c.expected {
			t.Errorf("Test with %s: expected %t, but actual %t", c.description, c.expected, actual)
		}
	}
}

func TestInMapKeys(t *testing.T) {
	type Input struct {
		element interface{}
		slice   []interface{}
	}
	type Case struct {
		description string
		input       Input
		expected    bool
	}

	cases := []Case{
		{
			description: "existed string in keys",
			input: Input{
				element: "foo",
				slice: []interface{}{
					map[string]interface{}{
						"foo": 1,
						"bar": 5,
					},
				},
			},
			expected: true,
		},
		{
			description: "existed string in 2nd map of keys",
			input: Input{
				element: "foo",
				slice: []interface{}{
					map[string]interface{}{
						"ok": "ok",
					},
					map[string]interface{}{
						"foo": 1,
						"bar": 5,
					},
				},
			},
			expected: true,
		},
		{
			description: "invalid type of slice element",
			input: Input{
				element: "foo",
				slice: []interface{}{
					map[interface{}]interface{}{
						"foo": 1,
						"bar": 5,
					},
				},
			},
			expected: false,
		},
		{
			description: "non-existed string in strings",
			input: Input{
				element: "qux",
			},
			expected: false,
		},
	}

	for _, c := range cases {
		actual := helpers.InMapKeys(c.input.element, c.input.slice)
		if actual != c.expected {
			t.Errorf("Test with %s: expected %t, but actual %t", c.description, c.expected, actual)
		}
	}
}
