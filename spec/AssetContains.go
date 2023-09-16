package spec

import (
  "strings"
  "testing"
)

func AssertContains(t *testing.T, s string, subs string, msg string) {
  if !strings.Contains(s, subs) {
    t.Error(msg)
  }
}

func AssertNotContains(t *testing.T, s string, subs string, msg string) {
  if strings.Contains(s, subs) {
    t.Error(msg)
  }
}
