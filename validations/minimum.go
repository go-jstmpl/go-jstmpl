package validations

import (
	"fmt"

	"github.com/lestrrat/go-jsschema"
)

type MinimumValidation struct {
	Minimum   float64
	Exclusive bool
}

func NewMinimumValidation(s *schema.Schema) (v MinimumValidation, err error) {
	if !s.Minimum.Initialized {
		return MinimumValidation{}, fmt.Errorf("not initialized")
	}
	if s.ExclusiveMinimum.Initialized {
		return MinimumValidation{s.Minimum.Val, s.ExclusiveMinimum.Val}, nil
	}
	return MinimumValidation{s.Minimum.Val, false}, nil
}

func (v MinimumValidation) Func() string {
	return "Minimum"
}

func (v MinimumValidation) Args() map[string]interface{} {
	return map[string]interface{}{
		"Minimum":   v.Minimum,
		"Exclusive": v.Exclusive,
	}
}
