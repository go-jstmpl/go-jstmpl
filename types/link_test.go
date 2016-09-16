package types_test

import (
	"net/url"
	"testing"

	"github.com/go-jstmpl/go-jstmpl/types"
	hschema "github.com/lestrrat/go-jshschema"
)

func TestNewLink(t *testing.T) {
	u, err := url.Parse("http://example.com")
	if err != nil {
		t.Fatalf("fail to parse URL %s", err)
	}
	l, err := types.NewLink(
		&hschema.Link{
			Title: "title",
			Extras: map[string]interface{}{
				"description": "description",
			},
		},
		types.Object{},
		types.Object{},
		&types.Root{URL: u},
	)
	if err != nil {
		t.Fatalf("fail to NewLink %s", err)
	}
	if l.Title != "title" {
		t.Errorf("Title expected %s but actual %s", "title", l.Title)
	}
	if l.Description != "description" {
		t.Errorf("Description expected %s but actual %s", "description", l.Description)
	}
}

func TestNewLinkWithoutDescription(t *testing.T) {
	u, err := url.Parse("http://example.com")
	if err != nil {
		t.Fatalf("fail to parse URL %s", err)
	}
	l, err := types.NewLink(
		&hschema.Link{
			Title:  "title",
			Extras: map[string]interface{}{},
		},
		types.Object{},
		types.Object{},
		&types.Root{URL: u},
	)
	if err != nil {
		t.Fatalf("fail to NewLink %s", err)
	}
	if l.Title != "title" {
		t.Errorf("Title expected %s but actual %s", "title", l.Title)
	}
	if l.Description != "" {
		t.Errorf("Description expected %s but actual %s", "", l.Description)
	}
}
