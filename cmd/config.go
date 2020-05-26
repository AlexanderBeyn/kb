package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/AlexanderBeyn/kb/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sort"
	"strings"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage kb configuration",
}

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all configuration settings",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		keys := viper.AllKeys()
		sort.Strings(keys)

		for _, key := range keys {
			if strings.HasPrefix(key, "cache.") {
				continue
			}
			value := viper.Get(key)
			fmt.Printf("%s: %v\n", key, value)
		}

		return nil
	},
}

var configGetCmd = &cobra.Command{
	Use:   "get [flags] <setting>",
	Short: "Display a configuration setting",
	Args:  cobra.ExactValidArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var keys []string
		for _, key := range viper.AllKeys() {
			if strings.HasPrefix(key, "cache.") || strings.HasPrefix(key, "selected.") {
				continue
			}
			keys = append(keys, key)
		}
		sort.Strings(keys)
		return keys, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%s: %v\n", args[0], viper.Get(args[0]))
		return nil
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set [flags] <key> <value>",
	Short: "Set a configuration setting to the specified value",
	Args: func(cmd *cobra.Command, args []string) error {
		switch len(args) {
		case 2:
			for _, key := range viper.AllKeys() {
				if key == args[0] {
					return nil
				}
			}
			return fmt.Errorf("unknown setting %q", args[0])

		case 1:
			return fmt.Errorf("missing value for %q", args[0])

		case 0:
			return fmt.Errorf("not enough arguments")

		default:
			return fmt.Errorf("too many arguments")
		}
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) == 1 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		var keys []string
		for _, key := range viper.AllKeys() {
			if strings.HasPrefix(key, "cache.") || strings.HasPrefix(key, "selected.") {
				continue
			}
			keys = append(keys, key)
		}
		sort.Strings(keys)
		return keys, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		viper.Set(args[0], args[1])
		return viper.WriteConfig()
	},
}

var configDefaultsCmd = &cobra.Command{
	Use:   "defaults [flags] [%%proj] [%col]",
	Short: "Set default project and column",

	Example: `
# Set the default project to "mine" and default columnt to "backlog":
kb config defaults %%mine %backlog`,

	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		common := lib.ParseCommonArgs(args)

		project, err := lib.PickProject("", common.Project)
		if err != nil {
			return err
		}

		column, err := lib.PickColumn("", *project, common.Column)
		if err != nil {
			return err
		}

		viper.Set("selected.project.id", project.ID)
		viper.Set("selected.project.name", project.Name)
		viper.Set("selected.column.id", column.ID)
		viper.Set("selected.column.title", column.Title)

		return viper.WriteConfig()
	},

	ValidArgsFunction: lib.CliSigilCompletions,
}

var configPromptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Prompt for basic configuration options",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		qs := []*survey.Question{
			{
				Name: "url",
				Prompt: &survey.Input{
					Message: "Enter Kanboard URL (including /jsonrpc.php):",
					Default: viper.GetString("server.url"),
				},
			},
			{
				Name: "user",
				Prompt: &survey.Input{
					Message: "Enter username:",
					Default: viper.GetString("server.user"),
				},
			},
			{
				Name:   "password",
				Prompt: &survey.Password{Message: "Enter password or API key (leave blank to not change):"},
			},
		}

		answers := struct {
			URL      string
			User     string
			Password string
		}{}

		err := survey.Ask(qs, &answers)
		if err != nil {
			return err
		}

		viper.Set("server.url", answers.URL)
		viper.Set("server.user", answers.User)
		if answers.Password != "" {
			viper.Set("server.password", answers.Password)
		}

		return viper.WriteConfig()
	},
}

func init() {
	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configDefaultsCmd)
	configCmd.AddCommand(configPromptCmd)

	rootCmd.AddCommand(configCmd)
}
