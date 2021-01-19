package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/amimof/huego"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	setupCmd.PersistentFlags().StringVar(&setupParams.bridgeIP, "bridgeIP", "", "IP of the bridge to setup")
	setupCmd.PersistentFlags().StringVar(&setupParams.user, "user", "hue-cli", "user to setup")
	setupCmd.PersistentFlags().BoolVar(&setupParams.wait, "wait", true, "user to setup")
	setupCmd.PersistentFlags().BoolVar(&setupParams.verbose, "verbose", false, "verbosity")

	bridgesCmd.AddCommand(setupCmd)
	bridgesCmd.AddCommand(discoverCmd)
	bridgesCmd.AddCommand(statusCmd)

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
	Short: "Discover hue bridges on your network",
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

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "show the status of the configured bridge",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		bridgeIP := viper.GetString("bridgeip")
		user := viper.GetString("huecliuser")

		if bridgeIP == "" {
			fmt.Println("A bridgeIP is not configured. Use discover and setup commands to set up a bridge")
			os.Exit(1)
		}

		if user == "" {
			fmt.Println("A bridgeIP is not configured. Use discover and setup commands to set up a bridge")
			os.Exit(1)
		}

		fmt.Printf("Bridge IP: %v\n", bridgeIP)
		fmt.Printf("User: %v\n", user)

		bridge := huego.New(bridgeIP, user)
		bridgeConfig, err := bridge.GetConfig()
		if err != nil {
			s := fmt.Sprintf("unable to connect to bridge: %v", err)
			panic(s)
		}

		fmt.Printf("Bridge Configuration:\n\n")
		fmt.Printf("Name: %v\n", bridgeConfig.Name)
		fmt.Println("Software Update Configuration:")
		fmt.Printf("  Check For Updates: %v\n", bridgeConfig.SwUpdate2.CheckForUpdate)
		fmt.Printf("  State: %v\n", bridgeConfig.SwUpdate2.State)
		fmt.Printf("  AutoInstall:\n")
		fmt.Printf("    Enabled: %v\n", bridgeConfig.SwUpdate2.AutoInstall.On)
		fmt.Printf("    Last Update Time: %v\n", bridgeConfig.SwUpdate2.AutoInstall.UpdateTime)
		fmt.Printf("  Last Change: %v\n", bridgeConfig.SwUpdate2.LastChange)
		fmt.Printf("  Last Install: %v\n", bridgeConfig.SwUpdate2.LastInstall)
		fmt.Println("Allowed Clients:")
		for k, _ := range bridgeConfig.WhitelistMap {
			fmt.Printf("  %v\n", k)
		}
		fmt.Printf("")
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
		Short: "setup a new user with the hue bridge",
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
