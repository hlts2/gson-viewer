package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/hlts2/gson-viewer/pkg/gson-viewer"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gson-viewer",
	Short: "A CLI tool to view JSON",
	Run: func(cmd *cobra.Command, args []string) {
		if err := gsonViewer(cmd, args); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

var jsonFileN string

func init() {
	rootCmd.PersistentFlags().StringVarP(&jsonFileN, "set", "s", "", "set json file")
}

func gsonViewer(cmd *cobra.Command, args []string) error {
	if len(jsonFileN) < 1 {
		return errors.New("json file dose not set")
	}

	gson, err := jsonviewer.LoadJSON(jsonFileN)
	if err != nil {
		return err
	}

	return NewREPL(gson).Run()
}

// Execute execute gson-viewer
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
