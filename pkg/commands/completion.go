package commands

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	addCompletion(rootCmd)
}

func addCompletion(topLevel *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "completion",
		Short: "Generates bash completion scripts",
		Long: `To load completion run

. <(hue-cli completion)

To configure your bash shell to load completions for each session add to your bashrc

# ~/.bashrc or ~/.profile
. <(hue-cli completion)
`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = topLevel.GenBashCompletion(os.Stdout)
		},
	}

	topLevel.AddCommand(cmd)
}
