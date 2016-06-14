package jstmpl

import (
	"bytes"
	"net/url"
	"regexp"
	"sort"
	"strings"

	"github.com/lestrrat/go-jshschema"
	"github.com/lestrrat/go-jsschema"
)

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Build(hs *hschema.HyperSchema) (*TSModel, error) {
	var err error

	m := &TSModel{
		Links: make(LinkList, len(hs.Links)),
	}
	str := getString(hs.Schema.Extras, "href")
	m.URL, err = url.Parse(str)
	if err != nil {
		return nil, err
	}

	var os, as, ss, ns, is, bs []Schema

	for _, prop := range hs.Definitions {
		s, err := resolve(prop, hs.Schema, "")
		if err != nil {
			return nil, err
		}
		switch ts := s.(type) {
		case Object:
			os = append(os, ts)
		case Array:
			as = append(as, ts)
		case String:
			ss = append(ss, ts)
		case Number:
			ns = append(ns, ts)
		case Integer:
			is = append(is, ts)
		case Boolean:
			bs = append(bs, ts)
		}
	}

	sort.Sort(ByClassName(os))
	sort.Sort(ByClassName(as))
	sort.Sort(ByClassName(ss))
	sort.Sort(ByClassName(ns))
	sort.Sort(ByClassName(is))
	sort.Sort(ByClassName(bs))

	m.Objects = make([]Object, len(os))
	for i, v := range os {
		m.Objects[i] = v.(Object)
	}
	m.Arrays = make([]Array, len(as))
	for i, v := range as {
		m.Arrays[i] = v.(Array)
	}
	m.Strings = make([]String, len(ss))
	for i, v := range ss {
		m.Strings[i] = v.(String)
	}
	m.Numbers = make([]Number, len(ns))
	for i, v := range ns {
		m.Numbers[i] = v.(Number)
	}
	m.Integers = make([]Integer, len(is))
	for i, v := range is {
		m.Integers[i] = v.(Integer)
	}
	m.Booleans = make([]Boolean, len(bs))
	for i, v := range bs {
		m.Booleans[i] = v.(Boolean)
	}

	// for i, l := range hs.Links {
	// 	var (
	// 		s, ts *Schema
	// 	)
	//
	// 	if l.Schema != nil {
	// 		s, err = resolve(l.Schema, hs.Schema)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 	}
	//
	// 	if l.TargetSchema != nil {
	// 		ts, err = resolve(l.TargetSchema, hs.Schema)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 	}
	//
	// 	m.Links[i] = &Link{
	// 		Link:         *l,
	// 		Schema:       s,
	// 		TargetSchema: ts,
	// 	}
	// }

	return m, nil
}

func resolve(src, ctx *schema.Schema, propName string) (Schema, error) {
	rs, err := src.Resolve(ctx)
	if err != nil {
		return nil, err
	}

	var ts Schema
	if rs.Type.Contains(schema.ObjectType) {
		obj := NewObject(propName, rs)
		for key, sp := range rs.Properties {
			if sp != nil {
				dp, err := resolve(sp, ctx, key)
				if err != nil {
					return nil, err
				}
				obj.Properties = append(obj.Properties, dp)
			}
		}
		ts = obj
	} else if rs.Type.Contains(schema.ArrayType) {
		arr := NewArray(propName, rs)
		for i, sp := range rs.Items.Schemas {
			dp, err := resolve(sp, ctx, "")
			if err != nil {
				return nil, err
			}
			arr.Items.Schemas[i] = dp
			if i == 0 {
				arr.Item = dp
			}
		}
		ts = arr
	} else if rs.Type.Contains(schema.StringType) {
		ts = NewString(propName, rs)
	} else if rs.Type.Contains(schema.NumberType) {
		ts = NewNumber(propName, rs)
	} else if rs.Type.Contains(schema.IntegerType) {
		ts = NewInteger(propName, rs)
	} else if rs.Type.Contains(schema.BooleanType) {
		ts = NewBoolean(propName, rs)
	} else {
	}

	// if rs.Items != nil {
	// 	dest.Items = &ItemSpec{
	// 		ItemSpec: rs.Items,
	// 		Schemas:  make([]*Schema, len(rs.Items.Schemas)),
	// 	}
	// 	for i, sp := range rs.Items.Schemas {
	// 		dp, err := resolve(sp, ctx)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		dest.Items.Schemas[i] = dp
	// 	}
	// }

	return ts, nil
}

var rspace = regexp.MustCompile(`\s+`)
var rsnake = regexp.MustCompile(`_`)

func TitleToClassName(s string) string {
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

func KeyToPropName(s string) string {
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

func getString(extras map[string]interface{}, key string) string {
	v, _ := extras[key].(string)
	return v
}
