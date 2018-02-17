package util

import (
	"fmt"
	"os"
)

//CheckErr check error.
func CheckErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
