package app

import (
	"fmt"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v3"
)
// m := make(map[string]int)

func MakePage(name string, options PageOptions) {
  var err error
  path := filepath.Join("_boilerplates", name)
  template, header, err := readBoilerplate(path)

  if err != nil { log.Fatal(err) }

  var headerOptions PageOptions
  err = yaml.Unmarshal([]byte(header), &headerOptions)

  if err != nil { log.Fatal(err) }

  fmt.Println(headerOptions)
  fmt.Println(template)

  // data, err := os.ReadFile(path)

  // strings.Cut(::)


}

