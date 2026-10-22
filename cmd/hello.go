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

		colour, err := cmd.Flags().GetStringSlice("colour")
		if err != nil {
			fmt.Println("Error getting favourite colours")
			fmt.Println(err)
		} else {
			fmt.Println(colour)
			fmt.Println("My favourite colours are:")
			for _, c := range colour { // This is a foreach loop in go - arg 0 = index, arg 1 = value
				fmt.Println(c)
			}
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
	// An array of values can be passed in using the TSlice function
	// (where T is the type of the array)
	// These arrays can the be passed in using one of two methods:
	// 1. --colours="red,yellow,green"
	// 2. --colours red --colours yellow --colours green
	// BUT NOT --colours red yellow green - That doesn't work
	helloCmd.Flags().StringSliceP("colour", "c", []string{}, "A list of favourite colours")
}

