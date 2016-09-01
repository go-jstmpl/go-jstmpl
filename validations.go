package jstmpl

import (
	"fmt"

	"github.com/lestrrat/go-jsschema"
)

type Validation interface {
	Func() string
	Args() map[string]interface{}
}

type Enum []interface{}

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
