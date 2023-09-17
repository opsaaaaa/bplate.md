package app

import (
	"gopkg.in/yaml.v3"
)

type PageOptions struct {
  Path string `yaml:"path"`
  Timestamp bool `yaml:"timestamp"`
}

func parseYamlOptions(txt string) (opt PageOptions, err error) {
  err = yaml.Unmarshal([]byte(txt), &opt)
  return
}


