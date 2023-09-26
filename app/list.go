package app

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"
)


func PrintList() (err error) {

  list := ListBoilerplateFiles()

  for _, path := range list {
    fmt.Printf("    %v\n", path)
  }

  return
}



func ListBoilerplateFiles() (list []string) {

  rgx := regexp.MustCompile(`\.boilerplates\/|\/{0,1}\.boilerplate(\..*){0,1}`)

  // .boilerplate/**
  _ = filepath.WalkDir(BOILERPLATES_FOLDER,func(path string, info fs.DirEntry, err error) error {
    if info.Type().IsRegular() {
      list = append(list, rgx.ReplaceAllString(path, EMPTYSTRING))
    }
    return err
  })

  // **/*.boilerplate*
  _ = filepath.WalkDir("./", func(path string, info fs.DirEntry, err error) error {

    if info.IsDir() && len(info.Name()) > 1 && strings.HasPrefix(info.Name(),".") {
      return filepath.SkipDir
    }

    if info.Type().IsRegular() && strings.Contains(info.Name(), BOILERPLATE_FILE) {
      list = append(list, rgx.ReplaceAllString(path, EMPTYSTRING))
    }
    return err
  })

  return list
}

