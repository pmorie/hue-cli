package cmd

import (
	"fmt"

	"github.com/amimof/huego"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "setup a new user with the hue bridge",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		bridge, err := huego.Discover()
		if err != nil {
			s := fmt.Sprintf("unable to discover bridge: %v", err)
			panic(s)
		}

		fmt.Printf("found bridge: %+v", bridge)

		user, err := bridge.CreateUser("my awesome hue app") // Link button needs to be pressed
		if err != nil {
			s := fmt.Sprintf("unable to create user: %v", err)
			panic(s)
		}

		// no error checking? what does this do?
		bridge = bridge.Login(user)

		// TODO: save off the username (what about auth??) and ip address of the bridge

		// re-load the username from the state file and ensure it saved correctly
	},
}
