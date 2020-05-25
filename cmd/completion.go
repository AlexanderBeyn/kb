package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var completionCmd = &cobra.Command{
	Use:       "completion [flags] <bash|zsh|fish|powershell>",
	Short:     "Generate shell completion code for the specified shell",
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
	RunE: func(cmd *cobra.Command, args []string) error {
		switch args[0] {
		case "bash":
			return rootCmd.GenBashCompletion(os.Stdout)
		case "zsh":
			return rootCmd.GenZshCompletion(os.Stdout)
		case "fish":
			return rootCmd.GenFishCompletion(os.Stdout, true)
		case "powershell":
			return rootCmd.GenPowerShellCompletion(os.Stdout)
		default:
			return fmt.Errorf("invalid shell '%s'", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
