package cmd

import (
	"fmt"

	"github.com/amimof/huego"
	"github.com/spf13/cobra"
)

func init() {
	bridgesCmd.AddCommand(discoverCmd)

	rootCmd.AddCommand(bridgesCmd)
}

var bridgesCmd = &cobra.Command{
	Use:   "bridges",
	Short: "work with hue bridges",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		// display help
	},
}

var discoverCmd = &cobra.Command{
	Use:   "discover",
	Short: "Discover hue bridges on your local network",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		bridges, err := huego.DiscoverAll()
		if err != nil {
			s := fmt.Sprintf("unable to discover bridges: %v", err)
			panic(s)
		}

		fmt.Println("HOST                   ID")
		fmt.Println("-------------------------------------------")

		for i := range bridges {
			bridge := bridges[i]
			fmt.Printf("%v         %v\n", bridge.Host, bridge.ID)
		}
	},
}
