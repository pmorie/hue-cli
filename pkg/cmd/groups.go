package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	groupsCmd.AddCommand(groupsListCmd)

	rootCmd.AddCommand(groupsCmd)
}

var groupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "work with hue groups",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		// display help
	},
}

var groupsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list groups",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		// load config
		// login with user
	},
}
