package jstmpl_test

import (
	"io"
	"testing"

	"github.com/go-jstmpl/go-jstmpl"
	"github.com/go-jstmpl/go-jstmpl/types"
)

type TestGenerateCase struct {
	Template []byte
	Root     *types.Root
	Pass     bool
	Explain  string
}

func TestGenerater(t *testing.T) {
	gen := jstmpl.NewGenerator()
	tests := []TestGenerateCase{{
		Template: []byte{},
		Root:     &types.Root{},
		Pass:     true,
		Explain:  "pass",
	}}

	var out io.Writer
	for _, test := range tests {
		err := gen.Process(out, test.Root, test.Template)
		if err != nil {
			t.Fatalf("generate error: %s: %v", test.Explain, err)
		}
	}
}
