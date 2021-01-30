package commands

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/spf13/cobra"
	
	"github.com/pmorie/hue-cli/pkg/commands/options"
)

func init() {
	addLights(rootCmd)
}

func addLights(topLevel *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "lights",
		Short: "work with hue lights",
		Long:  `TODO`,
		Run: func(cmd *cobra.Command, args []string) {
			// display help
			cmd.Help()
		},
	}

	addLightsList(cmd)
	addLightsGet(cmd)
	addLightsSet(cmd)
	addLightsToggle(cmd)

	topLevel.AddCommand(cmd)
}


func addLightsList(topLevel *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list lights",
		Long:  "TODO",
		Run: func(cmd *cobra.Command, args []string) {
			bridge := getBridgeFromConfig()

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

			fmt.Fprintln(w, "Name\tID\tOn?")
			fmt.Fprintln(w, "----\t---")

			for i := range lights {
				light := lights[i]
				fmt.Fprintf(w, "%v\t%v\t%v\n", light.Name, light.ID, light.State.On)
			}

			w.Flush()
		},
	}
	topLevel.AddCommand(cmd)
}

func addLightsGet(topLevel *cobra.Command) {
	io := &options.IDOptions{}
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get a specific light",
		Long:  "TODO",
		Run: func(cmd *cobra.Command, args []string) {
			bridge := getBridgeFromConfig()

			light, err := bridge.GetLight(io.ID)
			if err != nil {
				s := fmt.Sprintf("unable to connect to bridge: %v", err)
				panic(s)
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

			fmt.Fprintf(w, "Name:\t%v\n", light.Name)
			fmt.Fprintf(w, "ID:\t%v\n", light.ID)
			fmt.Fprintf(w, "Unique ID:\t%v\n", light.UniqueID)
			fmt.Fprintf(w, "Product ID:\t%v\n", light.ProductID)
			fmt.Fprintf(w, "Model ID:\t%v\n", light.ModelID)
			fmt.Fprintf(w, "Type:\t%v\n", light.Type)
			fmt.Fprintf(w, "Manufacturer Name:\t%v\n", light.ManufacturerName)
			fmt.Fprintf(w, "Software Version:\t%v\n", light.SwVersion)
			fmt.Fprintf(w, "Software Config ID:\t%v\n", light.SwConfigID)
			fmt.Fprintln(w, "State:\t")
			fmt.Fprintf(w, "  On:\t%v\n", light.State.On)
			fmt.Fprintf(w, "  Brightness:\t%v\n", light.State.Bri)
			fmt.Fprintf(w, "  Hue:\t%v\n", light.State.Hue)
			fmt.Fprintf(w, "  Saturation:\t%v\n", light.State.Sat)
			if len(light.State.Xy) >= 2 {
				fmt.Fprintf(w, "  XY:\t[%v, %v]\n", light.State.Xy[0], light.State.Xy[1])
			}
			fmt.Fprintf(w, "  Color Temperature:\t%v\n", light.State.Ct)
			fmt.Fprintf(w, "  Alert:\t%v\n", light.State.Alert)
			fmt.Fprintf(w, "  Effect:\t%v\n", light.State.Effect)
			fmt.Fprintf(w, "  Transition Time:\t%v\n", light.State.TransitionTime)
			fmt.Fprintf(w, "  Brightness Increment:\t%v\n", light.State.BriInc)
			fmt.Fprintf(w, "  Saturation Increment:\t%v\n", light.State.SatInc)
			fmt.Fprintf(w, "  Hue Increment:\t%v\n", light.State.HueInc)
			fmt.Fprintf(w, "  Color Temperature Increment:\t%v\n", light.State.CtInc)
			fmt.Fprintf(w, "  XY Increment:\t%v\n", light.State.XyInc)
			fmt.Fprintf(w, "  Color Mode:\t%v\n", light.State.ColorMode)
			fmt.Fprintf(w, "  Reachable:\t%v\n", light.State.Reachable)
			fmt.Fprintf(w, "  Scene:\t%v\n", light.State.Scene)

			w.Flush()
		},
	}

	options.AddIDArgs(cmd, io)

	topLevel.AddCommand(cmd)
}

func addLightsSet(topLevel *cobra.Command) {
	io := &options.IDOptions{}
	l := &options.LightOptions{}

	cmd := &cobra.Command{
		Use:   "set",
		Short: "set a specific light",
		Long:  "TODO",
		Run: func(cmd *cobra.Command, args []string) {
			bridge := getBridgeFromConfig()

			light, err := bridge.GetLight(io.ID)
			if err != nil {
				s := fmt.Sprintf("unable to connect to bridge: %v", err)
				panic(s)
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

			fmt.Fprintf(w, "Name:\t%v\n", light.Name)
			fmt.Fprintf(w, "ID:\t%v\n", light.ID)
			fmt.Fprintf(w, "Type:\t%v\n", light.Type)
			fmt.Fprintln(w, "State:\t")

			if err := light.Bri(l.Brightness); err != nil {
				s := fmt.Sprintf("unable to set brightness: %v", err)
				panic(s)
			}

			fmt.Fprintf(w, "  On:\t%v\n", light.State.On)
			fmt.Fprintf(w, "  Brightness:\t%v\n", light.State.Bri)

			w.Flush()
		},
	}

	options.AddIDArgs(cmd, io)
	options.AddLightArgs(cmd, l)

	topLevel.AddCommand(cmd)
}

func addLightsToggle(topLevel *cobra.Command) {
	io := &options.IDOptions{}

	cmd := &cobra.Command{
		Use:   "toggle",
		Short: "toggle a specific light",
		Long:  "TODO",
		Run: func(cmd *cobra.Command, args []string) {
			bridge := getBridgeFromConfig()

			light, err := bridge.GetLight(io.ID)
			if err != nil {
				s := fmt.Sprintf("unable to connect to bridge: %v", err)
				panic(s)
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

			fmt.Fprintf(w, "Name:\t%v\n", light.Name)
			fmt.Fprintf(w, "ID:\t%v\n", light.ID)
			fmt.Fprintf(w, "Type:\t%v\n", light.Type)
			fmt.Fprintln(w, "State:\t")

			if light.IsOn() {
				light.Off()
			} else {
				light.On()
			}

			fmt.Fprintf(w, "  On:\t%v\n", light.State.On)

			w.Flush()
		},
	}

	options.AddIDArgs(cmd, io)

	topLevel.AddCommand(cmd)
}