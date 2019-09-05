package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wayneashleyberry/numeronym"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:  "numeronym",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(numeronym.FromString(args[0]))
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
