package main

import (
	"fmt"
	"github.com/hekonsek/osexit"
	"github.com/hekonsek/ver"
	"github.com/spf13/cobra"
)

func init() {
	verCommand.AddCommand(currentCommand)
}

var currentCommand = &cobra.Command{
	Use: "current",
	Run: func(cmd *cobra.Command, args []string) {
		version, err := ver.ReadCurrentVersion(nil)
		osexit.ExitOnError(err)

		fmt.Println(version)
	},
}
