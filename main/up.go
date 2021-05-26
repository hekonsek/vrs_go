package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/hekonsek/osexit"
	"github.com/hekonsek/vrs/vrs"
	"github.com/spf13/cobra"
)

var upCommandProfiles []string

func init() {
	upCommand.Flags().StringSliceVar(&upCommandProfiles, "profile", []string{}, "")
	verCommand.AddCommand(upCommand)
}

var upCommand = &cobra.Command{
	Use: "up",
	Run: func(cmd *cobra.Command, args []string) {
		oldVersion, err := vrs.ReadCurrentVersion(nil)
		osexit.ExitOnError(err)

		bumpOptions, err := vrs.NewDefaultBumpOptions()
		osexit.ExitOnError(err)
		bumpOptions.ActiveProfiles = upCommandProfiles
		err = vrs.Bump(bumpOptions)
		osexit.ExitOnError(err)

		newVersion, err := vrs.ReadCurrentVersion(nil)
		osexit.ExitOnError(err)

		fmt.Printf("Version %s bumped to version %s.\n", color.GreenString(oldVersion), color.GreenString(newVersion))
	},
}
