package helpers

import schema "github.com/lestrrat/go-jsschema"

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
