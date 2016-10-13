package helpers_test

import (
	"testing"

	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/lestrrat/go-jsschema"
)

type TestCaseTag struct {
	Tag     string
	Expect  string
	Message string
}

type TestCaseExtras struct {
	Extras             map[string]interface{}
	ExpectColumnName   string
	ExpectTableName    string
	ExpectDbType       string
	ExpectPrivateState bool
	Message            string
}
type TestCaseConvertArrayForGo struct {
	Input   []string
	Expect  string
	Message string
}

func TestConvertArrayForGo(t *testing.T) {
	cases := []TestCaseConvertArrayForGo{{
		Input:   []string{"foo_bar", "bar"},
		Expect:  "[]string{\"FooBar\",\"Bar\"}",
		Message: "pass test",
	}, {
		Input:   []string{},
		Expect:  "[]string{}",
		Message: "empty test",
	}}
	for _, c := range cases {
		if c.Expect != helpers.ConvertArrayForGo(c.Input) {
			t.Errorf("fail to %s: Expect %s, But %s", c.Message, c.Expect, c.Input)
		}
	}
}

func TestConvertJSONTagForGo(t *testing.T) {
	cases := []TestCaseTag{
		{
			Tag:     "",
			Expect:  "json:\"-\"",
			Message: "Fail to convert empty",
		},
		{
			Tag:     "-",
			Expect:  "json:\"-\"",
			Message: "Fail to convert hypen",
		},
		{
			Tag:     "foo",
			Expect:  "json:\"foo,omitempty\"",
			Message: "Fail to convert string",
		},
	}

	for _, c := range cases {
		s := helpers.ConvertJSONTagForGo(c.Tag)
		if c.Expect != s {
			t.Errorf("%s: Expect: %s, Result: %s", c.Message, c.Expect, s)
		}
	}
}

func TestConvertXORMTagForGo(t *testing.T) {
	cases := []TestCaseTag{
		{
			Tag:     "",
			Expect:  "xorm:\"-\"",
			Message: "Fail to convert empty",
		},
		{
			Tag:     "-",
			Expect:  "xorm:\"-\"",
			Message: "Fail to convert hyphen",
		},
		{
			Tag:     "foo",
			Expect:  "xorm:\"foo\"",
			Message: "Fail to convert string",
		},
	}

	for _, c := range cases {
		s := helpers.ConvertXORMTagForGo(c.Tag)
		if c.Expect != s {
			t.Errorf("%s: Expect: %s, Result: %s", c.Message, c.Expect, s)
		}
	}
}

func TestGetExtraData(t *testing.T) {

	cases := []TestCaseExtras{{
		Extras: map[string]interface{}{
			"column": map[string]interface{}{
				"db_type": "db_type_test",
				"name":    "column_name_test",
			}},
		ExpectDbType:     "db_type_test",
		ExpectColumnName: "column_name_test",
		Message:          "pass all column have value test",
	}, {
		Extras: map[string]interface{}{
			"column": map[string]interface{}{
				"db_type": "db_type_test",
				"name":    "column_name_test",
			},
		},
		ExpectDbType:     "db_type_test",
		ExpectColumnName: "column_name_test",
		Message:          "pass column column have value test",
	}, {
		Extras: map[string]interface{}{
			"table": map[string]interface{}{
				"name": "table_name_test",
			},
		},
		ExpectTableName: "table_name_test",
		Message:         "pass table column have value test",
	}}

	for _, c := range cases {
		s := schema.New()
		s.Extras = c.Extras

		cn, ct, err := helpers.GetColumn(s)
		if err != nil {
			t.Errorf("%s: in GetColumnData, Extras: %s, Error: %s", c.Message, c.Extras, err)
		}

		tn, err := helpers.GetTable(s)
		if err != nil {
			t.Errorf("%s: in GetTableData, Extras: %s, Error: %s", c.Message, c.Extras, err)
		}

		if cn != c.ExpectColumnName || ct != c.ExpectDbType || tn != c.ExpectTableName {
			t.Errorf("%s: Expect(%s, %s, %s), Result(%s, %s, %s)", c.Message,
				c.ExpectColumnName, c.ExpectDbType, c.ExpectTableName,
				cn, ct, tn)
		}
	}
}
