package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type PageOptions struct {

}

const (
  TAB = "	"
  DOUBLESPACE = "  "
  NEWLINE = "\n"
)

func MakePage(name string, options PageOptions) {
  path := filepath.Join("_boilerplates", name)
  header, template := ReadBoilerplate(path)
  fmt.Println(header)
  fmt.Println(template)

  // data, err := os.ReadFile(path)

  // strings.Cut(::)


}

const (
  BEFORE_HEADER = 0
  INSIDE_HEADER = 1
  AFTER_HEADER = 2
)
/*
  Extract the header and template from the boilerplate file.
*/
func ReadBoilerplate(path string) (header string, template string) {
    inHeader := BEFORE_HEADER

    file, err := os.Open(path)

    if err != nil {
        log.Fatalf("Error when opening file: %s", err)
    }

    fileScanner := bufio.NewScanner(file)

    for fileScanner.Scan() {
      line := fileScanner.Text()+NEWLINE

      switch inHeader {
      case BEFORE_HEADER:
        if strings.HasPrefix(line, "_boilerplate:") {
          inHeader = INSIDE_HEADER
          header += line
          continue
        }
        template += line
      case INSIDE_HEADER:
        if line != NEWLINE && !strings.HasPrefix(line, DOUBLESPACE) && !strings.HasPrefix(line, TAB) {
          inHeader = AFTER_HEADER
          template += line
          continue
        }
        header += line
      default:
        template += line
      }

    }

    if err := fileScanner.Err(); err != nil {
        log.Fatalf("Error while reading file: %s", err)
    }

    file.Close()
    return
}





