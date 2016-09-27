package helpers

func ConvertTypeForJS(s string) string {
	v, ok := map[string]string{
		"integer": "number",
	}[s]
	if !ok {
		return s
	}
	return v
}
