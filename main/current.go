package main

import (
	"fmt"
	"github.com/hekonsek/osexit"
	"github.com/hekonsek/vrs/vrs"
	"github.com/spf13/cobra"
)

var currentVersionPrefix string

func init() {
	currentCommand.Flags().StringVarP(&currentVersionPrefix, "version-prefix", "v", "v", "")
	verCommand.AddCommand(currentCommand)
}

var currentCommand = &cobra.Command{
	Use: "current",
	Run: func(cmd *cobra.Command, args []string) {
		o, err := vrs.NewDefaultReadCurrentOptions()
		osexit.ExitOnError(err)
		o.VersionPrefix = currentVersionPrefix
		version, _, err := vrs.ReadCurrentVersion(o)
		osexit.ExitOnError(err)

		fmt.Println(version)
	},
}
