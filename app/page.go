package app

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

func MakePage(name string, args []string) {
  var err error
  path := filepath.Join("_boilerplates", name)

  body, header, err := readBoilerplate(path)
  if err != nil { log.Fatal(err) }

  opt, err := parseOptions(name,header,args)
  if err != nil { log.Fatal(err) }

  // TODO: create a func for templating a string
  slug, err := templateBoilerplate(opt.Slug, &opt)
  if err != nil { log.Fatalf("Can't parse slug:\n`%v`\n%v", opt.Slug, err) }

  _, err = os.Stat(slug)
  if !errors.Is(err, os.ErrNotExist) {
    log.Fatalf("File `%v` already exists.", slug)
  }

  // 1 execute, write 2, read 4. 0777 ugo+rwx, 666, 0750,
  // $ stat -c "%a" some/folder/ => 0775
  err = os.MkdirAll(filepath.Dir(slug), 0775)
  if err != nil { log.Fatalf("Can't create directory `%v`: %v", filepath.Dir(slug), err) }

  // TODO: create another func for templating writting to a file.
  body, err = templateBoilerplate(body, &opt)
  if err != nil { log.Fatalf("Can't parse boilerplate:\n```\n%v\n```\n%v", body, err) }

  file, err := os.Create(slug)
  if err != nil {
    log.Fatalf("Cant create file %v: %v", slug, err)
  }
  defer file.Close()

  _, err = file.WriteString(body)
  if err != nil {
    log.Fatalf("Cant create file %v: %v", slug, err)
  }

  log.Printf("Created `%v` from `%v` boilerplate.", slug,name)
  // data, err := os.ReadFile(path)
  // m := make(map[string]int)
  // strings.Cut(::)


}

