package jstmpl

import (
	"encoding/json"
	"sort"

	"github.com/lestrrat/go-jsschema"
)

type SchemasByKey []Schema

func (a SchemasByKey) Len() int           { return len(a) }
func (a SchemasByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SchemasByKey) Less(i, j int) bool { return a[i].Key() < a[j].Key() }

type ByTitle []Schema

func (a ByTitle) Len() int           { return len(a) }
func (a ByTitle) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTitle) Less(i, j int) bool { return a[i].Title() < a[j].Title() }

type ByKey []Header

func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

func NewObject(ctx *Context, s *schema.Schema) *Object {
	return &Object{
		Schema:     s,
		NativeType: "object",
		Type:       spaceToUpperCamelCase(s.Title),
		Name:       spaceToUpperCamelCase(s.Title),
		key:        ctx.Key,
		IsPrivate:  false,
		Properties: []Schema{},
	}
}

func (o Object) Raw() *schema.Schema {
	return o.Schema
}

func (o Object) Title() string {
	return o.Schema.Title
}

func (o Object) Format() string {
	return string(o.Schema.Format)
}

func (o Object) Key() string {
	return o.key
}

func (o Object) Example() interface{} {
	e := make(map[string]interface{})
	for _, s := range o.Properties {
		e[s.Key()] = s.Example()
	}
	return e
}

func NewArray(ctx *Context, s *schema.Schema) *Array {
	return &Array{
		Schema:     s,
		NativeType: "array",
		Type:       spaceToUpperCamelCase(s.Title),
		Name:       spaceToUpperCamelCase(s.Title),
		key:        ctx.Key,
		IsPrivate:  false,
		Items: &ItemSpec{
			ItemSpec: s.Items,
			Schemas:  make([]Schema, len(s.Items.Schemas)),
		},
	}
}

func (o Array) Raw() *schema.Schema {
	return o.Schema
}

func (o Array) Title() string {
	return o.Schema.Title
}

func (o Array) Format() string {
	return string(o.Schema.Format)
}

func (o Array) Key() string {
	return o.key
}

func (o Array) Example() interface{} {
	e := make([]interface{}, len(o.Items.Schemas))
	for i, s := range o.Items.Schemas {
		e[i] = s.Example()
	}
	return e
}

func NewString(ctx *Context, s *schema.Schema) *String {
	vs := []Validation{}
	if v, err := NewFormatValidation(s); err == nil {
		ctx.AddValidation(v)
		vs = append(vs, v)
	}
	if v, err := NewMinLengthValidation(s); err == nil {
		ctx.AddValidation(v)
		vs = append(vs, v)
	}
	if v, err := NewMaxLengthValidation(s); err == nil {
		ctx.AddValidation(v)
		vs = append(vs, v)
	}
	if v, err := NewPatternValidation(s); err == nil {
		ctx.AddValidation(v)
		vs = append(vs, v)
	}
	return &String{
		Schema:      s,
		NativeType:  "string",
		Type:        "string",
		Name:        spaceToUpperCamelCase(s.Title),
		key:         ctx.Key,
		IsPrivate:   true,
		Validations: vs,
	}
}

func (o String) Raw() *schema.Schema {
	return o.Schema
}

func (o String) Title() string {
	return o.Schema.Title
}

func (o String) Format() string {
	return string(o.Schema.Format)
}

func (o String) Key() string {
	return o.key
}

func (o String) Example() interface{} {
	e := o.Schema.Extras["example"]
	if e != nil {
		return e
	}
	return ""
}

func NewNumber(ctx *Context, s *schema.Schema) *Number {
	vs := []Validation{}
	if v, err := NewMaximumValidation(s); err == nil {
		ctx.AddValidation(v)
		vs = append(vs, v)
	}
	if v, err := NewMinimumValidation(s); err == nil {
		ctx.AddValidation(v)
		vs = append(vs, v)
	}
	return &Number{
		Schema:      s,
		NativeType:  "number",
		Type:        "number",
		Name:        spaceToUpperCamelCase(s.Title),
		key:         ctx.Key,
		IsPrivate:   true,
		Validations: vs,
	}
}

func (o Number) Raw() *schema.Schema {
	return o.Schema
}

func (o Number) Title() string {
	return o.Schema.Title
}

func (o Number) Format() string {
	return string(o.Schema.Format)
}

func (o Number) Key() string {
	return o.key
}

func (o Number) Example() interface{} {
	e := o.Schema.Extras["example"]
	if e != nil {
		return e
	}
	return 0
}

func NewInteger(ctx *Context, s *schema.Schema) *Integer {
	vs := []Validation{}
	if v, err := NewMaximumValidation(s); err == nil {
		ctx.AddValidation(v)
		vs = append(vs, v)
	}
	if v, err := NewMinimumValidation(s); err == nil {
		ctx.AddValidation(v)
		vs = append(vs, v)
	}
	return &Integer{
		Schema:      s,
		NativeType:  "number",
		Type:        "number",
		Name:        spaceToUpperCamelCase(s.Title),
		key:         ctx.Key,
		IsPrivate:   true,
		Validations: vs,
	}
}

func (o Integer) Raw() *schema.Schema {
	return o.Schema
}

func (o Integer) Title() string {
	return o.Schema.Title
}

func (o Integer) Format() string {
	return string(o.Schema.Format)
}

func (o Integer) Key() string {
	return o.key
}

func (o Integer) Example() interface{} {
	e := o.Schema.Extras["example"]
	if e != nil {
		return e
	}
	return 0
}

func NewBoolean(ctx *Context, s *schema.Schema) *Boolean {
	vs := []Validation{}
	return &Boolean{
		Schema:      s,
		NativeType:  "boolean",
		Type:        "boolean",
		Name:        spaceToUpperCamelCase(s.Title),
		key:         ctx.Key,
		IsPrivate:   true,
		Validations: vs,
	}
}

func (o Boolean) Raw() *schema.Schema {
	return o.Schema
}

func (o Boolean) Title() string {
	return o.Schema.Title
}

func (o Boolean) Format() string {
	return string(o.Schema.Format)
}

func (o Boolean) Key() string {
	return o.key
}

func (o Boolean) Example() interface{} {
	e := o.Schema.Extras["example"]
	if e != nil {
		return e
	}
	return false
}

func (l Link) ReqHeaders() []Header {
	h := []Header{
		Header{
			Key:   "Host",
			Value: l.URL.Host,
		},
		Header{
			Key:   "Content-Type",
			Value: "application/json",
		},
	}
	sort.Sort(ByKey(h))
	return h
}

func (l Link) ReqBody() string {
	if l.Schema == nil {
		return ""
	}

	e := l.Schema.Example()
	if e == nil {
		return ""
	}
	j, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return ""
	}
	return string(j)
}

var (
	statusCodes = map[string]int{
		"GET":    200,
		"POST":   201,
		"PUT":    204,
		"DELETE": 204,
	}
	reasonPhrases = map[int]string{
		100: "Continue",
		101: "Switching Protocols",
		200: "OK",
		201: "Created",
		202: "Accepted",
		203: "Non-Authoritative Information",
		204: "No Content",
		205: "Reset Content",
		206: "Partial Content",
		300: "Multiple Choices",
		301: "Moved Permanently",
		302: "Found",
		303: "See Other",
		304: "Not Modified",
		305: "Use Proxy",
		307: "Temporary Redirect",
		400: "Bad Request",
		401: "Unauthorized",
		402: "Payment Required",
		403: "Forbidden",
		404: "Not Found",
		405: "Method Not Allowed",
		406: "Not Acceptable",
		407: "Proxy Authentication Required",
		408: "Request Time-out",
		409: "Conflict",
		410: "Gone",
		411: "Length Required",
		412: "Precondition Failed",
		413: "Request Entity Too Large",
		414: "Request-URI Too Large",
		415: "Unsupported Media Type",
		416: "Requested range not satisfiable",
		417: "Expectation Failed",
		500: "Internal Server Error",
		501: "Not Implemented",
		502: "Bad Gateway",
		503: "Service Unavailable",
		504: "Gateway Time-out",
		505: "HTTP Version not supported",
	}
)

func (l Link) ResStatusCode() int {
	return statusCodes[l.Method]
}

func (l Link) ResReasonPhrase() string {
	return reasonPhrases[l.ResStatusCode()]
}

func (l Link) ResHeaders() []Header {
	h := []Header{
		Header{
			Key:   "Content-Type",
			Value: "application/json",
		},
	}
	sort.Sort(ByKey(h))
	return h
}

func (l Link) ResBody() string {
	if l.TargetSchema == nil {
		return ""
	}

	e := l.TargetSchema.Example()
	if e == nil {
		return ""
	}
	j, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return ""
	}
	return string(j)
}
