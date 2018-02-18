package main

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/hlts2/goson"
	cmdutil "github.com/hlts2/goson/cmd/goson/util"
	"github.com/spf13/cobra"
)

var (
	selectKey string

	rootCmdUsageTemplate = ``
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

	rootCmd.PersistentFlags().StringVarP(&selectKey, "select", "s", "", "set json key to extract json value")
	return rootCmd
}

func runRootCmd(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New(rootCmdUsageTemplate)
	}

	jsonData, err := ioutil.ReadFile(args[0])
	if err != nil {
		return err
	}

	g, err := goson.NewGoson(jsonData)
	if err != nil {
		return nil
	}

	if len(selectKey) == 0 {
		str, err := g.JSONObjectToPrettyJSONString()
		if err != nil {
			return err
		}
		fmt.Println(str)
	}

	return nil
}

//execute run command
func execute() {
	rootCmd := NewRootCmd()
	cmdutil.CheckErr(rootCmd.Execute())
}
