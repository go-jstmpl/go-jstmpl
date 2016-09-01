package hschema

import (
	"encoding/json"
	"io"
	"os"

	"github.com/lestrrat/go-jsschema"
)

func New() *HyperSchema {
	return &HyperSchema{
		Schema: schema.New(),
	}
}

func ReadFile(f string) (*HyperSchema, error) {
	in, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer in.Close()
	return Read(in)
}

func Read(in io.Reader) (*HyperSchema, error) {
	s := New()
	if err := s.Decode(in); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *HyperSchema) Decode(in io.Reader) error {
	var m map[string]interface{}
	if err := json.NewDecoder(in).Decode(&m); err != nil {
		return err
	}
	return s.Extract(m)
}

