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
	sensorsCmd.AddCommand(sensorsListCmd)

	rootCmd.AddCommand(sensorsCmd)
}

var sensorsCmd = &cobra.Command{
	Use:   "sensors",
	Short: "work with hue sensors",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		// display help
	},
}

var sensorsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list sensors",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		bridgeIP, user := getLoginFromConfig()
		bridge := huego.New(bridgeIP, user)

		sensors, err := bridge.GetSensors()
		if err != nil {
			s := fmt.Sprintf("unable to connect to bridge: %v", err)
			panic(s)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

		sort.Slice(sensors, func(i, j int) bool {
			if len(sensors[i].Name) == len(sensors[j].Name) {
				return sensors[i].Name < sensors[j].Name
			}

			return len(sensors[i].Name) < len(sensors[j].Name)
		})

		fmt.Fprintln(w, "Name\tType")
		fmt.Fprintln(w, "----\t----")

		for i := range sensors {
			sensor := sensors[i]
			fmt.Fprintf(w, "%v\t%v\n", sensor.Name, sensor.Type)
		}

		w.Flush()
	},
}
