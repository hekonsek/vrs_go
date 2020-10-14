package main

import (
	"fmt"
	"github.com/hekonsek/osexit"
	"github.com/hekonsek/ver"
	"github.com/spf13/cobra"
	"os"
	"github.com/fatih/color"
)

func init() {
	verCommand.AddCommand(bumpCommand)
}

var bumpCommand = &cobra.Command {
	Use:                "bump",
	Run: func(cmd *cobra.Command, args []string) {
		wd, err := os.Getwd()
		osexit.ExitOnError(err)

		oldVersion, err := ver.ReadCurrentVersion(wd)
		osexit.ExitOnError(err)

		err = ver.Bump(nil)
		osexit.ExitOnError(err)

		newVersion, err := ver.ReadCurrentVersion(wd)
		osexit.ExitOnError(err)

		fmt.Printf("Version %s bumped to version %s.", color.GreenString(oldVersion), color.GreenString(newVersion))
	},
}
