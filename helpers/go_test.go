package helpers_test

import (
	"testing"

	"github.com/go-jstmpl/go-jstmpl/helpers"
	"github.com/lestrrat/go-jsschema"
)

type TestCaseConvertTagsForGo struct {
	ColumnName string
	Name       string
	Expect     string
	Title      string
}

type TestCaseExtras struct {
	Extras           map[string]interface{}
	ExpectColumnName string
	ExpectTableName  string
	ExpectDbType     string
	Title            string
}

type TestCaseConvertArrayForGo struct {
	Input  []string
	Expect string
	Title  string
}

func TestConvertArrayForGo(t *testing.T) {
	tests := []TestCaseConvertArrayForGo{{
		Input:  []string{"foo", "bar"},
		Expect: "[]string{\"foo\",\"bar\"}",
		Title:  "pass test",
	}, {
		Input:  []string{},
		Expect: "[]string{}",
		Title:  "empty test",
	}}
	for _, test := range tests {
		if test.Expect != helpers.ConvertArrayForGo(test.Input) {
			t.Errorf("fail to %s: Expect %s, But %s", test.Title, test.Expect, test.Input)
		}
	}
}

func TestConvertTagsForGo(t *testing.T) {
	tests := []TestCaseConvertTagsForGo{{
		ColumnName: "test_column",
		Name:       "test_name",
		Expect:     "`json:\"test_name, omitempty\" xorm:\"test_column\"`",
		Title:      "pass all column have value test",
	}, {
		ColumnName: "",
		Name:       "test_name",
		Expect:     "`json:\"test_name, omitempty\" xorm:\"-\"`",
		Title:      "pass ColumnName column empty test",
	}, {
		ColumnName: "test_column",
		Name:       "",
		Expect:     "`json:\"-\" xorm:\"test_column\"`",
		Title:      "pass Name column empty test",
	}, {
		ColumnName: "",
		Name:       "",
		Expect:     "`json:\"-\" xorm:\"-\"`",
		Title:      "pass all column empty test",
	}}

	for _, test := range tests {
		s := helpers.ConvertTagsForGo(test.Name, test.ColumnName)
		if test.Expect != s {
			t.Errorf("%s: Expect: %s, Result: %s", test.Title, test.Expect, s)
		}
	}
}

func TestGetExtraData(t *testing.T) {

	tests := []TestCaseExtras{{
		Extras: map[string]interface{}{
			"column": map[string]interface{}{
				"db_type": "db_type_test",
				"name":    "column_name_test",
			}},
		ExpectDbType:     "db_type_test",
		ExpectColumnName: "column_name_test",
		Title:            "pass all column have value test",
	}, {
		Extras: map[string]interface{}{
			"column": map[string]interface{}{
				"db_type": "db_type_test",
				"name":    "column_name_test",
			},
		},
		ExpectDbType:     "db_type_test",
		ExpectColumnName: "column_name_test",
		Title:            "pass column column have value test",
	}, {
		Extras: map[string]interface{}{
			"table": map[string]interface{}{
				"name": "table_name_test",
			},
		},
		ExpectTableName: "table_name_test",
		Title:           "pass table column have value test",
	}}

	for _, test := range tests {
		s := schema.New()
		s.Extras = test.Extras

		cn, ct, err := helpers.GetColumnData(s)
		if err != nil {
			t.Errorf("%s: in GetColumnData, Extras: %s, Error: %s", test.Title, test.Extras, err)
		}

		tn, err := helpers.GetTableData(s)
		if err != nil {
			t.Errorf("%s: in GetTableData, Extras: %s, Error: %s", test.Title, test.Extras, err)
		}

		if cn != test.ExpectColumnName || ct != test.ExpectDbType || tn != test.ExpectTableName {
			t.Errorf("%s: Expect(%s, %s, %s), Result(%s, %s, %s)", test.Title,
				test.ExpectColumnName, test.ExpectDbType, test.ExpectTableName,
				cn, ct, tn)
		}
	}
}
