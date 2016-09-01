package jstmpl

import (
	"bytes"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/lestrrat/go-jshschema"
	"github.com/lestrrat/go-jsschema"
)

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Build(hs *hschema.HyperSchema) (*Root, error) {
	var err error

	m := &Root{
		HyperSchema: hs,
		Links:       make(LinkList, len(hs.Links)),
	}
	str := getString(hs.Schema.Extras, "href")
	m.URL, err = url.Parse(str)
	if err != nil {
		return nil, err
	}

	ctx := &Context{
		validations: map[string]bool{},
	}

	var ds, os, as, ss, ns, is, bs, ps []Schema

	for k, s := range hs.Definitions {
		ctx.Key = k
		rs, err := resolve(s, hs.Schema, ctx)
		if err != nil {
			return nil, err
		}
		ds = append(ds, rs)
		func(s interface{}) {
			switch ts := s.(type) {
			case *Object:
				os = append(os, ts)
			case *Array:
				as = append(as, ts)
			case *String:
				ss = append(ss, ts)
			case *Number:
				ns = append(ns, ts)
			case *Integer:
				is = append(is, ts)
			case *Boolean:
				bs = append(bs, ts)
			}
		}(rs)
	}

	for k, s := range hs.Properties {
		ctx.Key = k
		rs, err := resolve(s, hs.Schema, ctx)
		if err != nil {
			return nil, err
		}
		ps = append(ps, rs)
	}

	sort.Sort(SchemasByKey(ds))
	sort.Sort(SchemasByKey(os))
	sort.Sort(SchemasByKey(as))
	sort.Sort(SchemasByKey(ss))
	sort.Sort(SchemasByKey(ns))
	sort.Sort(SchemasByKey(is))
	sort.Sort(SchemasByKey(bs))
	sort.Sort(SchemasByKey(ps))

	m.Definitions = ds
	m.Objects = make([]*Object, len(os))
	for i, v := range os {
		m.Objects[i] = v.(*Object)
	}
	m.Arrays = make([]*Array, len(as))
	for i, v := range as {
		m.Arrays[i] = v.(*Array)
	}
	m.Strings = make([]*String, len(ss))
	for i, v := range ss {
		m.Strings[i] = v.(*String)
	}
	m.Numbers = make([]*Number, len(ns))
	for i, v := range ns {
		m.Numbers[i] = v.(*Number)
	}
	m.Integers = make([]*Integer, len(is))
	for i, v := range is {
		m.Integers[i] = v.(*Integer)
	}
	m.Booleans = make([]*Boolean, len(bs))
	for i, v := range bs {
		m.Booleans[i] = v.(*Boolean)
	}
	m.Properties = ps

	for i, l := range hs.Links {
		var (
			s, ts Schema
		)

		if l.Schema != nil {
			ctx.Key = ""
			s, err = resolve(l.Schema, hs.Schema, ctx)
			if err != nil {
				return nil, err
			}
		}

		if l.TargetSchema != nil {
			ctx.Key = ""
			ts, err = resolve(l.TargetSchema, hs.Schema, ctx)
			if err != nil {
				return nil, err
			}
		}

		u, err := url.Parse(fmt.Sprintf("%s%s", m.URL.String(), l.Href))
		if err != nil {
			return nil, err
		}
		m.Links[i] = &Link{
			Link:         *l,
			URL:          u,
			Schema:       s,
			TargetSchema: ts,
		}
	}

	m.RequiredValidations = ctx.RequiredValidations()

	return m, nil
}

type Context struct {
	Key         string
	validations map[string]bool
}

func (ctx *Context) AddValidation(v Validation) {
	ctx.validations[v.Func()] = true
}

func (ctx *Context) RequiredValidations() []string {
	vs := []string{}
	for v, b := range ctx.validations {
		if !b {
			continue
		}
		vs = append(vs, v)
	}
	sort.Strings(vs)
	return vs
}

func resolve(s, t *schema.Schema, ctx *Context) (Schema, error) {
	rs, err := s.Resolve(t)
	if err != nil {
		return nil, err
	}

	var ts Schema
	if rs.Type.Contains(schema.ObjectType) {
		obj := NewObject(ctx, rs)
		for key, sp := range rs.Properties {
			if sp != nil {
				ctx.Key = key
				dp, err := resolve(sp, t, ctx)
				if err != nil {
					return nil, err
				}
				ps := append(obj.Properties, dp)
				sort.Sort(SchemasByKey(ps))
				obj.Properties = ps
			}
		}
		ts = obj
	} else if rs.Type.Contains(schema.ArrayType) {
		arr := NewArray(ctx, rs)
		for i, sp := range rs.Items.Schemas {
			ctx.Key = ""
			dp, err := resolve(sp, t, ctx)
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
		ts = NewString(ctx, rs)
	} else if rs.Type.Contains(schema.NumberType) {
		ts = NewNumber(ctx, rs)
	} else if rs.Type.Contains(schema.IntegerType) {
		ts = NewInteger(ctx, rs)
	} else if rs.Type.Contains(schema.BooleanType) {
		ts = NewBoolean(ctx, rs)
	} else {
		return nil, fmt.Errorf("undefined type: %+v", rs)
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
