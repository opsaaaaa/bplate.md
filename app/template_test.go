package app

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/opsaaaaa/bplate.md/spec"
)

type exampleProps struct {
  Count int
}

func Test_template(t *testing.T) {
  tmpl, err := template.New("example").Funcs(template.FuncMap{
    "funci": func() string {return "sploosh"},
  }).Parse("{{ .Count }} {{ funci }}")

  spec.AssertErrorNil(t,err)

  buf := new(bytes.Buffer)
  err = tmpl.Execute(buf,exampleProps{Count: 30})
  out := buf.String()

  spec.AssertErrorNil(t,err)
  spec.AssertSame(t,out, "30 sploosh")
}


