package hschema

import (
	"errors"

	"github.com/lestrrat/go-jsschema"
	"github.com/lestrrat/go-pdebug"
)

func extractString(s *string, m map[string]interface{}, n string, required bool) error {
	v, ok := m[n]
	if !ok {
		if required {
			return errors.New("'" + n + "' is required")
		}
		return nil
	}

	switch v.(type) {
	case string:
	default:
		return errors.New("'" + n + "' must be a string")
	}
	*s = v.(string)

	return nil
}

func extractSchema(s **schema.Schema, m map[string]interface{}, name string) error {
	v, ok := m[name]
	if !ok {
		return nil
	}

	if pdebug.Enabled {
		pdebug.Printf("Found property '%s'", name)
	}
	switch v.(type) {
	case map[string]interface{}:
		s1 := schema.New()
		if err := s1.Extract(v.(map[string]interface{})); err != nil {
			return err
		}
		*s = s1
	case *schema.Schema:
		*s = v.(*schema.Schema)
	default:
		return errors.New("key '" + name + "' must be a JSON Schema")
	}

	return nil
}

func (s *HyperSchema) Extract(m map[string]interface{}) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("HyperSchema.Extract").BindError(&err)
		defer g.End()
	}

	s.Schema = schema.New()
	if err := s.Schema.Extract(m); err != nil {
		return err
	}

	if err := extractString(&s.PathStart, m, "pathStart", false); err != nil {
		return err
	}

	if v, ok := m["links"]; ok {
		if err := s.Links.Extract(v); err != nil {
			return err
		}
		s.Links.SetParent(s)
	}

	if v, ok := m["media"]; ok {
		if err := s.Media.Extract(v); err != nil {
			return err
		}
	}

	for _, k := range []string{"pathStart", "links", "media"} {
		delete(s.Schema.Extras, k)
	}

	return nil
}

func (ll *LinkList) Extract(v interface{}) error {
	switch v.(type) {
	case []interface{}:
	default:
		return errors.New("value is not a slice")
	}

	m := v.([]interface{})
	r := make(LinkList, len(m))
	for i, e := range m {
		l1 := Link{}
		if err := l1.Extract(e); err != nil {
			return err
		}
		r[i] = &l1
	}
	*ll = LinkList(r)
	return nil
}

func (l *Link) Extract(v interface{}) (err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("Link.Extract").BindError(&err)
		defer g.End()
	}

	switch v.(type) {
	case map[string]interface{}:
	default:
		return errors.New("value is not a map")
	}

	m := v.(map[string]interface{})

	if err := extractString(&l.Href, m, "href", true); err != nil {
		return err
	}

	if err := extractString(&l.Rel, m, "rel", true); err != nil {
		return err
	}

	if err := extractString(&l.Title, m, "title", false); err != nil {
		return err
	}

	if err := extractString(&l.MediaType, m, "mediaType", false); err != nil {
		return err
	}

	if err := extractString(&l.Method, m, "method", false); err != nil {
		return err
	}

	if err := extractString(&l.EncType, m, "encType", false); err != nil {
		return err
	}

	if err := extractSchema(&l.Schema, m, "schema"); err != nil {
		return err
	}

	if err := extractSchema(&l.TargetSchema, m, "targetSchema"); err != nil {
		return err
	}

	l.Extras = make(map[string]interface{})
	for k, v := range m {
		switch k {
		case "href", "rel", "title", "targetSchema", "mediaType", "method", "encType", "schema":
			continue
		}
		l.Extras[k] = v
	}

	return nil
}

func (m *Media) Extract(v interface{}) error {
	switch v.(type) {
	case map[string]interface{}:
	default:
		return errors.New("value is not a map")
	}

	m1 := v.(map[string]interface{})

	if err := extractString(&m.Type, m1, "type", false); err != nil {
		return err
	}

	if err := extractString(&m.BinaryEncoding, m1, "binaryEncoding", false); err != nil {
		return err
	}

	return nil
}
