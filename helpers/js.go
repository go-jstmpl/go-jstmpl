package helpers

func ConvertTypeForJS(s string) string {
	conv := map[string]string{
		"array":   "array",
		"boolean": "boolean",
		"integer": "number",
		"number":  "number",
		"object":  "object",
		"string":  "string",
	}
	return conv[s]
}
