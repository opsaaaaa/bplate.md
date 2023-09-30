package app

import (
	"os"
	"path/filepath"
)



func CreateNewBoilerplate(path string, args []string) (file string, err error) {
  /*
    So what does this do again

    .boilerplates/path already esists > error, that boilerplate aready esists
    .boilerplates/path is a folder > create the boilerplate file without complaint.
    .boilerplates/path doesn't exist. > create it without complaint.
  */
  file = filepath.Join(BOILERPLATES_FOLDER, path)
  _, err = os.Stat(file)

  if !os.IsNotExist(err) {
    if err == nil {
      err = os.ErrExist
    }
    return
  }
  err = os.MkdirAll(filepath.Dir(file), 0777)
  if err != nil { return }

  newfile, err := os.Create(file)
  if err != nil { return }
  defer newfile.Close()

  _, err = newfile.WriteString("This is a placeholder for the new file.")
  if err != nil { return }

  // check how other go pkgs are installed.

  // create the file
  // print a message
  return
}




