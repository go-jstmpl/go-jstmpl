package helpers

import (
	"fmt"

	schema "github.com/lestrrat/go-jsschema"
)

func ConvertTypeForGo(ts schema.PrimitiveTypes) string {
	if ts.Len() >= 2 {
		return ""
	}

	conv := map[string]string{
		"integer": "int",
		"boolean": "bool",
	}

	for _, t := range ts {
		if conv[t.String()] == "" {
			return t.String()
		}
	}
	return conv[ts[0].String()]
}

func ConvertTagsForGo(n, cn string) string {
	s := "`"
	if n != "" {
		s += fmt.Sprintf("json:\"%s, omitempty\" ", n)
	} else {
		s += fmt.Sprint("json:\"-\" ")
	}

	if cn != "" {
		s += fmt.Sprintf("xorm:\"%s\"", cn)
	} else {
		s += fmt.Sprint("xorm:\"-\"")
	}
	s += "`"
	return s
}

func GetGoTypeData(ts *schema.Schema) (gt string, err error) {
	if ts.Extras["go_type"] == nil {
		return
	}

	s, ok := ts.Extras["go_type"].(string)
	if ok != true {
		err = fmt.Errorf("go_type %v is invalid type", ts.Extras["go_type"])
	}
	gt = s
	return
}

func GetTableData(ts *schema.Schema) (tn string, err error) {
	if ts.Extras["table"] == nil {
		return
	}

	s, ok := ts.Extras["table"].(map[string]interface{})
	if ok != true {
		err = fmt.Errorf("table %v is invalid type", ts.Extras["table"])
	}

	t, ok := s["name"].(string)
	if ok != true {
		err = fmt.Errorf("table[name] %v is invalid type", s["name"])
	}

	tn = t
	return
}

func GetColumnData(ts *schema.Schema) (cn, ct string, err error) {
	if ts.Extras["column"] == nil {
		return
	}

	c, ok := ts.Extras["column"].(map[string]interface{})
	if ok != true {
		err = fmt.Errorf("column %v is invalid type", ts.Extras["column"])
	}

	n, ok := c["name"].(string)
	if ok != true {
		err = fmt.Errorf("column[name] %v is invalid type", c["name"])
	}
	cn = n

	t, ok := c["db_type"].(string)
	if ok != true {
		err = fmt.Errorf("column[db_type] %v is invalid type", c["db_type"])
	}
	ct = t
	return
}
