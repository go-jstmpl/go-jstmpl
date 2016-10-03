package jstmpl_test

import (
	"testing"

	"github.com/go-jstmpl/go-jstmpl"
	"github.com/go-jstmpl/go-jstmpl/types"
)

type TestGenerateCase struct {
	Root      *types.Root
	Template  string
	Extension string
	Expected  string
	Message   string
}

func TestGenerater(t *testing.T) {
	gen := jstmpl.NewGenerator()
	cases := []TestGenerateCase{
		{
			Root:      &types.Root{},
			Template:  "package main\n\nimport \"fmt\"\n\nfunc main () {\n    fmt.Println(\"foo\")\n}",
			Extension: ".go",
			Expected:  "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"foo\")\n}\n",
			Message:   "Format with Go",
		},
	}

	for _, c := range cases {
		b, err := gen.Process(c.Root, []byte(c.Template), c.Extension)
		if err != nil {
			t.Fatalf("%s: returns error %+v", c.Message, err)
		}
		if actual := string(b); actual != c.Expected {
			t.Errorf("%s: \nexpected\n-----\n%s\n-----\nactual\n-----\n%s\n-----", c.Message, c.Expected, actual)
		}
	}
}
