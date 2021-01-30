package options

import (
	"github.com/spf13/cobra"
)

// LightOptions
type LightOptions struct {
	Brightness     uint8
	// TODO: hue, color, temp, etc...
}

func AddLightArgs(cmd *cobra.Command, o *LightOptions) {
	cmd.Flags().Uint8VarP(&o.Brightness, "brightness", "b", 0, "Brightness of light, 0-254")
}
