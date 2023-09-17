package app

import (
	"testing"

  "github.com/opsaaaaa/bplate.md/spec"
)

func Test_parseYamlOptions(t *testing.T) {
  o, err := parseYamlOptions("path: _posts\ntimestamp: true")

  spec.AssertErrorNil(t, err)
  spec.AssertSame(t, o.Path, "_posts")
  spec.AssertSame(t, o.Timestamp, true)

  o, _ = parseYamlOptions("")
  spec.AssertSame(t, o.Path, "")
  spec.AssertSame(t, o.Timestamp, false)
}


