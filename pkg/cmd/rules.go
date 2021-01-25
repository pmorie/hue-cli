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
	rulesCmd.AddCommand(rulesListCmd)

	rootCmd.AddCommand(rulesCmd)
}

var rulesCmd = &cobra.Command{
	Use:   "rules",
	Short: "work with hue rules",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		// display help
	},
}

var rulesListCmd = &cobra.Command{
	Use:   "list",
	Short: "list rules",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		bridgeIP, user := getLoginFromConfig()
		bridge := huego.New(bridgeIP, user)

		rules, err := bridge.GetRules()
		if err != nil {
			s := fmt.Sprintf("unable to connect to bridge: %v", err)
			panic(s)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

		sort.Slice(rules, func(i, j int) bool {
			if len(rules[i].Name) == len(rules[j].Name) {
				return rules[i].Name < rules[j].Name
			}

			return len(rules[i].Name) < len(rules[j].Name)
		})

		fmt.Fprintln(w, "Name\tStatus")
		fmt.Fprintln(w, "----\t------")

		for i := range rules {
			rule := rules[i]
			fmt.Fprintf(w, "%v\t%v\n", rule.Name, rule.Status)
		}

		w.Flush()
	},
}
