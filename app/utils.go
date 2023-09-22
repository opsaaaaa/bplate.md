package app

import (
  "os"
  "errors"

)

func DoesFileExist(path string) bool {
  _, err := os.Stat(path)
  return !errors.Is(err, os.ErrNotExist)
}

