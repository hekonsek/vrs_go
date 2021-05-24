package main

import (
	"fmt"
	"github.com/hekonsek/osexit"
	"github.com/hekonsek/vrs/vrs"
	"github.com/spf13/cobra"
)

func init() {
	verCommand.AddCommand(currentCommand)
}

var currentCommand = &cobra.Command{
	Use: "current",
	Run: func(cmd *cobra.Command, args []string) {
		version, err := vrs.ReadCurrentVersion(nil)
		osexit.ExitOnError(err)

		fmt.Print(version)
	},
}
