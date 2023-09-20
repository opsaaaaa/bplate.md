package app

import (
	"fmt"
	"log"
	"path/filepath"
)

func MakePage(name string, args []string) {
  var err error
  path := filepath.Join("_boilerplates", name)
  template, header, err := readBoilerplate(path)

  if err != nil { log.Fatal(err) }

  opt, err := parseOptions(name,header,args)

  if err != nil { log.Fatal(err) }

  fmt.Println(opt)
  fmt.Println(template)

  // data, err := os.ReadFile(path)
  // m := make(map[string]int)
  // strings.Cut(::)


}

