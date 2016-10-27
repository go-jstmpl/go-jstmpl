package helpers_test

import (
	"errors"
	"testing"

	"github.com/go-jstmpl/go-jstmpl/helpers"
)

func TestCatchErrorForString(t *testing.T) {
	type Case struct {
		message  string
		input    string
		expected string
		fn       func(s string) (string, error)
	}
	cases := []Case{
		{
			message:  "no error",
			input:    "foo",
			expected: "foo",
			fn: func(s string) (string, error) {
				return s, nil
			},
		},
		{
			message:  "with error",
			input:    "foo",
			expected: "",
			fn: func(s string) (string, error) {
				return s, errors.New("not good")
			},
		},
	}
	for _, c := range cases {
		actual := helpers.CatchErrorForString(c.fn)(c.input)
		if actual != c.expected {
			t.Errorf("Test with %s: expected %s, but actual %s", c.message, c.expected, actual)
		}
	}
}
