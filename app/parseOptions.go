package app

import (
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type PageOptions struct {
  Path string `yaml:"path"`
  Args []string `yaml:"args"`
  Command []string `yaml:"command"`
  // Defaults []string `yaml:"defaults"`
  Slug string `yaml:"slug"`
  Ext string `yaml:"ext"`
  Boilerplate string
  // Basename string
  // Slug string
  // File string
  // Title string
  // filepath.Ext
  // filepath.Base
  // filepath.Dir
}

func parseOptions(name string, txt string, args []string) (opt PageOptions,err error) {
  opt = defaultPageOptions()
  opt.Ext = filepath.Ext(name)
  err = parseYamlOptions(&opt, txt)
  if len(args) > 0 {
    opt.Args = args
  }
  opt.Boilerplate = name
  // TODO input flag options 
  return
}

func parseYamlOptions(opt *PageOptions, txt string) (err error) {
  err = yaml.Unmarshal([]byte(txt), opt)
  return
}

func defaultPageOptions() PageOptions {
  return PageOptions{
    Args: []string{},
    Command: []string{"title"},
    Slug: "{{ path }}/{{ title | asFileSlug }}{{ ext }}",
    Path: "./",
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

