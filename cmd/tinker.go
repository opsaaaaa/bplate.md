/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

import "github.com/opsaaaaa/bplate.md/cmd"
*/
package cmd

import (
  "fmt"
  "os"
  "log"

  "github.com/spf13/cobra"
)

// tinkerCmd represents the tinker command
var tinkerCmd = &cobra.Command{
  Use:   "tinker",
  Short: "A brief description of your command",
  Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("tinker called")
    entries, err := os.ReadDir("./")
    if err != nil { log.Fatal(err) }
    for _, e := range entries {
      if e.IsDir() { continue }

      fmt.Println(e.Name())
    }
  },
}

func init() {
  rootCmd.AddCommand(tinkerCmd)

  // Here you will define your flags and configuration settings.

  // Cobra supports Persistent Flags which will work for this command
  // and all subcommands, e.g.:
  // tinkerCmd.PersistentFlags().String("foo", "", "A help for foo")

  // Cobra supports local flags which will only run when this command
  // is called directly, e.g.:
  // tinkerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
