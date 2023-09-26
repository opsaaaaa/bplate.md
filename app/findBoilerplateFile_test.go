package app

import (
	"os"
	"testing"

	"github.com/opsaaaaa/bplate.md/spec"
)

func Test_findBoilerplateFile(t *testing.T) {
  _ = os.Chdir("../test_cases")

  file, err := findBoilerplateFile("zero/example_boilerplate.md")
  spec.AssertSame(t,file,"zero/example_boilerplate.md")
  spec.AssertErrorNil(t,err)

  file, err = findBoilerplateFile("notes.md")
  spec.AssertSame(t,file,".boilerplates/notes.md")
  spec.AssertErrorNil(t,err)

  file, err = findBoilerplateFile("one")
  spec.AssertSame(t,file,"one/.boilerplate")
  spec.AssertErrorNil(t,err)

  file, err = findBoilerplateFile("two/")
  spec.AssertSame(t,file,"two/.boilerplate.md")
  spec.AssertErrorNil(t,err)

  file, err = findBoilerplateFile("three/post")
  spec.AssertSame(t,file,"three/post.boilerplate")
  spec.AssertErrorNil(t,err)

  file, err = findBoilerplateFile("four/post")
  spec.AssertSame(t,file,"four/post.boilerplate.txt")
  spec.AssertErrorNil(t,err)

  _, err = findBoilerplateFile("a/missing/file/path")
  if err == nil { t.Error("It should error on missing file paths.") }

  _ = os.Chdir("../app")
}


