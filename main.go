package main

import (
	"fmt"

	"github.com/amimof/huego"
)

func main() {
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

	bridge = bridge.Login(user)

}
