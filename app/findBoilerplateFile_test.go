package app

import (
	"os"
	"testing"

	"github.com/opsaaaaa/bplate.md/spec"
)


func Test_findBoilerplateFile_directPath(t *testing.T) {
  _ = os.Chdir("../")
  file, err := findBoilerplateFile(".boilerplates/example.md")
  spec.AssertSame(t,file,".boilerplates/example.md")
  spec.AssertErrorNil(t,err)
  _ = os.Chdir("app")
}

func Test_findBoilerplateFile_boilerplatesFolder(t *testing.T) {
  _ = os.Chdir("../")
  file, err := findBoilerplateFile("example.md")
  spec.AssertSame(t,file,".boilerplates/example.md")
  spec.AssertErrorNil(t,err)
  _ = os.Chdir("app")
}

func Test_findBoilerplateFile_boilerplateFile(t *testing.T) {
  _ = os.Chdir("../")
  file, err := findBoilerplateFile(".boilerplates/caseone")
  spec.AssertSame(t,file,".boilerplates/caseone/.boilerplate")
  spec.AssertErrorNil(t,err)
  _ = os.Chdir("app")
}

func Test_findBoilerplateFile_boilerplateMdFile(t *testing.T) {
  _ = os.Chdir("../")
  file, err := findBoilerplateFile(".boilerplates/casetwo")
  spec.AssertSame(t,file,".boilerplates/casetwo/.boilerplate.md")
  spec.AssertErrorNil(t,err)
  _ = os.Chdir("app")
}


