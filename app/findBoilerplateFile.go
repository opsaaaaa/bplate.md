package app

import (
	"os"
	"path/filepath"
)

/*
  for the given path, check all the possible boilerplate locations and return the first valid match.
  ./path
  .boilerplates/path
  path/.boilerplate
  path/.boilerplate.*
*/
func findBoilerplateFile(path string) (file string, err error)  {

  info, err := os.Stat(path)

  if err != nil {
    file = filepath.Join(BOILERPLATES_FOLDER, path)
    info, err = os.Stat(file)
    if err == nil && info.Mode().IsRegular() { return }
  } else if info.Mode().IsRegular() {
    file = path
    return
  } else if info.IsDir() {
    file = filepath.Join(path, BOILERPLATE_FILE)
    info, err = os.Stat(file)
    if err == nil && info.Mode().IsRegular() { return }

    var files []string
    files, err = filepath.Glob(filepath.Join(path, BOILERPLATE_FILE_GLOB))
    if err != nil { return }

    for _, file = range files {
      info, err = os.Stat(file)
      if err == nil && info.Mode().IsRegular() { return }
    }
  }

  file = ""
  err = os.ErrNotExist
  return
}
