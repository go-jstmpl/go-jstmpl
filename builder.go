package jstmpl

import (
	"fmt"
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
	r, err := types.NewRoot(hs)
	if err != nil {
		return nil, err
	}
	ctx := &types.Context{
		Validations: map[string]bool{},
	}
	var ds, os, as, ss, ns, is, bs, ps []types.Schema

	for k, s := range hs.Definitions {
		ctx.Key = k
		rs, err := Resolve(s, hs.Schema, ctx)
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
		rs, err := Resolve(s, hs.Schema, ctx)
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

	r.Definitions = ds
	r.Objects = make([]*types.Object, len(os))
	for i, v := range os {
		r.Objects[i] = v.(*types.Object)
	}
	r.Arrays = make([]*types.Array, len(as))
	for i, v := range as {
		r.Arrays[i] = v.(*types.Array)
	}
	r.Strings = make([]*types.String, len(ss))
	for i, v := range ss {
		r.Strings[i] = v.(*types.String)
	}
	r.Numbers = make([]*types.Number, len(ns))
	for i, v := range ns {
		r.Numbers[i] = v.(*types.Number)
	}
	r.Integers = make([]*types.Integer, len(is))
	for i, v := range is {
		r.Integers[i] = v.(*types.Integer)
	}
	r.Booleans = make([]*types.Boolean, len(bs))
	for i, v := range bs {
		r.Booleans[i] = v.(*types.Boolean)
	}

	r.Properties = ps

	for i, l := range hs.Links {
		var (
			s, ts types.Schema
			us    []types.Schema
		)
		m, err := GetUrlParameters(l.Href, hs.Schema, ctx)
		if err != nil {
			return nil, err
		}
		us = append(us, m...)

		if l.Schema != nil {
			ctx.Key = ""
			s, err = Resolve(l.Schema, hs.Schema, ctx)
			if err != nil {
				return nil, err
			}
		}

		if l.TargetSchema != nil {
			ctx.Key = ""
			ts, err = Resolve(l.TargetSchema, hs.Schema, ctx)
			if err != nil {
				return nil, err
			}
		}

		rl, err := types.NewLink(l, s, ts, r, us)
		if err != nil {
			return nil, err
		}
		r.Links[i] = rl

	}

	r.RequiredValidations = ctx.RequiredValidations()
	return r, nil
}

// GetUrlParameters is resolve a JSON Schema from JSON Path and
// catch properties. for example
// if path = '/path/to/resources/{#/definitions/resource}', then
// getting specified schema
func GetUrlParameters(h string, t *schema.Schema, ctx *types.Context) ([]types.Schema, error) {
	var m []types.Schema
	var i, j int
	b := false
	// parse {...} type string
	for j = 0; j < len(h); j++ {
		switch h[j : j+1] {
		case "{":
			i, b = j, true
		case "}":
			if !b {
				return nil, fmt.Errorf("fail to parse url parameter: invalid URL: %+v", h)
			}
			s := schema.New()
			s.Reference = h[i:j]
			scm, err := Resolve(s, t, ctx)
			if err != nil {
				return nil, fmt.Errorf("fail to parse url parameter: resolve: %+v", s)
			}
			m = append(m, scm)
			b = false
		}
	}
	return m, nil
}

func Resolve(s, t *schema.Schema, ctx *types.Context) (types.Schema, error) {
	rs, err := s.Resolve(t)
	ctx.Raw = s
	if err != nil {
		return nil, err
	}
	var ts types.Schema
	if rs.Type.Contains(schema.ObjectType) {
		obj := types.NewObject(ctx, rs)
		for key, sp := range rs.Properties {
			if sp != nil {
				ctx.Key = key
				dp, err := Resolve(sp, t, ctx)
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
			dp, err := Resolve(sp, t, ctx)
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
