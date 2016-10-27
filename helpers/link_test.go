package helpers_test

import (
	"reflect"
	"testing"

	"github.com/go-jstmpl/go-jstmpl/helpers"
)

func TestBuilURLToken(t *testing.T) {
	type Case struct {
		message  string
		input    string
		expected []helpers.Token
	}
	cases := []Case{
		{
			message: "empty",
			input:   "",
			expected: []helpers.Token{
				helpers.Chunk(""),
			},
		},
		{
			message: "a slash",
			input:   "/",
			expected: []helpers.Token{
				helpers.Chunk(""),
				helpers.Chunk(""),
			},
		},
		{
			message: "a chunk",
			input:   "foo",
			expected: []helpers.Token{
				helpers.Chunk("foo"),
			},
		},
		{
			message: "chunks",
			input:   "foo/bar/baz",
			expected: []helpers.Token{
				helpers.Chunk("foo"),
				helpers.Chunk("bar"),
				helpers.Chunk("baz"),
			},
		},
		{
			message: "a slashed chunk",
			input:   "/foo",
			expected: []helpers.Token{
				helpers.Chunk(""),
				helpers.Chunk("foo"),
			},
		},
		{
			message: "slashed chunks",
			input:   "/foo/bar/baz",
			expected: []helpers.Token{
				helpers.Chunk(""),
				helpers.Chunk("foo"),
				helpers.Chunk("bar"),
				helpers.Chunk("baz"),
			},
		},
		{
			message: "a ref",
			input:   "{#/definitions/foo}",
			expected: []helpers.Token{
				helpers.Ref("#/definitions/foo"),
			},
		},
		{
			message: "refs",
			input:   "{#/definitions/foo}/{#/definitions/bar}/{#/definitions/baz}",
			expected: []helpers.Token{
				helpers.Ref("#/definitions/foo"),
				helpers.Ref("#/definitions/bar"),
				helpers.Ref("#/definitions/baz"),
			},
		},
		{
			message: "a slashed ref",
			input:   "/{#/definitions/foo}",
			expected: []helpers.Token{
				helpers.Chunk(""),
				helpers.Ref("#/definitions/foo"),
			},
		},
		{
			message: "slashed refs",
			input:   "/{#/definitions/foo}/{#/definitions/bar}/{#/definitions/baz}",
			expected: []helpers.Token{
				helpers.Chunk(""),
				helpers.Ref("#/definitions/foo"),
				helpers.Ref("#/definitions/bar"),
				helpers.Ref("#/definitions/baz"),
			},
		},
		{
			message: "mixed",
			input:   "foo/{#/definitions/bar}/baz",
			expected: []helpers.Token{
				helpers.Chunk("foo"),
				helpers.Ref("#/definitions/bar"),
				helpers.Chunk("baz"),
			},
		},
		{
			message: "slashed mixed",
			input:   "/foo/{#/definitions/bar}/baz",
			expected: []helpers.Token{
				helpers.Chunk(""),
				helpers.Chunk("foo"),
				helpers.Ref("#/definitions/bar"),
				helpers.Chunk("baz"),
			},
		},
	}
	for _, c := range cases {
		actual, err := helpers.BuildURLToken(c.input)
		if err != nil {
			t.Errorf("fail to parse: %s", err)
			continue
		}
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("Test with %s: expected %+v, but actual %+v", c.message, c.expected, actual)
		}
	}
}

func TestToParamsPathLikeGorilla(t *testing.T) {
	type Case struct {
		message  string
		input    string
		expected string
	}
	cases := []Case{
		{
			message:  "empty",
			input:    "",
			expected: "",
		},
		{
			message:  "a slash",
			input:    "/",
			expected: "/",
		},
		{
			message:  "a chunk",
			input:    "foo",
			expected: "foo",
		},
		{
			message:  "chunks",
			input:    "foo/bar/baz",
			expected: "foo/bar/baz",
		},
		{
			message:  "a slashed chunk",
			input:    "/foo",
			expected: "/foo",
		},
		{
			message:  "slashed chunks",
			input:    "/foo/bar/baz",
			expected: "/foo/bar/baz",
		},
		{
			message:  "a ref",
			input:    "{#/definitions/foo}",
			expected: "{foo}",
		},
		{
			message:  "refs",
			input:    "{#/definitions/foo}/{#/definitions/bar}/{#/definitions/baz}",
			expected: "{foo}/{bar}/{baz}",
		},
		{
			message:  "a slashed ref",
			input:    "/{#/definitions/foo}",
			expected: "/{foo}",
		},
		{
			message:  "slashed refs",
			input:    "/{#/definitions/foo}/{#/definitions/bar}/{#/definitions/baz}",
			expected: "/{foo}/{bar}/{baz}",
		},
		{
			message:  "mixed",
			input:    "foo/{#/definitions/bar}/baz",
			expected: "foo/{bar}/baz",
		},
		{
			message:  "slashed mixed",
			input:    "/foo/{#/definitions/bar}/baz",
			expected: "/foo/{bar}/baz",
		},
	}
	for _, c := range cases {
		actual, err := helpers.ToPathLikeGorilla(c.input)
		if err != nil {
			t.Errorf("fail to parse: %s", err)
			continue
		}
		if actual != c.expected {
			t.Errorf("Test with %s: expected '%s', but actual '%s'", c.message, c.expected, actual)
		}
	}
}

func TestToParamsPathLikeSinatra(t *testing.T) {
	type Case struct {
		message  string
		input    string
		expected string
	}
	cases := []Case{
		{
			message:  "empty",
			input:    "",
			expected: "",
		},
		{
			message:  "a slash",
			input:    "/",
			expected: "/",
		},
		{
			message:  "a chunk",
			input:    "foo",
			expected: "foo",
		},
		{
			message:  "chunks",
			input:    "foo/bar/baz",
			expected: "foo/bar/baz",
		},
		{
			message:  "a slashed chunk",
			input:    "/foo",
			expected: "/foo",
		},
		{
			message:  "slashed chunks",
			input:    "/foo/bar/baz",
			expected: "/foo/bar/baz",
		},
		{
			message:  "a ref",
			input:    "{#/definitions/foo}",
			expected: ":foo",
		},
		{
			message:  "refs",
			input:    "{#/definitions/foo}/{#/definitions/bar}/{#/definitions/baz}",
			expected: ":foo/:bar/:baz",
		},
		{
			message:  "a slashed ref",
			input:    "/{#/definitions/foo}",
			expected: "/:foo",
		},
		{
			message:  "slashed refs",
			input:    "/{#/definitions/foo}/{#/definitions/bar}/{#/definitions/baz}",
			expected: "/:foo/:bar/:baz",
		},
		{
			message:  "mixed",
			input:    "foo/{#/definitions/bar}/baz",
			expected: "foo/:bar/baz",
		},
		{
			message:  "slashed mixed",
			input:    "/foo/{#/definitions/bar}/baz",
			expected: "/foo/:bar/baz",
		},
	}
	for _, c := range cases {
		actual, err := helpers.ToPathLikeSinatra(c.input)
		if err != nil {
			t.Errorf("fail to parse: %s", err)
			continue
		}
		if actual != c.expected {
			t.Errorf("Test with %s: expected '%s', but actual '%s'", c.message, c.expected, actual)
		}
	}
}
