package commands

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/amimof/huego"
	"github.com/spf13/cobra"
)

func init() {
	resourceLinksCmd.AddCommand(resourceLinksListCmd)

	rootCmd.AddCommand(resourceLinksCmd)
}

var resourceLinksCmd = &cobra.Command{
	Use:   "resourcelinks",
	Short: "work with hue resourceLinks",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		// display help
	},
}

var resourceLinksListCmd = &cobra.Command{
	Use:   "list",
	Short: "list resourcelinks",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		bridgeIP, user := getLoginFromConfig()
		bridge := huego.New(bridgeIP, user)

		resourceLinks, err := bridge.GetResourcelinks()
		if err != nil {
			s := fmt.Sprintf("unable to connect to bridge: %v", err)
			panic(s)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

		sort.Slice(resourceLinks, func(i, j int) bool {
			if len(resourceLinks[i].Name) == len(resourceLinks[j].Name) {
				return resourceLinks[i].Name < resourceLinks[j].Name
			}

			return len(resourceLinks[i].Name) < len(resourceLinks[j].Name)
		})

		fmt.Fprintln(w, "Name\tType")
		fmt.Fprintln(w, "----\t----")

		for i := range resourceLinks {
			resourceLink := resourceLinks[i]
			fmt.Fprintf(w, "%v\t%v\n", resourceLink.Name, resourceLink.Type)
		}

		w.Flush()
	},
}
