package jstmpl_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/ghodss/yaml"
	"github.com/go-jstmpl/go-jstmpl"
	jstypes "github.com/go-jstmpl/go-jstmpl/types"
	"github.com/lestrrat/go-jshschema"
)

func ParseHschema(file string) (*hschema.HyperSchema, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("fail to read the source JSON Schema file: %s", err)
	}

	var m map[string]interface{}
	if err := yaml.Unmarshal(b, &m); err != nil {
		return nil, fmt.Errorf("fail to unmarshal YAML: %s", err)
	}

	hs := hschema.New()
	if err := hs.Extract(m); err != nil {
		return nil, fmt.Errorf("fail to extract JSON Schema: %s", err)
	}

	return hs, nil
}

func TestBuilderParseQuery(t *testing.T) {
	hs, err := ParseHschema("./test/query.yml")
	if err != nil {
		t.Fatal(err)
	}
	b := jstmpl.NewBuilder()
	ts, err := b.Build(hs)
	if err != nil {
		t.Fatalf("fail to build: %s", err)
	}

	for _, l := range ts.Links {
		if l.UrlParameters == nil || len(l.UrlParameters) != 2 {
			t.Fatal("fail to parse UrlParameters: %+v", ts)
		}
		for _, p := range l.UrlParameters {
			switch pp := p.(type) {
			case *jstypes.Integer:
				if pp.Title() != "test integer" {
					t.Errorf("fail to resolve UrlParameters: title: Expect: test integer, Actual: %+v", pp.Title())
				}
			case *jstypes.Boolean:
				if pp.Title() != "test bool" {
					t.Errorf("fail to resolve UrlParameters: title: Expect: test bool, Actual: %+v", pp.Title())
				}
			default:
				t.Errorf("fail to resolve UrlParameters: unknown type: %+v", pp)
			}
		}
	}
}

func TestBuilderLoopRef(t *testing.T) {
	hs, err := ParseHschema("./test/ref_loop.yml")
	if err != nil {
		t.Fatal(err)
	}
	b := jstmpl.NewBuilder()
	_, err = b.Build(hs)
	if err == nil {
		t.Fatalf("build should be failed: %+v", b)
	}
}

func TestBuillderNotHaveHref(t *testing.T) {
	hs, err := ParseHschema("./test/has_not_href.yml")
	if err != nil {
		t.Fatal(err)
	}
	b := jstmpl.NewBuilder()
	ts, err := b.Build(hs)
	if err != nil {
		t.Fatalf("fail to build: %s", err)
	}

	if ts.HyperSchema.Title != "has not href" {
		t.Error("fail to get title")
	}
}

func TestBuilderPassBuild(t *testing.T) {
	hs, err := ParseHschema("./test/pass.yml")
	if err != nil {
		t.Fatal(err)
	}
	b := jstmpl.NewBuilder()
	ts, err := b.Build(hs)
	if err != nil {
		t.Fatalf("fail to build: %s", err)
	}

	if len(ts.Objects) != 1 {
		t.Fatal("fail to parse Object type schema")
	}

	if ts.Objects[0] == nil || ts.Objects[0].Title() != "test object" {
		t.Errorf("fail to get Object type schema: %+v", ts.Objects[0])
	}

	if len(ts.Arrays) != 1 {
		t.Fatal("fail to parse Array type schema")
	}

	if ts.Arrays[0] == nil || ts.Arrays[0].Title() != "test array" {
		t.Errorf("fail to get Arrays type schema: %+v", ts.Arrays[0])
	}

	if len(ts.Booleans) != 1 {
		t.Fatal("fail to parse Boolean type schema")
	}

	if ts.Booleans[0] == nil || ts.Booleans[0].Title() != "test boolean" {
		t.Errorf("fail to get Boolean type schema: %+v", ts.Booleans[0])
	}

	if len(ts.Numbers) != 1 {
		t.Fatal("fail to parse Number type schema")
	}

	if ts.Numbers[0] == nil || ts.Numbers[0].Title() != "test number" {
		t.Errorf("fail to get Numbers type schema: %+v", ts.Numbers[0])
	}

	if len(ts.Integers) != 2 {
		t.Fatal("fail to parse Integer type schema")
	}

	if ts.Integers[0] == nil || ts.Integers[0].Title() != "test integer" {
		t.Errorf("fail to get Integer type schema: %+v", ts.Integers[0])
	}

	if len(ts.Strings) != 1 {
		t.Fatal("fail to parse String type schema")
	}

	if ts.Strings[0] == nil || ts.Strings[0].Title() != "test string" {
		t.Errorf("fail to get Strings type schema: %+v", ts.Strings[0])
	}

	if len(ts.Properties) != 2 {
		t.Fatal("fail to parse Properties type schema")
	}

	if len(ts.Definitions) != 7 {
		t.Fatal("fail to parse Definitions")
	}

	for _, v := range ts.Properties {
		switch v.Key() {
		case "test_multitype":
			i, ok := v.(*jstypes.Integer)
			if !ok {
				t.Errorf("fail to get Properties type not link extra schema: Type Convert:%+v", v)
				continue
			}
			if i.ColumnName != "test multitype" || i.ColumnType != "int" {
				t.Errorf("fail to get Properties type not link extra schema: Parse: %+v", i)
			}

		case "test_multitype_link":
			i, ok := v.(*jstypes.Integer)
			if !ok {
				t.Errorf("fail to get Properties type link extra schema: Type Convert:%+v", v)
				continue
			}
			if i.ColumnName != "test multitype" || i.ColumnType != "int" {
				t.Errorf("fail to get Properties type link extra schema: Parse: %+v", i)
			}

		default:
			t.Errorf("fail to get Properties type schema, specify one of key: %s", v)
		}
	}

	if len(ts.Links) != 1 {
		t.Fatal("fail to parse Links type schema")
	}

	for _, v := range ts.Links {
		if v.Title != "Create test" {
			t.Errorf("fail to get Title: expect Create test, but %s", v.Title)
		}
		if v.Method != "POST" {
			t.Errorf("fail to get Method: expect POST, but %s", v.Method)
		}
		switch obj := v.Schema.(type) {
		case *jstypes.Object:
			if len(obj.Required) != 1 || obj.Required[0] != "test" {
				t.Errorf("fail to get Required: %v", obj.Required)
			}
			for _, p := range obj.Properties {
				switch p.Key() {
				case "test_multitype":
					i, ok := p.(*jstypes.Integer)
					if !ok {
						t.Errorf("fail to get Links type not link extra schema: Type Convert:%+v", v)
						continue
					}
					if i.ColumnName != "test multitype" || i.ColumnType != "int" {
						t.Errorf("fail to get Properties type not link extra schema: Parse: %+v", i)
					}

				case "test_multitype_link":
					i, ok := p.(*jstypes.Integer)
					if !ok {
						t.Errorf("fail to get Links type link extra schema: Type Convert:%+v", v)
						continue
					}
					if i.ColumnName != "test multitype" || i.ColumnType != "int" {
						t.Errorf("fail to get Properties type link extra schema: Parse: %+v", i)
					}
				default:
					t.Errorf("failt to get Links type schema, specify one of keys: %+v", p)
				}
			}
		default:
			t.Errorf("fail to get Links type, should be object schema: %+v", v)
		}
	}
}

func TestBuilderConbinatrial(t *testing.T) {
	hs, err := ParseHschema("./test/combine.yml")
	if err != nil {
		t.Fatalf("fail to parse resolve: %v", err)
	}

	b := jstmpl.NewBuilder()
	ts, err := b.Build(hs)
	if err != nil {
		t.Fatalf("fail to build: %s", err)
	}

	p, ok := ts.Objects[0].Properties[0].(*jstypes.Integer)
	if !ok {
		t.Fatalf("fail to type convert: should be integer type: %T", ts.Objects[0].Properties[0])
	}

	if p.Name != "TestParts" || p.ColumnName != "test_parts" {
		t.Errorf("fail to parse conbinatrial definition: %+v", p)
	}

}

func TestResolvePass(t *testing.T) {
	hs, err := ParseHschema("./test/resolve_pass.yml")
	if err != nil {
		t.Fatalf("fail to parse resolve: %v", err)
	}

	if hs.Schema.Definitions["test_link"] == nil {
		t.Fatalf("fail to get schema: %+v", hs)
	}

	sc := hs.Schema.Definitions["test_link"]
	ctx := &jstypes.Context{
		Validations: map[string]bool{},
	}

	for k := range hs.Definitions {
		ctx.Key = k
	}

	scm, err := jstmpl.Resolve(sc, hs.Schema, ctx)
	if err != nil {
		t.Fatalf("fail to resolve: %s", err)
	}

	if scm.Title() != "test" {
		t.Errorf("fail to resolve link schema: title should be test: %+v", scm)
	}
}
