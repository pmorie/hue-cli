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
	scenesCmd.AddCommand(scenesListCmd)

	rootCmd.AddCommand(scenesCmd)
}

var scenesCmd = &cobra.Command{
	Use:   "scenes",
	Short: "work with hue scenes",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		// display help
	},
}

var scenesListCmd = &cobra.Command{
	Use:   "list",
	Short: "list scenes",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		bridgeIP, user := getLoginFromConfig()
		bridge := huego.New(bridgeIP, user)

		scenes, err := bridge.GetScenes()
		if err != nil {
			s := fmt.Sprintf("unable to connect to bridge: %v", err)
			panic(s)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

		sort.Slice(scenes, func(i, j int) bool {
			if len(scenes[i].Name) == len(scenes[j].Name) {
				return scenes[i].Name < scenes[j].Name
			}

			return len(scenes[i].Name) < len(scenes[j].Name)
		})

		fmt.Fprintln(w, "Name\tType")
		fmt.Fprintln(w, "----\t----")

		for i := range scenes {
			scene := scenes[i]
			fmt.Fprintf(w, "%v\t%v\n", scene.Name, scene.Type)
		}

		w.Flush()
	},
}
