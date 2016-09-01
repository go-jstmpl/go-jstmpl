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

type Format struct {
	Format string
}

func NewFormat(s *schema.Schema) (Format, error) {
	f := string(s.Format)
	if f == "" {
		return Format{}, fmt.Errorf("not initialized")
	}
	return Format{f}, nil
}

func (v Format) Func() string {
	return "Format"
}

func (v Format) Args() map[string]interface{} {
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

type MaxLength struct {
	MaxLength int
}

func NewMaxLength(s *schema.Schema) (MaxLength, error) {
	if !s.MaxLength.Initialized {
		return MaxLength{}, fmt.Errorf("not initialized")
	}
	return MaxLength{s.MaxLength.Val}, nil
}

func (v MaxLength) Func() string {
	return "MaxLength"
}

func (v MaxLength) Args() map[string]interface{} {
	return map[string]interface{}{
		"MaxLength": v.MaxLength,
	}
}

type MinLength struct {
	MinLength int
}

func NewMinLength(s *schema.Schema) (MinLength, error) {
	if !s.MinLength.Initialized {
		return MinLength{}, fmt.Errorf("not initialized")
	}
	return MinLength{s.MinLength.Val}, nil
}

func (v MinLength) Func() string {
	return "MinLength"
}

func (v MinLength) Args() map[string]interface{} {
	return map[string]interface{}{
		"MinLength": v.MinLength,
	}
}

type Pattern struct {
	Pattern string
}

func NewPattern(s *schema.Schema) (Pattern, error) {
	if s.Pattern == nil {
		return Pattern{}, fmt.Errorf("not initialized")
	}
	return Pattern{s.Pattern.String()}, nil
}

func (v Pattern) Func() string {
	return "Pattern"
}

func (v Pattern) Args() map[string]interface{} {
	return map[string]interface{}{
		"Pattern": v.Pattern,
	}
}
