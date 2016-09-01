package validations

import (
	"fmt"

	"github.com/lestrrat/go-jsschema"
)

type FormatValidation struct {
	Format string
}

func NewFormatValidation(s *schema.Schema) (FormatValidation, error) {
	f := string(s.Format)
	if f == "" {
		return FormatValidation{}, fmt.Errorf("not initialized")
	}
	return FormatValidation{f}, nil
}

func (v FormatValidation) Func() string {
	return "Format"
}

func (v FormatValidation) Args() map[string]interface{} {
	return map[string]interface{}{
		"Format": v.Format,
	}
}
