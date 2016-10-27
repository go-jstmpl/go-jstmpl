package helpers

import (
	"errors"
	"fmt"
	"strings"
)

var (
	UnclosedBracketError = errors.New("can not find the pair of bracket")
)

type UncaughtCharacterError struct {
	Path      string
	Position  int
	Character byte
}

func (err UncaughtCharacterError) Error() string {
	return fmt.Sprintf("Uncaught character '%s' at %d in '%s'", err.Character, err.Position, err.Path)
}

type Token interface {
	Gorilla() string
	Sinatra() string
}
type Chunk string

func NewChunk(s string) *Chunk {
	c := Chunk(s)
	return &c
}

func (c *Chunk) Append(b byte) {
	*c += Chunk(b)
}

func (c Chunk) Gorilla() string {
	return string(c)
}

func (c Chunk) Sinatra() string {
	return string(c)
}

type Ref string

func NewRef(s string) *Ref {
	r := Ref(s)
	return &r
}

func (r *Ref) Append(b byte) {
	*r += Ref(b)
}

func (r Ref) Gorilla() string {
	return fmt.Sprintf("{%s}", ParseParam(string(r)))
}

func (r Ref) Sinatra() string {
	return fmt.Sprintf(":%s", ParseParam(string(r)))
}

func BuildURLToken(p string) ([]Token, error) {
	l := len(p)
	tokens := []Token{}
	chunk := NewChunk("")
	var ref *Ref
	for i := 0; i < l; i++ {
		c := p[i]
		switch {
		case c == '{':
			chunk = nil
			ref = NewRef("")
		case ref != nil && c == '}':
			tokens = append(tokens, *ref)
			ref = nil
		case ref == nil && c == '/':
			if chunk != nil {
				tokens = append(tokens, *chunk)
			}
			chunk = NewChunk("")
		case chunk != nil:
			chunk.Append(c)
		case ref != nil:
			ref.Append(c)
		default:
			return tokens, UncaughtCharacterError{
				Path:      p,
				Position:  i,
				Character: c,
			}
		}
	}
	if ref != nil {
		return tokens, UnclosedBracketError
	}
	if chunk != nil {
		tokens = append(tokens, *chunk)
	}
	return tokens, nil
}

func ParseParam(p string) string {
	i := strings.LastIndex(p, "/")
	return p[i+1:]
}

func ToPathLikeGorilla(p string) (string, error) {
	tokens, err := BuildURLToken(p)
	if err != nil {
		return "", err
	}
	s := make([]string, len(tokens))
	for i, t := range tokens {
		s[i] = t.Gorilla()
	}
	return strings.Join(s, "/"), nil
}

func ToPathLikeSinatra(p string) (string, error) {
	tokens, err := BuildURLToken(p)
	if err != nil {
		return "", err
	}
	s := make([]string, len(tokens))
	for i, t := range tokens {
		s[i] = t.Sinatra()
	}
	return strings.Join(s, "/"), nil
}
