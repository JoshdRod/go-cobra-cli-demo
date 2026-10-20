/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "The short description is shown when you run `go run main.go --help`",
	Long: `The long description

which usually spans
multiple
lines

is shown when you run ` + "`go run main.go hello --help`",
	Run: func(cmd *cobra.Command, args []string) {
		yay, err := cmd.Flags().GetBool("yay")
		if err != nil || yay != true {
			fmt.Println("RAAAAAHHHH")
		} else {
			fmt.Println("yay!")
		}

		num, err := cmd.Flags().GetInt("num")
		if err != nil {
			fmt.Println("Got no num :(")
		} else {
			fmt.Println(num)
		}
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	helloCmd.Flags().BoolP("yay", "y", false, "Changes message to yay") // BoolP is like Bool, but with a shorthand
	helloCmd.Flags().IntP("num", "n", 0, "An integer that's printed")
}
