package main

import (
	"github.com/hekonsek/osexit"
	"github.com/spf13/cobra"
)

var verCommand = &cobra.Command{
	Use: "vrs",
	Run: func(cmd *cobra.Command, args []string) {
		osexit.ExitOnError(cmd.Help())
	},
}

func main() {
	osexit.ExitOnError(verCommand.Execute())
}
