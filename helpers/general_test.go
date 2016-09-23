package helpers_test

import (
	"testing"

	"github.com/go-jstmpl/go-jstmpl/helpers"
)

type TestCaseJSONPathEscape struct {
	Input  string
	Expect string
	Title  string
}

func TestEscapeJSONPath(t *testing.T) {
	tests := []TestCaseJSONPathEscape{{
		Input:  "/sample/{#/sample/params}",
		Expect: "/sample/:params",
		Title:  "pass in one espace",
	}, {
		Input:  "/sample/{#/sample/params1}/sample/{#/sample/params2}",
		Expect: "/sample/:params1/sample/:params2",
		Title:  "pass in two escape",
	}}

	for _, test := range tests {
		if out := helpers.EscapeJSONPath(test.Input); out != test.Expect {
			t.Errorf("%s: Expect: %s, but Actual: %s", test.Title, test.Expect, out)
		}
	}
}
