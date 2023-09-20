package app

import (
	"gopkg.in/yaml.v3"
)

type PageOptions struct {
  Path string `yaml:"path"`
  Timestamp bool `yaml:"timestamp"`
  Args []string `yaml:"args"`
  Command []string `yaml:"command"`
  Slug string `yaml:"slug"`
  // Basename string
  // Slug string
  // File string
  // Title string
  // filepath.Ext
  // filepath.Base
  // filepath.Dir
}

func parseOptions(txt string) (opt PageOptions,err error) {
  opt = defaultPageOptions()
  err = parseYamlOptions(&opt, txt)
  // TODO input flag options 
  return
}

func parseYamlOptions(opt *PageOptions, txt string) (err error) {
  err = yaml.Unmarshal([]byte(txt), opt)
  return
}

func defaultPageOptions() PageOptions {
  return PageOptions{
    Timestamp: false,
    Args: []string{},
    Command: []string{"title"},
    Slug: "{ safe(title) }{ ext }",
  }
}

/*
{b title b}
{b ext b}
{b date(%Y-%M-%D) b}

*/


// func parseOptions(yamlTxt string) {

//   var headerOpt PageOptions{
//     Ext: 
//   }
//   err = parseYamlOptions(&headerOpt,header)
//   err = parseYamlOptions(&headerOpt,header)

// }


// {
//       suffix: plate_path[/\.\w+$/],
//       name: plate_path[/.*(?=\.)/] || plate_path,
//       basename: basename,
//       title: basename,
//       slug: '{{ title }}',
//       path: '_posts/',
//       file: '{{ slug }}{{ suffix }}',
//     }

