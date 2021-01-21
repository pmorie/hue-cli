package cmd

import (
	"fmt"

	"github.com/amimof/huego"
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
		bridgeIP, user := getLoginFromConfig()

		fmt.Printf("Bridge IP: %v\n", bridgeIP)
		fmt.Printf("User: %v\n", user)

		bridge := huego.New(bridgeIP, user)

		groups, err := bridge.GetGroups()
		if err != nil {
			s := fmt.Sprintf("unable to connect to bridge: %v", err)
			panic(s)
		}

		for i := range groups {
			group := groups[i]
			fmt.Printf("%v\t\t\t%v\n", group.Name, group.GroupState.AnyOn)
		}
	},
}
