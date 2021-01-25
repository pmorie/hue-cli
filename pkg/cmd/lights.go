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
	lightsGetCmd.PersistentFlags().IntVar(&getLightParams.id, "id", 1, "ID of the light to get")

	lightsCmd.AddCommand(lightsListCmd)
	lightsCmd.AddCommand(lightsGetCmd)

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

		fmt.Fprintln(w, "Name\tID\tOn?")
		fmt.Fprintln(w, "----\t---")

		for i := range lights {
			light := lights[i]
			fmt.Fprintf(w, "%v\t%v\t%v\n", light.Name, light.ID, light.State.On)
		}

		w.Flush()
	},
}

type getLightParamz struct {
	id int
}

var getLightParams = &getLightParamz{}

var lightsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get a specific light",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		bridgeIP, user := getLoginFromConfig()
		bridge := huego.New(bridgeIP, user)

		light, err := bridge.GetLight(getLightParams.id)
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
		fmt.Fprintf(w, "  XY:\t[%v, %v]\n", light.State.Xy[0], light.State.Xy[1])
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
