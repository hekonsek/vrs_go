package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/hekonsek/osexit"
	"github.com/hekonsek/ver"
	"github.com/spf13/cobra"
)

func init() {
	verCommand.AddCommand(initCommand)
}

var initCommand = &cobra.Command{
	Use: "init",
	Run: func(cmd *cobra.Command, args []string) {
		err := ver.Init(nil)
		osexit.ExitOnError(err)

		fmt.Printf("Created and commited %s file.\n", color.GreenString(ver.VersioonConfigFileName))
	},
}
