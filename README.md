# go-jstmpl

Template renderer using JSON Schema as data

## Installation

```
go get -u github.com/go-jstmpl/go-jstmpl/cmd/jstmpl
```

## Usage

```
jstmpl -s schema.json -t schema/template
jstmpl -s schema.yml -t schema/template
```

## Features

- Can generate structs and validators in Go from JSON Schema definitions.
- Can generate router, controllers and mock server in Go from JSON Schema links.
- Can generate models and validators in TypeScript from JSON Schema definitions.
- Can generate model fetcher in TypeScript from JSON Schema links.

And more ...

# References

| Name                                                     | Notes                            |
|:--------------------------------------------------------:|:---------------------------------|
| [go-jsschema](https://github.com/lestrrat/go-jsschema)   | JSON Schema implementation       |
| [go-jshschema](https://github.com/lestrrat/go-jshschema) | JSON Hyper Schema implementation |
| [go-jsref](https://github.com/lestrrat/go-jsref)         | JSON Reference implementation    |
| [go-jspointer](https://github.com/lestrrat/go-jspointer) | JSON Pointer implementations     |
