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

func GetColumnData(ts *schema.Schema) (cn, ct string, err error) {
	for k, v := range ts.Extras {
		switch k {
		case "column":
			m, ok := v.(map[string]interface{})
			if ok != true {
				err = fmt.Errorf("column %v is invalid type", v)
			}
			for ck, cv := range m {
				switch ck {
				case "name":
					s, ok := cv.(string)
					if ok != true {
						err = fmt.Errorf("column name %v is invalid type", v)
					}
					cn = s
				case "db_type":
					s, ok := cv.(string)
					if ok != true {
						err = fmt.Errorf("column type %v is invalid type", v)
					}
					ct = s
				}
			}
		}
	}
	return
}
