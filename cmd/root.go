package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gson-viewer",
	Short: "A CLI tool to view JSON",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var jsonFileN string

func init() {
	rootCmd.PersistentFlags().StringVarP(&jsonFileN, "set", "s", "", "set json file")
}

// Execute execute gson-viewer
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
