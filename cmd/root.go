package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	cmdutil "github.com/hlts2/goson/cmd/util"
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

	jsonObj := []map[string]interface{}{}
	err = json.Unmarshal(jsonData, &jsonObj)
	if err != nil {
		return err
	}

	if len(selectKey) == 0 {
		jsonData, _ := json.Marshal(jsonObj)

		var buf bytes.Buffer
		json.Indent(&buf, jsonData, "", "  ")

		fmt.Println(buf.String())
		return nil
	}

	return nil
}

//Execute run command
func Execute() {
	rootCmd := NewRootCmd()
	cmdutil.CheckErr(rootCmd.Execute())
}
