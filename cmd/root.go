package cmd

import (
	"fmt"
	"github.com/AlexanderBeyn/kb/api"
	"github.com/AlexanderBeyn/kb/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use:               "kb",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	err := lib.ReadConfig()
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}

	url := viper.GetString("server.url")
	user := viper.GetString("server.user")
	password := viper.GetString("server.password")

	api.InitRPC(url, user, password)
}
