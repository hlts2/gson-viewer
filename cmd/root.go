package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/hlts2/gson"
	gsonviewer "github.com/hlts2/gson-viewer/pkg/gson-viewer"
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

var path string

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&path, "set", "s", "", "set json file",
	)
}

func gsonViewer(cmd *cobra.Command, args []string) (err error) {
	finfo, err := os.Stdin.Stat()
	if err != nil {
		return err
	}

	var gson *gson.Gson

	// read from standard input
	if finfo.Mode()&os.ModeCharDevice == 0 {
		gson, err = gsonviewer.LoadWithReader(os.Stdin)
		if err != nil {
			return err
		}
	} else {
		if len(path) < 1 {
			return errors.New("json file dose not set")
		}

		gson, err = gsonviewer.Load(path)
		if err != nil {
			return err
		}

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
