/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
  "fmt"
  "os"

  // "path/filepath"

  "github.com/opsaaaaa/bplate.md/app"
  "github.com/spf13/cobra"
)

func RunRootCommand(cmd *cobra.Command, args []string) {
  fmt.Printf("%v <boilerplate> [...args]\n", cmd.Name())
  err := app.PrintList() 
  if err != nil {
    fmt.Printf("Failure: %v", err)
  }
}

func AutoCompleteRootCommand(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
  var comps []string
  if len(args) == 0 {
    comps = cobra.AppendActiveHelp(app.ListBoilerplateFiles(), "Specify a boilerplate to use.")
  } 
  return comps, cobra.ShellCompDirectiveNoFileComp
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "bplate.md",
  Short: "A brief description of your application",
  Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
  // Uncomment the following line if your bare application
  // has an action associated with it:
  Args: func(_ *cobra.Command, _ []string) error {
    // if args[0] == nil {
    //   return nil
    // } 
    return nil
  },
  Run: RunRootCommand,
  ValidArgsFunction: AutoCompleteRootCommand,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  // rootCmd.Run
  err := rootCmd.Execute()
  if err != nil {
    os.Exit(1)
  }
}

func init() {
  // Here you will define your flags and configuration settings.
  // Cobra supports persistent flags, which, if defined here,
  // will be global for your application.

  // rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bplate.md.yaml)")

  // Cobra also supports local flags, which will only run
  // when this action is called directly.
  rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


