package commands

import (
	"fmt"
	"github.com/amimof/huego"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string
	homeDir string

	rootCmd = &cobra.Command{
		Use:   "hue-cli",
		Short: "hue-cli is a CLI tool for philips hue ecosystem",
		Long:  `TODO`,
		Run: func(cmd *cobra.Command, args []string) {
			// display help
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			s := fmt.Sprintf("error determining homedir: %v", err)
			panic(s)
		}

		var (
			configName = ".hue-cli"
			configType = "yml"
		)
		viper.AddConfigPath(home)
		viper.SetConfigName(configName)
		viper.SetConfigType(configType)
		configPath := filepath.Join(home, configName+"."+configType)

		_, err = os.Stat(configPath)
		if os.IsNotExist(err) {
			fmt.Printf("Creating config file at %v\n", configPath)
			if _, err := os.Create(configPath); err != nil {
				s := fmt.Sprintf("error creating config file at path %q: %v", configPath, err)
				panic(s)
			}
		}
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func getLoginFromConfig() (string, string) {
	bridgeIP := viper.GetString("bridgeip")
	user := viper.GetString("huecliuser")

	if bridgeIP == "" {
		fmt.Println("A bridgeIP is not configured. Use discover and setup commands to set up a bridge")
		os.Exit(1)
	}

	if user == "" {
		fmt.Println("A user is not configured. Use discover and setup commands to set up a bridge")
		os.Exit(1)
	}

	return bridgeIP, user
}

func getBridgeFromConfig() *huego.Bridge {
	bridgeIP, user := getLoginFromConfig()
	return huego.New(bridgeIP, user)
}
