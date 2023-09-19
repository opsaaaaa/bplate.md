package app

import (
	"testing"

  "github.com/opsaaaaa/bplate.md/spec"
)

func Test_parseYamlOptions(t *testing.T) {
  happyPath := func () {
    var o PageOptions
    err := parseYamlOptions(&o,"path: _posts\ntimestamp: true\nargs: [one, two, three]")

    spec.AssertErrorNil(t, err)
    spec.AssertSame(t, o.Path, "_posts")
    spec.AssertSame(t, o.Timestamp, true)
    spec.AssertSame(t, len(o.Args), 3)
  }
  happyPath()

  emptyTxt := func () {
    var o PageOptions
    _ = parseYamlOptions(&o, "")
    spec.AssertSame(t, o.Path, "")
    spec.AssertSame(t, o.Timestamp, false)
  }
  emptyTxt()
}


