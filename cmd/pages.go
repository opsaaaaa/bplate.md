/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
  "fmt"
  "os"
  "log"

  "github.com/opsaaaaa/bplate.md/app"
  "github.com/spf13/cobra"
)

func init() {
  entries, err := os.ReadDir("./_boilerplates/")

  if err != nil {
    // TODO: provide some more usefull info if the director isn't present.
    log.Print(err)
    return
  }

  for _, d := range entries {
    if d.IsDir() { continue }

    var pagesCmd = &cobra.Command{
      Use: d.Name(),
      Short: fmt.Sprintf("Create a new page from `%s` boilerplate.",d.Name()),
      Long: `A longer description that spans multiple lines and likely contains examples
              and usage of using your command. For example:

              Cobra is a CLI library for Go that empowers applications.
              This application is a tool to generate the needed files
              to quickly create a Cobra application.`,

      Run: func(cmd *cobra.Command, args []string) {
        app.MakePage(d.Name(), args)
      },
    }

    rootCmd.AddCommand(pagesCmd)
  }

  // Here you will define your flags and configuration settings.

  // Cobra supports Persistent Flags which will work for this command
  // and all subcommands, e.g.:
  // pagesCmd.PersistentFlags().String("foo", "", "A help for foo")

  // Cobra supports local flags which will only run when this command
  // is called directly, e.g.:
  // pagesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
