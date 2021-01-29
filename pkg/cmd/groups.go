package cmd

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"

	"github.com/amimof/huego"
	"github.com/erikh/colorwriter"
	"github.com/gookit/color"
	pcolor "github.com/pmorie/hue-cli/pkg/color"
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

		w := colorwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
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
		fmt.Fprintf(w, "Class:\t%v\n", group.Class)

		fmt.Fprintln(w, "State:\t")
		fmt.Fprintf(w, "  Any On:\t%v\n", group.GroupState.AnyOn)
		fmt.Fprintf(w, "  All On:\t%v\n", group.GroupState.AllOn)

		fmt.Fprintf(w, "Recycle:\t%v\n", group.Recycle)

		// TODO sort

		fmt.Fprintln(w, "Lights:\t")
		for i := range group.Lights {

			id, err := strconv.Atoi(group.Lights[i])
			if err != nil {
				panic(err)
			}

			light, err := bridge.GetLight(id)
			if err != nil {
				panic(err)
			}

			bold := func(str string) string {
				if light.State.On {
					if len(light.State.Xy) < 2 {
						return color.OpBold.Render(str)
					} else {
						return color.OpBold.Render(color.RGB(pcolor.HueToRGB(light)).Sprint(str))
					}
				}

				return str
			}

			stateStr := "off"
			if light.State.On {
				stateStr = "on"
			}

			fmt.Fprintf(w, "  %v:\t%v\n", light.Name, bold(stateStr))
		}

		w.Flush()
	},
}
