package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/amimof/huego"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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

			// no error checking? what does this do?
			bridge = bridge.Login(user)
			bridgeConfig, err := bridge.GetConfig()
			if err != nil {
				s := fmt.Sprintf("unable to connect to bridge: %v", err)
				panic(s)
			}
			fmt.Printf("Bridge Configuration:\n%+v", bridgeConfig)
			// TODO: save off the username (what about auth??) and ip address of the bridge

			viper.Set("bridgeIP", setupParams.bridgeIP)
			viper.Set("user", setupParams.user)
			err = viper.WriteConfig()
			if err != nil {
				s := fmt.Sprintf("unable to write config: %v", err)
				panic(s)
			}
		},
	}
)

func init() {
	setupCmd.PersistentFlags().StringVar(&setupParams.bridgeIP, "bridgeIP", "", "IP of the bridge to setup")
	setupCmd.PersistentFlags().StringVar(&setupParams.user, "user", "hue-cli", "user to setup")
	setupCmd.PersistentFlags().BoolVar(&setupParams.wait, "wait", true, "user to setup")
	setupCmd.PersistentFlags().BoolVar(&setupParams.verbose, "verbose", false, "verbosity")

	rootCmd.AddCommand(setupCmd)
}
