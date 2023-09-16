package app

import (
  "fmt"
  "log"
  "path/filepath"
)

type PageOptions struct {

}


func MakePage(name string, options PageOptions) {
  path := filepath.Join("_boilerplates", name)
  template, header, err := readBoilerplate(path)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(header)
  fmt.Println(template)

  // data, err := os.ReadFile(path)

  // strings.Cut(::)


}

