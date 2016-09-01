package validations

import (
	"fmt"

	"github.com/lestrrat/go-jsschema"
)

type PatternValidation struct {
	Pattern string
}

func NewPatternValidation(s *schema.Schema) (PatternValidation, error) {
	if s.Pattern == nil {
		return PatternValidation{}, fmt.Errorf("not initialized")
	}
	return PatternValidation{s.Pattern.String()}, nil
}

func (v PatternValidation) Func() string {
	return "Pattern"
}

func (v PatternValidation) Args() map[string]interface{} {
	return map[string]interface{}{
		"Pattern": v.Pattern,
	}
}
