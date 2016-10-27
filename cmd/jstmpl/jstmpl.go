package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/go-jstmpl/go-jstmpl"
	"github.com/jessevdk/go-flags"
	"github.com/lestrrat/go-jshschema"
	"github.com/pkg/errors"
)

func main() {
	os.Exit(_main())
}

type options struct {
	DumpFile string `short:"d" long:"dump" description:"intermediate data"`
	OutDir   string `short:"o" long:"output" description:"output directory"`
	Schema   string `short:"s" long:"schema" description:"JSON Schema file"`
	Template string `short:"t" long:"template" description:"template directory"`
}

func _main() int {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		log.Printf("fail to parse flags: %s", err)
		return 1
	}

	f, err := os.Open(opts.Schema)
	if err != nil {
		log.Printf("fail to open the source JSON Schema file: %s", err)
		return 1
	}
	defer f.Close()

	var m map[string]interface{}
	switch ext := filepath.Ext(opts.Schema); ext {
	case ".json":
		if err := json.NewDecoder(f).Decode(&m); err != nil {
			log.Printf("fail to decode JSON: %s", err)
			return 1
		}
	case ".yml", ".yaml":
		b, err := ioutil.ReadFile(opts.Schema)
		if err != nil {
			log.Printf("fail to read the source JSON Schema file: %s", err)
			return 1
		}
		if err := yaml.Unmarshal(b, &m); err != nil {
			log.Printf("fail to unmarshal YAML: %s", err)
			return 1
		}
	default:
		log.Printf("undefined extension: %s", ext)
		return 1
	}

	hs := hschema.New()
	if err := hs.Extract(m); err != nil {
		log.Printf("fail to extract JSON Schema: %s", err)
		return 1
	}

	b := jstmpl.NewBuilder()
	ts, err := b.Build(hs)
	if err != nil {
		log.Printf("fail to build: %s", err)
		return 1
	}

	if d := opts.DumpFile; d != "" {
		b, err := json.MarshalIndent(ts, "", "  ")
		if err != nil {
			log.Printf("fail to dump: %s", err)
		}

		switch d {
		case "stdout":
			fmt.Printf("%s\n", b)
		default:
			if err := ioutil.WriteFile(d, b, 0775); err != nil {
				log.Printf("fail to write dump data at %s: %s", d, err)
			}
		}
	}

	err = filepath.Walk(opts.Template, func(i string, info os.FileInfo, err error) error {
		if err := (func() error {
			if info == nil {
				return fmt.Errorf("fail to find a template file or dir: %s", i)
			}
			if info.IsDir() {
				return nil
			}
			r, err := filepath.Rel(opts.Template, i)
			if err != nil {
				return err
			}
			o := filepath.Join(opts.OutDir, r)
			ext := filepath.Ext(o)
			if ext == ".tmpl" {
				o = strings.TrimRight(o, ext)
			}

			var tmpl []byte
			if i != "" {
				f, err := os.Open(i)
				if err != nil {
					return err
				}
				defer f.Close()
				tmpl, err = ioutil.ReadAll(f)
				if err != nil {
					return err
				}
			}

			g := jstmpl.NewGenerator()
			b, gpErr := g.Process(ts, tmpl, filepath.Ext(o))
			if gpErr != nil {
				if _, ok := gpErr.(jstmpl.FormatError); !ok {
					return gpErr
				}
			}

			if o != "" {
				d := filepath.Dir(o)
				_, err := os.Stat(d)
				if err != nil {
					err := os.MkdirAll(d, 0755)
					if err != nil {
						return err
					}
				}
				if err := ioutil.WriteFile(o, b, 0644); err != nil {
					return err
				}
			}

			if gpErr != nil {
				return gpErr
			}

			return nil
		})(); err != nil {
			return errors.Wrapf(err, "in '%s'", i)
		}
		return nil
	})
	if err != nil {
		log.Println(errors.Wrap(err, "fail to generate"))
		return 1
	}

	return 0
}
