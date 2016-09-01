package validations

import (
	"fmt"

	"github.com/lestrrat/go-jsschema"
)

type MaximumValidation struct {
	Maximum   float64
	Exclusive bool
}

func NewMaximumValidation(s *schema.Schema) (v MaximumValidation, err error) {
	if !s.Maximum.Initialized {
		return MaximumValidation{}, fmt.Errorf("not initialized")
	}
	if s.ExclusiveMaximum.Initialized {
		return MaximumValidation{s.Maximum.Val, s.ExclusiveMaximum.Val}, nil
	}
	return MaximumValidation{s.Maximum.Val, false}, nil
}

func (v MaximumValidation) Func() string {
	return "Maximum"
}

func (v MaximumValidation) Args() map[string]interface{} {
	return map[string]interface{}{
		"Maximum":   v.Maximum,
		"Exclusive": v.Exclusive,
	}
}
