/*
  Extract the header and template from a boilerplate file.
*/

package app

import (
  "bufio"
  "os"
  "strings"
)

const (
  BEFORE_HEADER = 0
  INSIDE_HEADER = 1
  AFTER_HEADER = 2
)

func readBoilerplate(path string) (template string, header string, err error) {
    inHeader := BEFORE_HEADER

    file, err := os.Open(path)

    if err != nil { 
      file.Close()
      return
    }

    fileScanner := bufio.NewScanner(file)

    for fileScanner.Scan() {
      line := fileScanner.Text() + NEWLINE

      switch inHeader {
      case BEFORE_HEADER:
        if strings.HasPrefix(line, "_boilerplate:") {
          inHeader = INSIDE_HEADER
          continue
        }
        template += line
      case INSIDE_HEADER:
        if line == NEWLINE { continue }
        if !strings.HasPrefix(line, DOUBLESPACE) && !strings.HasPrefix(line, TAB) {
          inHeader = AFTER_HEADER
          template += line
          continue
        }
        header += strings.TrimPrefix(strings.TrimPrefix(line, TAB), DOUBLESPACE)
      default:
        template += line
      }

    }

    if err = fileScanner.Err(); err != nil {
      file.Close()
      return
    }

    file.Close()
    return
}



