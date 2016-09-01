package validations

import (
	"fmt"

	"github.com/lestrrat/go-jsschema"
)

type MinLengthValidation struct {
	MinLength int
}

func NewMinLengthValidation(s *schema.Schema) (MinLengthValidation, error) {
	if !s.MinLength.Initialized {
		return MinLengthValidation{}, fmt.Errorf("not initialized")
	}
	return MinLengthValidation{s.MinLength.Val}, nil
}

func (v MinLengthValidation) Func() string {
	return "MinLength"
}

func (v MinLengthValidation) Args() map[string]interface{} {
	return map[string]interface{}{
		"MinLength": v.MinLength,
	}
}
