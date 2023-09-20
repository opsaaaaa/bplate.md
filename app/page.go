package app

import (
	"fmt"
	"log"
	"path/filepath"
)

func MakePage(name string, args []string) {
  var err error
  path := filepath.Join("_boilerplates", name)

  body, header, err := readBoilerplate(path)
  if err != nil { log.Fatal(err) }

  opt, err := parseOptions(name,header,args)
  if err != nil { log.Fatal(err) }

  slug, err := templateBoilerplate(opt.Slug, &opt)
  if err != nil { log.Fatalf("Error while parsing slug:\n`%v`\n%v", opt.Slug, err) }
  
  body, err = templateBoilerplate(body, &opt)
  if err != nil { log.Fatalf("Error while parsing boilerplate:\n```\n%v\n```\n%v", body, err) }

  fmt.Println(body)
  fmt.Println(slug)

  // data, err := os.ReadFile(path)
  // m := make(map[string]int)
  // strings.Cut(::)


}

