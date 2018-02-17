package cmd

import (
	cmdutil "github.com/hlts2/goson/cmd/util"
	"github.com/spf13/cobra"
)

//NewRootCmd returns Command pointer
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Short: "A CLI tool for simple json viewer",
		Use:   "goson",
	}
	return rootCmd
}

//Execute run command
func Execute() {
	rootCmd := NewRootCmd()
	cmdutil.CheckErr(rootCmd.Execute())
}
