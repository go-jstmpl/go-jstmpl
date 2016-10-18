package validations

import (
	"fmt"

	"github.com/lestrrat/go-jsschema"
)

type StringEnumValidation struct {
	Enum []string
}

func NewStringEnumValidation(s *schema.Schema) (StringEnumValidation, error) {
	if len(s.Enum) == 0 {
		return StringEnumValidation{}, fmt.Errorf("not initialized")
	}

	arr := make([]string, len(s.Enum))
	for i, v := range s.Enum {
		arr[i] = v.(string)
	}

	return StringEnumValidation{arr}, nil
}

func (v StringEnumValidation) Func() string {
	return "Enum"
}

func (v StringEnumValidation) Args() map[string]interface{} {
	return map[string]interface{}{
		"Enum": v.Enum,
	}
}

type BooleanEnumValidation struct {
	Enum []bool
}

func NewBooleanEnumValidation(s *schema.Schema) (BooleanEnumValidation, error) {
	if len(s.Enum) == 0 {
		return BooleanEnumValidation{}, fmt.Errorf("not initialized")
	}

	arr := make([]bool, len(s.Enum))
	for i, v := range s.Enum {
		arr[i] = v.(bool)
	}

	return BooleanEnumValidation{arr}, nil
}

func (v BooleanEnumValidation) Func() string {
	return "Enum"
}

func (v BooleanEnumValidation) Args() map[string]interface{} {
	return map[string]interface{}{
		"Enum": v.Enum,
	}
}

type IntegerEnumValidation struct {
	Enum []int
}

func NewIntegerEnumValidation(s *schema.Schema) (IntegerEnumValidation, error) {
	if len(s.Enum) == 0 {
		return IntegerEnumValidation{}, fmt.Errorf("not initialized")
	}

	arr := make([]int, len(s.Enum))
	for i, v := range s.Enum {
		arr[i] = v.(int)
	}

	return IntegerEnumValidation{arr}, nil
}

func (v IntegerEnumValidation) Func() string {
	return "Enum"
}

func (v IntegerEnumValidation) Args() map[string]interface{} {
	return map[string]interface{}{
		"Enum": v.Enum,
	}
}

type NumberEnumValidation struct {
	Enum []string
}

func NewNumberEnumValidation(s *schema.Schema) (NumberEnumValidation, error) {
	if len(s.Enum) == 0 {
		return NumberEnumValidation{}, fmt.Errorf("not initialized")
	}

	arr := make([]string, len(s.Enum))
	for i, v := range s.Enum {
		arr[i] = v.(string)
	}

	return NumberEnumValidation{arr}, nil
}

func (v NumberEnumValidation) Func() string {
	return "Enum"
}

func (v NumberEnumValidation) Args() map[string]interface{} {
	return map[string]interface{}{
		"Enum": v.Enum,
	}
}
