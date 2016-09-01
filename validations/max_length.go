package validations

import (
	"fmt"

	"github.com/lestrrat/go-jsschema"
)

type MaxLengthValidation struct {
	MaxLength int
}

func NewMaxLengthValidation(s *schema.Schema) (MaxLengthValidation, error) {
	if !s.MaxLength.Initialized {
		return MaxLengthValidation{}, fmt.Errorf("not initialized")
	}
	return MaxLengthValidation{s.MaxLength.Val}, nil
}

func (v MaxLengthValidation) Func() string {
	return "MaxLength"
}

func (v MaxLengthValidation) Args() map[string]interface{} {
	return map[string]interface{}{
		"MaxLength": v.MaxLength,
	}
}
