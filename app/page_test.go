package app_test

import (
	"os"
	"testing"

	"github.com/opsaaaaa/bplate.md/app"
	"github.com/opsaaaaa/bplate.md/spec"
)

func TestReadBoilerplate(t *testing.T) {
  exampleSuccessful := func () {
    template, header, err := app.ReadBoilerplate("../_boilerplates/test.md")

    if err != nil { 
      t.Log(os.Getwd())
      t.Fatal(err)
    }

    spec.AssertContains(t, header, "_boilerplate:", "header must include _boilerplate:")
    spec.AssertContains(t, header, "path:", "header must include path:")
    spec.AssertNotContains(t, header, "layout:", "header cant include layout:")
    spec.AssertNotContains(t, header, "author:", "header cant include author:")
    spec.AssertNotContains(t, header, "Default Test Heading", "header cant include Default Test Heading")

    spec.AssertNotContains(t, template, "_boilerplate:", "template cant include _boilerplate:")
    spec.AssertNotContains(t, template, "path:", "template cant include path:")
    spec.AssertContains(t, template, "layout:", "template must include layout:")
    spec.AssertContains(t, template, "author:", "template must include author:")
    spec.AssertContains(t, template, "Default Test Heading", "template must include Default Test Heading")
  }
  exampleSuccessful()
}


