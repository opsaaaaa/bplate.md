package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CreatePage(name string, args []string) {
  // var err error
  // DoesFileExist(name)
  // boilerplateFile := filepath.Join(BOILERPLATES_FOLDER, name)
  boilerplateFile, err := findBoilerplateFile(name)
  if err != nil {
    fmt.Printf("A boilerpate called `%v` doesn't exist.\nuse the `--new` flag to create one.\n%v\n",name,err)
    os.Exit(1)
    return
  }

  body, header, err := readBoilerplate(boilerplateFile)
  if err != nil { log.Fatal(err) }

  opt, err := parseOptions(name,header,args)
  if err != nil { log.Fatal(err) }

  // TODO: create a func for templating a string
  slug, err := templateBoilerplate(opt.Slug, &opt)
  if err != nil { log.Fatalf("Can't parse slug:\n`%v`\n%v", opt.Slug, err) }

  if DoesFileExist(slug) {
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
  // data, err := os.ReadFile(boilerplateFile)
  // m := make(map[string]int)
  // strings.Cut(::)


}

