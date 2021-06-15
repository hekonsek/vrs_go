package main

import (
	"github.com/hekonsek/osexit"
	"github.com/spf13/cobra"
)

var verCommand = &cobra.Command{
	Use:   "vrs",
	Short: "vrs - project versioning made easy",
	Long:  "vrs is a command line tool simplifying versioning of git projects.",

	Run: func(cmd *cobra.Command, args []string) {
		osexit.ExitOnError(cmd.Help())
	},
}

func main() {
	osexit.ExitOnError(verCommand.Execute())
}
