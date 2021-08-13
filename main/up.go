package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/hekonsek/osexit"
	"github.com/hekonsek/vrs/vrs"
	"github.com/spf13/cobra"
)

var upVersionPrefix string
var upCommandProfiles []string

func init() {
	upCommand.Flags().StringVarP(&upVersionPrefix, "version-prefix", "v", "v", "")
	upCommand.Flags().StringSliceVar(&upCommandProfiles, "profile", []string{}, "")
	verCommand.AddCommand(upCommand)
}

var upCommand = &cobra.Command{
	Use: "up",
	Run: func(cmd *cobra.Command, args []string) {
		currentOptions, err := vrs.NewDefaultReadCurrentOptions()
		osexit.ExitOnError(err)
		currentOptions.VersionPrefix = upVersionPrefix

		oldVersion, isFileless, err := vrs.ReadCurrentVersion(currentOptions)
		osexit.ExitOnError(err)

		bumpOptions, err := vrs.NewDefaultBumpOptions()
		osexit.ExitOnError(err)
		bumpOptions.ActiveProfiles = upCommandProfiles
		bumpOptions.Fileless = isFileless
		bumpOptions.VersionPrefix = upVersionPrefix
		err = vrs.Bump(bumpOptions)
		osexit.ExitOnError(err)

		newVersion, _, err := vrs.ReadCurrentVersion(currentOptions)
		osexit.ExitOnError(err)

		fmt.Printf("Version %s bumped to version %s.\n", color.GreenString(oldVersion), color.GreenString(newVersion))
	},
}
