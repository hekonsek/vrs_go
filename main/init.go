package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/hekonsek/osexit"
	"github.com/hekonsek/vrs/vrs"
	"github.com/spf13/cobra"
)

func init() {
	verCommand.AddCommand(initCommand)
}

var initCommand = &cobra.Command{
	Use: "init",
	Run: func(cmd *cobra.Command, args []string) {
		err := vrs.Init(nil)
		osexit.ExitOnError(err)

		fmt.Printf("Created and commited %s file.\n", color.GreenString(vrs.VrsConfigFileName))
	},
}
