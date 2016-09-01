package hschema_test

import (
	"strings"
	"testing"

	"github.com/lestrrat/go-jshschema"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	const src = `{
  "definitions": {
    "uint": {
      "type": "integer",
      "minimum": 0
    }
  },
  "links": [
    {
      "href": "/schema",
      "method": "GET",
      "rel": "self",
      "targetSchema": {
        "$ref": "http://json-schema.org/draft-04/hyper-schema#"
      }
    },
    {
      "href": "/user/create",
      "method": "POST",
      "rel": "create user",
      "schema": {
        "type": "object",
				"extra": true,
        "properties": {
          "name": { "type": "string" },
          "age": { "$ref": "#/definitions/uint" },
          "address": { "type": "string" }
        }
      },
      "targetSchema": {
        "type": "array",
        "items": { "$ref": "#/definitions/uint" }
      }
    }
  ]
}`

	s, err := hschema.Read(strings.NewReader(src))
	if !assert.NoError(t, err, "hschema.Read should succeed") {
		return
	}

	if !assert.Len(t, s.Links, 2, "There should be 2 links") {
		return
	}

	l := s.Links[1]
	if !assert.Equal(t, l.Href, "/user/create", "l.Href matches") {
		return
	}

	if !assert.Equal(t, l.Method, "POST", "l.Method matches") {
		return
	}

	_, ok := l.Schema.Extras["extra"]
	if !assert.True(t, ok, "Extra item 'extra' should exist") {
		return
	}
}