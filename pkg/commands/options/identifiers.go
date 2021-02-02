package options

import (
	"github.com/spf13/cobra"
)

// IDOptions
type IDOptions struct {
	ID     int
	Name string
}

func AddIDArgs(cmd *cobra.Command, o *IDOptions) {
	cmd.Flags().IntVar(&o.ID, "id", 0,
		"The id.")
	cmd.Flags().StringVar(&o.Name, "name", "",
		"The name.")
}
