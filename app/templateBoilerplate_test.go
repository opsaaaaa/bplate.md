package app

import (
	"testing"
	"time"

	"github.com/opsaaaaa/bplate.md/spec"
)


func Test_templateBoilerplate_withSlug(t *testing.T) {
  txt, err := templateBoilerplate(
    "{{ path }}/{{ date }}-{{ title | asFileSlug }}{{ ext }}", 
    &PageOptions{
      Path: "_posts",
      Boilerplate: "post.md",
      Ext: ".md",
      Args: []string{"Foo Bar"},
      Command: []string{"title"},
    },
  )
  spec.AssertErrorNil(t, err)
  spec.AssertSame(t, txt, "_posts/"+time.Now().Format(time.DateOnly)+"-foo-bar.md")
}

func Test_templateBoilerplate_withMultipleArgs(t *testing.T) {
  txt, err := templateBoilerplate(
    "{{ one }}, {{ two }}, {{ three }}, {{ arg 0 }}, {{ arg 1 }}, {{ arg 2}}, {{ arg 3 }}", 
    &PageOptions{
      Args: []string{"1","2","3"},
      Command: []string{"one", "two", "three"},
    },
  )
  spec.AssertErrorNil(t, err)
  spec.AssertSame(t, txt, "1, 2, 3, 1, 2, 3, ")
}

func Test_asFileSlug(t *testing.T) {
  s := asFileSlug("*Foo;, Bar?.99_(2)-10%")
  spec.AssertSame(t, s, "foo-bar.99_(2)-10")
}

func Test_asUrlSlug(t *testing.T) {
  s := asFileSlug("*Foo;, Bar?.99_(2)-10%")
  spec.AssertSame(t, s, "foo-bar.99_(2)-10")
}

