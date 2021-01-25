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
	lightsCmd.AddCommand(lightsListCmd)

	rootCmd.AddCommand(lightsCmd)
}

var lightsCmd = &cobra.Command{
	Use:   "lights",
	Short: "work with hue lights",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		// display help
	},
}

var lightsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list lights",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		bridgeIP, user := getLoginFromConfig()
		bridge := huego.New(bridgeIP, user)

		lights, err := bridge.GetLights()
		if err != nil {
			s := fmt.Sprintf("unable to connect to bridge: %v", err)
			panic(s)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

		sort.Slice(lights, func(i, j int) bool {
			if len(lights[i].Name) == len(lights[j].Name) {
				return lights[i].Name < lights[j].Name
			}

			return len(lights[i].Name) < len(lights[j].Name)
		})

		for i := range lights {
			light := lights[i]
			fmt.Fprintf(w, "%v\t%v\n", light.Name, light.State.On)
		}

		w.Flush()
	},
}
