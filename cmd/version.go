package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show sts cli version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(RootCmd.Use + " " + VERSION)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
