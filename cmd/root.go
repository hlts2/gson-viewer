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
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(runRootCmd(cmd, args))
		},
	}
	return rootCmd
}

func runRootCmd(cmd *cobra.Command, args []string) error {
	return nil
}

//Execute run command
func Execute() {
	rootCmd := NewRootCmd()
	cmdutil.CheckErr(rootCmd.Execute())
}
