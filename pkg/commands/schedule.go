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
	scheduleCmd.AddCommand(scheduleListCmd)

	rootCmd.AddCommand(scheduleCmd)
}

var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "work with hue schedule",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		// display help
	},
}

var scheduleListCmd = &cobra.Command{
	Use:   "list",
	Short: "list schedule",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		bridgeIP, user := getLoginFromConfig()
		bridge := huego.New(bridgeIP, user)

		schedules, err := bridge.GetSchedules()
		if err != nil {
			s := fmt.Sprintf("unable to connect to bridge: %v", err)
			panic(s)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

		sort.Slice(schedules, func(i, j int) bool {
			if len(schedules[i].Name) == len(schedules[j].Name) {
				return schedules[i].Name < schedules[j].Name
			}

			return len(schedules[i].Name) < len(schedules[j].Name)
		})

		fmt.Fprintln(w, "Name\tStatus")
		fmt.Fprintln(w, "----\t------")

		for i := range schedules {
			schedule := schedules[i]
			fmt.Fprintf(w, "%v\t%v\n", schedule.Name, schedule.Status)
		}

		w.Flush()
	},
}
