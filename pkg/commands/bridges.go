package commands

import (
	"bufio"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/amimof/huego"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	setupCmd.PersistentFlags().StringVar(&setupParams.bridgeIP, "bridgeIP", "", "IP of the bridge to setup")
	setupCmd.PersistentFlags().StringVar(&setupParams.user, "user", "hue-cli", "user to setup")
	setupCmd.PersistentFlags().BoolVar(&setupParams.wait, "wait", true, "user to setup")

	statusCmd.PersistentFlags().BoolVar(&statusParams.verbose, "verbose", false, "verbosity")

	bridgeCmd.AddCommand(setupCmd)
	bridgeCmd.AddCommand(discoverCmd)
	bridgeCmd.AddCommand(statusCmd)
	bridgeCmd.AddCommand(capabilitiesCmd)

	rootCmd.AddCommand(bridgeCmd)
}

var bridgeCmd = &cobra.Command{
	Use:   "bridge",
	Short: "Work with hue bridges",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		// display help
	},
}

var discoverCmd = &cobra.Command{
	Use:   "discover",
	Short: "Discover hue bridges on your network",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		bridges, err := huego.DiscoverAll()
		if err != nil {
			s := fmt.Sprintf("unable to discover bridges: %v", err)
			panic(s)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

		fmt.Fprintln(w, "IP\tID")

		for i := range bridges {
			bridge := bridges[i]
			fmt.Fprintf(w, "%v\t%v\n", bridge.Host, bridge.ID)
		}

		w.Flush()
	},
}

type statusParamz struct {
	verbose bool
}

var statusParams = statusParamz{}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the status of the configured bridge",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		bridgeIP, user := getLoginFromConfig()
		bridge := huego.New(bridgeIP, user)
		bridgeConfig, err := bridge.GetConfig()
		if err != nil {
			s := fmt.Sprintf("unable to connect to bridge: %v", err)
			panic(s)
		}

		fmt.Printf("Bridge Configuration:\n\n")
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintf(w, "Name:\t%v\n", bridgeConfig.Name)
		fmt.Fprintln(w, "Software Update Configuration:\t")
		fmt.Fprintf(w, "  Check For Updates:\t%v\n", bridgeConfig.SwUpdate2.CheckForUpdate)
		fmt.Fprintf(w, "  State:\t%v\n", bridgeConfig.SwUpdate2.State)
		fmt.Fprintf(w, "  AutoInstall:\t\n")
		fmt.Fprintf(w, "    Enabled:\t%v\n", bridgeConfig.SwUpdate2.AutoInstall.On)
		fmt.Fprintf(w, "    Last Update Time:\t%v\n", bridgeConfig.SwUpdate2.AutoInstall.UpdateTime)
		fmt.Fprintf(w, "  Last Change:\t%v\n", bridgeConfig.SwUpdate2.LastChange)
		fmt.Fprintf(w, "  Last Install:\t%v\n", bridgeConfig.SwUpdate2.LastInstall)

		if statusParams.verbose {
			fmt.Fprintln(w, "Allowed Clients:\t")
			for k, _ := range bridgeConfig.WhitelistMap {
				fmt.Fprintf(w, "  %v\t\n", k)
			}
		} else {
			fmt.Fprintf(w, "Allowed clients:\t%v\n", len(bridgeConfig.WhitelistMap))
		}

		fmt.Fprintln(w, "Portal State:\t")
		fmt.Fprintf(w, "  Signed On:\t%v\n", bridgeConfig.PortalState.SignedOn)
		fmt.Fprintf(w, "  Incoming:\t%v\n", bridgeConfig.PortalState.Incoming)
		fmt.Fprintf(w, "  Outgoing:\t%v\n", bridgeConfig.PortalState.Outgoing)
		fmt.Fprintf(w, "  Communication:\t%v\n", bridgeConfig.PortalState.Communication)
		fmt.Fprintf(w, "API Version:\t%v\n", bridgeConfig.APIVersion)
		fmt.Fprintf(w, "SW Version:\t%v\n", bridgeConfig.SwVersion)
		fmt.Fprintf(w, "Proxy Address:\t%v\n", bridgeConfig.ProxyAddress)
		fmt.Fprintf(w, "Proxy Port:\t%v\n", bridgeConfig.ProxyPort)
		fmt.Fprintf(w, "Link Button:\t%v\n", bridgeConfig.LinkButton)
		fmt.Fprintf(w, "IP Address:\t%v\n", bridgeConfig.IPAddress)
		fmt.Fprintf(w, "MAC Address:\t%v\n", bridgeConfig.Mac)
		fmt.Fprintf(w, "Net Mask:\t%v\n", bridgeConfig.NetMask)
		fmt.Fprintf(w, "Gateway:\t%v\n", bridgeConfig.Gateway)
		fmt.Fprintf(w, "DHCP Address:\t%v\n", bridgeConfig.Dhcp)
		fmt.Fprintf(w, "Portal Services:\t%v\n", bridgeConfig.PortalServices)
		fmt.Fprintf(w, "UTC:\t%v\n", bridgeConfig.UTC)
		fmt.Fprintf(w, "Local Time:\t%v\n", bridgeConfig.LocalTime)
		fmt.Fprintf(w, "Time Zone:\t%v\n", bridgeConfig.TimeZone)
		fmt.Fprintf(w, "Zigbee Channel:\t%v\n", bridgeConfig.ZigbeeChannel)
		fmt.Fprintf(w, "Model ID:\t%v\n", bridgeConfig.ModelID)
		fmt.Fprintf(w, "Bridge ID:\t%v\n", bridgeConfig.BridgeID)
		fmt.Fprintf(w, "Factory New:\t%v\n", bridgeConfig.FactoryNew)
		fmt.Fprintf(w, "Replaces Bridge ID:\t%v\n", bridgeConfig.ReplacesBridgeID)
		fmt.Fprintf(w, "Datastore Version:\t%v\n", bridgeConfig.DatastoreVersion)
		fmt.Fprintf(w, "Starter Kit ID:\t%v\n", bridgeConfig.StarterKitID)
		fmt.Fprintln(w, "Internet Service:\t")
		fmt.Fprintf(w, "  Internet:\t%v\n", bridgeConfig.InternetService.Internet)
		fmt.Fprintf(w, "  Remote Access:\t%v\n", bridgeConfig.InternetService.RemoteAccess)
		fmt.Fprintf(w, "  Time:\t%v\n", bridgeConfig.InternetService.Time)
		fmt.Fprintf(w, "  SW Update:\t%v\n", bridgeConfig.InternetService.SwUpdate)
		w.Flush()
	},
}

type setupParamz struct {
	bridgeIP string
	user     string
	wait     bool
	verbose  bool
}

var (
	setupParams = setupParamz{}

	setupCmd = &cobra.Command{
		Use:   "setup",
		Short: "Setup a hue bridge to work with the CLI",
		Long:  `TODO`,
		Run: func(cmd *cobra.Command, args []string) {
			bridge := huego.New(setupParams.bridgeIP, "")
			if setupParams.wait {
				reader := bufio.NewReader(os.Stdin)
				fmt.Println("Press the button on the hue bridge")
				_, _ = reader.ReadString('\n')
			}

			user, err := bridge.CreateUser(setupParams.user)
			if err != nil {
				s := fmt.Sprintf("unable to create user: %v", err)
				panic(s)
			}

			bridge = bridge.Login(user)
			_, err = bridge.GetConfig()
			if err != nil {
				s := fmt.Sprintf("unable to connect to bridge: %v", err)
				panic(s)
			}

			viper.Set("bridgeip", setupParams.bridgeIP)
			viper.Set("huecliuser", user)
			err = viper.WriteConfig()
			if err != nil {
				s := fmt.Sprintf("unable to write config: %v", err)
				panic(s)
			}
		},
	}
)

var capabilitiesCmd = &cobra.Command{
	Use:   "capabilities",
	Short: "show capabilities of the linked hue bridge",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		bridgeIP, user := getLoginFromConfig()
		bridge := huego.New(bridgeIP, user)

		capabilities, err := bridge.GetCapabilities()
		if err != nil {
			s := fmt.Sprintf("unable to connect to bridge: %v", err)
			panic(s)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(w, "Name\tAvailable")
		fmt.Fprintln(w, "----\t---------")
		fmt.Fprintf(w, "Groups\t%v\n", capabilities.Groups.Available)
		fmt.Fprintf(w, "Lights\t%v\n", capabilities.Lights.Available)
		fmt.Fprintf(w, "Resource Links\t%v\n", capabilities.Resourcelinks.Available)
		fmt.Fprintf(w, "Rules\t%v\n", capabilities.Rules.Available)
		fmt.Fprintf(w, "Scenes\t%v\n", capabilities.Scenes.Available)
		fmt.Fprintf(w, "Sensors\t%v\n", capabilities.Sensors.Available)
		fmt.Fprintf(w, "Streaming\t%v\n", capabilities.Streaming.Available)

		w.Flush()
	},
}
