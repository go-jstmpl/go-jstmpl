package validations

import (
	"fmt"

	"github.com/lestrrat/go-jsschema"
)

type MinItemsValidation struct {
	MinItems int
}

func NewMinItemsValidation(s *schema.Schema) (MinItemsValidation, error) {
	if !s.MinItems.Initialized {
		return MinItemsValidation{}, fmt.Errorf("not initialized")
	}
	return MinItemsValidation{s.MinItems.Val}, nil
}

func (v MinItemsValidation) Func() string {
	return "MinItems"
}

func (v MinItemsValidation) Args() map[string]interface{} {
	return map[string]interface{}{
		"MinItems": v.MinItems,
	}
}
