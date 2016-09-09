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

func GetExtraData(ts *schema.Schema) (gt, cn, ct string, err error) {
	for k, v := range ts.Extras {
		switch k {
		case "go_type":
			s, ok := v.(string)
			if ok != true {
				err = fmt.Errorf("go_type %v is invalid type", v)
			}
			gt = s
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
