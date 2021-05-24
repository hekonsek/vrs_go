package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/hekonsek/osexit"
	"github.com/hekonsek/vrs/vrs"
	"github.com/spf13/cobra"
)

func init() {
	verCommand.AddCommand(bumpCommand)
}

var bumpCommand = &cobra.Command{
	Use: "bump",
	Run: func(cmd *cobra.Command, args []string) {
		oldVersion, err := vrs.ReadCurrentVersion(nil)
		osexit.ExitOnError(err)

		err = vrs.Bump(nil)
		osexit.ExitOnError(err)

		newVersion, err := vrs.ReadCurrentVersion(nil)
		osexit.ExitOnError(err)

		fmt.Printf("Version %s bumped to version %s.\n", color.GreenString(oldVersion), color.GreenString(newVersion))
	},
}
