package app

import (
	"os"
	"strings"
	"testing"

	"github.com/opsaaaaa/bplate.md/spec"
)

func Test_readBoilerplate_happyPath(t *testing.T) {
  template, header, err := readBoilerplate("../.boilerplates/test.md")

  if err != nil { 
    t.Log(os.Getwd())
    t.Fatal(err)
  }

  spec.AssertNotContains(t, header, "_boilerplate:", "header cant include _boilerplate:")
  spec.AssertContains(t, header, "path:", "header must include path:")
  spec.AssertNotContains(t, header, "layout:", "header cant include layout:")
  spec.AssertNotContains(t, header, "author:", "header cant include author:")
  spec.AssertNotContains(t, header, "Default Test Heading", "header cant include Default Test Heading")

  if strings.HasPrefix(header, DOUBLESPACE) {
    t.Error("header must have root indentation")
  }

  spec.AssertNotContains(t, template, "_boilerplate:", "template cant include _boilerplate:")
  spec.AssertNotContains(t, template, "path:", "template cant include path:")
  spec.AssertContains(t, template, "layout:", "template must include layout:")
  spec.AssertContains(t, template, "author:", "template must include author:")
  spec.AssertContains(t, template, "Default Test Heading", "template must include Default Test Heading")
}

func Test_readBoilerplate_fileErrPath(t *testing.T) {
  _, _, err := readBoilerplate("../.boilerplates/nonExistentFile.md")
  if err == nil {
    t.Fatal("It must return an err when reading non existent files")
  }
}


