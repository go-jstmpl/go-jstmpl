package helpers_test

import (
	"testing"

	"github.com/go-jstmpl/go-jstmpl/helpers"
)

func TestAdd(t *testing.T) {
	type Input struct {
		a int
		b int
	}
	type Case struct {
		Input    Input
		Expected int
	}

	cases := []Case{
		{
			Input: Input{
				a: 2,
				b: 1,
			},
			Expected: 3,
		},
		{
			Input: Input{
				a: 1,
				b: 2,
			},
			Expected: 3,
		},
		{
			Input: Input{
				a: 1,
				b: 0,
			},
			Expected: 1,
		},
		{
			Input: Input{
				a: 0,
				b: 1,
			},
			Expected: 1,
		},
		{
			Input: Input{
				a: -1,
				b: 0,
			},
			Expected: -1,
		},
		{
			Input: Input{
				a: 0,
				b: -1,
			},
			Expected: -1,
		},
		{
			Input: Input{
				a: 1,
				b: -1,
			},
			Expected: 0,
		},
		{
			Input: Input{
				a: -1,
				b: 1,
			},
			Expected: 0,
		},
	}

	for _, c := range cases {
		actual := helpers.Add(c.Input.a, c.Input.b)
		if actual != c.Expected {
			t.Errorf("Add(%d, %d) expected %d, but actual %d", c.Input.a, c.Input.b, c.Expected, actual)
		}
	}
}

func TestSub(t *testing.T) {
	type Input struct {
		a int
		b int
	}
	type Case struct {
		Input    Input
		Expected int
	}

	cases := []Case{
		{
			Input: Input{
				a: 2,
				b: 1,
			},
			Expected: 1,
		},
		{
			Input: Input{
				a: 1,
				b: 2,
			},
			Expected: -1,
		},
		{
			Input: Input{
				a: 1,
				b: 0,
			},
			Expected: 1,
		},
		{
			Input: Input{
				a: 0,
				b: 1,
			},
			Expected: -1,
		},
		{
			Input: Input{
				a: -1,
				b: 0,
			},
			Expected: -1,
		},
		{
			Input: Input{
				a: 0,
				b: -1,
			},
			Expected: 1,
		},
		{
			Input: Input{
				a: 1,
				b: -1,
			},
			Expected: 2,
		},
		{
			Input: Input{
				a: -1,
				b: 1,
			},
			Expected: -2,
		},
	}

	for _, c := range cases {
		actual := helpers.Sub(c.Input.a, c.Input.b)
		if actual != c.Expected {
			t.Errorf("Sub(%d, %d) expected %d, but actual %d", c.Input.a, c.Input.b, c.Expected, actual)
		}
	}
}
