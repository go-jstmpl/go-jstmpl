package types

import (
	"encoding/json"
	"net/url"
	"sort"

	hschema "github.com/lestrrat/go-jshschema"
)

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

type LinkList []*Link

type Link struct {
	hschema.Link
	URL          *url.URL
	Schema       Schema
	TargetSchema Schema
}

type Header struct {
	Key, Value string
}

type ByKey []Header

func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

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
