package app

import (
  "os"
  "errors"

)

func DoesFileExist(path string) bool {
  _, err := os.Stat(path)
  return !errors.Is(err, os.ErrNotExist)
}

func IsValidFile(path string) bool {
  info, err := os.Stat(path)
  if err != nil { return false }
  if !info.Mode().IsRegular() { return false }
  if info.IsDir() { return false }
  return true
}

// func exists(path string) (bool, error) {
//     fileInfo, err := os.Stat(path)
//     fileInfo.IsDir()
//     if err == nil { return true, nil }
//     if os.IsNotExist(err) { return false, nil }
//     return false, err
// }

