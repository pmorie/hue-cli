package cmd

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

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
		bridge := huego.New(bridgeIP, user)

		groups, err := bridge.GetGroups()
		if err != nil {
			s := fmt.Sprintf("unable to connect to bridge: %v", err)
			panic(s)
		}

		sort.Slice(groups, func(i, j int) bool {
			if len(groups[i].Name) == len(groups[j].Name) {
				return groups[i].Name < groups[j].Name
			}

			return len(groups[i].Name) < len(groups[j].Name)
		})

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

		for i := range groups {
			group := groups[i]
			fmt.Fprintf(w, "%v\t%v\n", group.Name, group.GroupState.AnyOn)
		}

		w.Flush()
	},
}
