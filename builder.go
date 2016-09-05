package jstmpl

import (
	"fmt"
	"net/url"
	"sort"

	"github.com/go-jstmpl/go-jstmpl/types"
	"github.com/lestrrat/go-jshschema"
	"github.com/lestrrat/go-jsschema"
)

type Builder struct {
	Schema *schema.Schema
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Build(hs *hschema.HyperSchema) (*types.Root, error) {
	var err error

	m := &types.Root{
		HyperSchema: hs,
		Links:       make(types.LinkList, len(hs.Links)),
	}
	str := hs.Schema.Extras["href"].(string)
	m.URL, err = url.Parse(str)
	if err != nil {
		return nil, err
	}

	ctx := &types.Context{
		Validations: map[string]bool{},
	}

	var ds, os, as, ss, ns, is, bs, ps []types.Schema

	for k, s := range hs.Definitions {
		ctx.Key = k
		rs, err := resolve(s, hs.Schema, ctx)
		if err != nil {
			return nil, err
		}
		ds = append(ds, rs)
		func(s interface{}) {
			switch ts := s.(type) {
			case *types.Object:
				os = append(os, ts)
			case *types.Array:
				as = append(as, ts)
			case *types.String:
				ss = append(ss, ts)
			case *types.Number:
				ns = append(ns, ts)
			case *types.Integer:
				is = append(is, ts)
			case *types.Boolean:
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

	sort.Sort(types.SchemasByKey(ds))
	sort.Sort(types.SchemasByKey(os))
	sort.Sort(types.SchemasByKey(as))
	sort.Sort(types.SchemasByKey(ss))
	sort.Sort(types.SchemasByKey(ns))
	sort.Sort(types.SchemasByKey(is))
	sort.Sort(types.SchemasByKey(bs))
	sort.Sort(types.SchemasByKey(ps))

	m.Definitions = ds
	m.Objects = make([]*types.Object, len(os))
	for i, v := range os {
		m.Objects[i] = v.(*types.Object)
	}
	m.Arrays = make([]*types.Array, len(as))
	for i, v := range as {
		m.Arrays[i] = v.(*types.Array)
	}
	m.Strings = make([]*types.String, len(ss))
	for i, v := range ss {
		m.Strings[i] = v.(*types.String)
	}
	m.Numbers = make([]*types.Number, len(ns))
	for i, v := range ns {
		m.Numbers[i] = v.(*types.Number)
	}
	m.Integers = make([]*types.Integer, len(is))
	for i, v := range is {
		m.Integers[i] = v.(*types.Integer)
	}
	m.Booleans = make([]*types.Boolean, len(bs))
	for i, v := range bs {
		m.Booleans[i] = v.(*types.Boolean)
	}
	m.Properties = ps

	for i, l := range hs.Links {
		var (
			s, ts types.Schema
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
		m.Links[i] = &types.Link{
			Link:         *l,
			URL:          u,
			Schema:       s,
			TargetSchema: ts,
		}
	}

	m.RequiredValidations = ctx.RequiredValidations()

	return m, nil
}

func resolve(s, t *schema.Schema, ctx *types.Context) (types.Schema, error) {
	rs, err := s.Resolve(t)
	if err != nil {
		return nil, err
	}

	var ts types.Schema
	if rs.Type.Contains(schema.ObjectType) {
		obj := types.NewObject(ctx, rs)
		for key, sp := range rs.Properties {
			if sp != nil {
				ctx.Key = key
				dp, err := resolve(sp, t, ctx)
				if err != nil {
					return nil, err
				}
				ps := append(obj.Properties, dp)
				sort.Sort(types.SchemasByKey(ps))
				obj.Properties = ps
			}
		}
		ts = obj
	} else if rs.Type.Contains(schema.ArrayType) {
		arr := types.NewArray(ctx, rs)
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
		ts = types.NewString(ctx, rs)
	} else if rs.Type.Contains(schema.NumberType) {
		ts = types.NewNumber(ctx, rs)
	} else if rs.Type.Contains(schema.IntegerType) {
		ts = types.NewInteger(ctx, rs)
	} else if rs.Type.Contains(schema.BooleanType) {
		ts = types.NewBoolean(ctx, rs)
	} else {
		ts = types.NewUndefined(ctx, rs)
	}

	return ts, nil
}
