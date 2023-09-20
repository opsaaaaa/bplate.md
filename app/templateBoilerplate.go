package app

import (
	"bytes"
	"net/url"
	"regexp"
	"strings"
	"text/template"
	"time"
)

type templateProps struct {
  Count int
}


func buildFuncMap(opt PageOptions) template.FuncMap {
  fm := template.FuncMap{
    // These don't need opt
    "asFileSlug": func(s string) string {return asFileSlug(s)},
    "asUrlSlug": func(s string) string {return asUrlSlug(s)},
    "timeFormat": func(f string) string {return time.Now().Format(f)},
    "dateTime": func() string {return time.Now().Format(time.DateTime)},
    "time": func() string {return time.Now().Format(time.TimeOnly)},
    "date": func() string {return time.Now().Format(time.DateOnly)},

    // these require opt
    // could pass it as data instead of functions.
    "ext": func() string {return opt.Ext },
    "boilerplate": func() string {return opt.Boilerplate},
    "path": func() string {return opt.Path},
    // filepath.Ext
    // filepath.Base
    // filepath.Dir
    // < > . , ; : = ? * [ ]
    // /\ : * " ? < > | 
    // / \ ? % * ; | " ' < > . , = SPACE ` `
  }
  for idx, key := range opt.Command {
    fm[key] = func() string {return opt.Args[idx]}
  }
  return fm
}

func templateBoilerplate(txt string, opt PageOptions) (out string, err error) {
  tmpl, err := template.New(opt.Boilerplate).Funcs(buildFuncMap(opt)).Parse(txt)
  if err != nil { return }
  buf := new(bytes.Buffer)
  err = tmpl.Execute(buf,templateProps{Count: 30})
  out = buf.String()
  return
}

const REGEX_BADFILECHAR = "[/\\?%*;|\"'<>,=]"



func asFileSlug(s string) string {
  reg := regexp.MustCompile(REGEX_BADFILECHAR)
  s = reg.ReplaceAllString(s, EMPTYSTRING)
  return strings.ReplaceAll(strings.ToLower(s), SPACE, HYPHEN)
}

func asUrlSlug(s string) string {
  reg := regexp.MustCompile(REGEX_BADFILECHAR)
  s = strings.ReplaceAll(strings.ToLower(s), SPACE, HYPHEN)
  s = reg.ReplaceAllString(s, EMPTYSTRING)
  s = url.QueryEscape(s)
  return s
}

