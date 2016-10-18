package validations

import (
	"fmt"

	"github.com/lestrrat/go-jsschema"
)

type MaxItemsValidation struct {
	MaxItems int
}

func NewMaxItemsValidation(s *schema.Schema) (MaxItemsValidation, error) {
	if !s.MaxItems.Initialized {
		return MaxItemsValidation{}, fmt.Errorf("not initialized")
	}
	return MaxItemsValidation{s.MaxItems.Val}, nil
}

func (v MaxItemsValidation) Func() string {
	return "MaxItems"
}

func (v MaxItemsValidation) Args() map[string]interface{} {
	return map[string]interface{}{
		"MaxItems": v.MaxItems,
	}
}
