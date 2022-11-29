/*
Copyright © 2022 Skyfall
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new \"package.json\" file.",
	Long: `The "init" command allows you to create a new "package.json" file in
working directory. It will ask you some questions, and if the "--yes" flag is not
set, it will ask whether you want to write the file to your working directory or not.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
