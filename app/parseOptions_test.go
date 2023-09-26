package app

import (
	"testing"

  "github.com/opsaaaaa/bplate.md/spec"
)

func Test_parseOptions_happyPath(t *testing.T) {
  o, err := parseOptions("example.md","path: _posts\nslug: \"{{ date }}.md\"\ncommand: [one, two, three]", []string{"1","2","3"})

  spec.AssertErrorNil(t, err)
  spec.AssertSame(t, o.Path, "_posts")
  spec.AssertSame(t, len(o.Command), 3)
  spec.AssertSame(t, len(o.Args), 3)
}

func Test_parseOptions_emptyTxt(t *testing.T) {
  o, _ := parseOptions("blank","", []string{})
  spec.AssertSame(t, o.Path, "./")
}


