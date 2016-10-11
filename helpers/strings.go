package helpers

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func isAlphabetical(c rune) bool {
	return isLower(c) || isUpper(c)
}

func isUpper(c rune) bool {
	return 'A' <= c && 'Z' >= c
}

func isLower(c rune) bool {
	return 'a' <= c && 'z' >= c
}

func isNumeric(c rune) bool {
	return '0' <= c && '9' >= c
}

func LowerSnakeCase(s string) string {
	return strings.ToLower(SnakeCase(s))
}

func UpperSnakeCase(s string) string {
	return strings.ToUpper(SnakeCase(s))
}

func SnakeCase(s string) string {
	return Chop(s, '_')
}

func Chop(s string, d rune) string {
	var b bytes.Buffer
	for i, c := range s {
		switch {
		case isUpper(c):
			if i > 0 {
				p := rune(s[i-1])
				if isLower(p) {
					b.WriteRune(d)
				}
			}
			b.WriteRune(c)
		case !(isAlphabetical(c) || isNumeric(c)):
			b.WriteRune(d)
		default:
			b.WriteRune(c)
		}
	}
	return b.String()
}

type convert func(rune) rune

func LowerCamelCase(s string) string {
	return CamelCase(s, unicode.ToLower)
}

func UpperCamelCase(s string) string {
	return CamelCase(s, unicode.ToUpper)
}

func CamelCase(s string, fn convert) string {
	var b bytes.Buffer
	first := true
	apply := false
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		switch {
		case !(isAlphabetical(c) || isNumeric(c)):
			apply = true
		case first:
			first = false
			apply = false
			b.WriteRune(fn(c))
		case apply:
			apply = false
			b.WriteRune(unicode.ToUpper(c))
		default:
			apply = false
			b.WriteRune(c)
		}
	}
	return b.String()
}

// SpaceToUpperCamelCase is deprecated
func SpaceToUpperCamelCase(s string) string {
	fmt.Printf("Warning: helpers.SpaceToUpperCamelCase is deprecated\n")
	if s == "" {
		return ""
	}
	buf := bytes.Buffer{}
	for _, p := range rspace.Split(s, -1) {
		buf.WriteString(strings.ToUpper(p[:1]))
		buf.WriteString(p[1:])
	}
	return buf.String()
}

// SnakeToUpperCamelCase is deprecated
func SnakeToUpperCamelCase(s string) string {
	fmt.Printf("Warning: helpers.SnakeToUpperCamelCase is deprecated\n")
	if s == "" {
		return ""
	}
	buf := bytes.Buffer{}
	for _, p := range rsnake.Split(s, -1) {
		buf.WriteString(strings.ToUpper(p[:1]))
		buf.WriteString(p[1:])
	}
	return buf.String()
}

// SnakeToLowerCamelCase is deprecated
func SnakeToLowerCamelCase(s string) string {
	fmt.Printf("Warning: helpers.SnakeToLowerCamelCase is deprecated\n")
	if s == "" {
		return ""
	}
	buf := bytes.Buffer{}
	for i, p := range rsnake.Split(s, -1) {
		if i == 0 {
			buf.WriteString(p)
			continue
		}
		buf.WriteString(strings.ToUpper(p[:1]))
		buf.WriteString(p[1:])
	}
	return buf.String()
}

// ToUpperFirst is deprecated
func ToUpperFirst(s string) string {
	fmt.Printf("Warning: helpers.ToUpperFirst is deprecated\n")
	buf := bytes.Buffer{}
	buf.WriteString(strings.ToUpper(s[:1]))
	buf.WriteString(s[1:])
	return buf.String()
}

// ToLowerFirst is deprecated
func ToLowerFirst(s string) string {
	fmt.Printf("Warning: helpers.ToLowerFirst is deprecated\n")
	buf := bytes.Buffer{}
	buf.WriteString(strings.ToLower(s[:1]))
	buf.WriteString(s[1:])
	return buf.String()
}
