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
	groupsGetCmd.PersistentFlags().IntVar(&getGroupParams.id, "id", 1, "ID of the group to get")

	groupsCmd.AddCommand(groupsListCmd)
	groupsCmd.AddCommand(groupsGetCmd)

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
		fmt.Fprintln(w, "Name\tID\tAny On?")
		fmt.Fprintln(w, "----\t----\t-------")

		for i := range groups {
			group := groups[i]
			fmt.Fprintf(w, "%v\t%v\t%v\n", group.Name, group.ID, group.GroupState.AnyOn)
		}

		w.Flush()
	},
}

type getGroupParamz struct {
	id int
}

var getGroupParams = &getGroupParamz{}

var groupsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get a specific group",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		bridgeIP, user := getLoginFromConfig()
		bridge := huego.New(bridgeIP, user)

		group, err := bridge.GetGroup(getGroupParams.id)
		if err != nil {
			s := fmt.Sprintf("unable to connect to bridge: %v", err)
			panic(s)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

		fmt.Fprintf(w, "Name:\t%v\n", group.Name)
		fmt.Fprintf(w, "ID:\t%v\n", group.ID)
		fmt.Fprintf(w, "Type:\t%v\n", group.Type)

		fmt.Fprintln(w, "State:\t")
		fmt.Fprintf(w, "  Any On:\t%v\n", group.GroupState.AnyOn)
		fmt.Fprintf(w, "  All On:\t%v\n", group.GroupState.AllOn)

		w.Flush()
	},
}
