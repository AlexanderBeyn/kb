package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kb/api"
	"kb/lib"
)

var showCmd = &cobra.Command{
	Use:   "show [flags] [[%%proj] %col] [/search]",
	Short: "Show tasks",

	Long: `Show tasks`,

	Example: `
# Show tasks in the default column:
kb show

# Show tasks in the "done" column containing "great task":
kb show %done "/great task"`,

	ValidArgsFunction: lib.CliSigilCompletions,
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

		tasks, err := api.GetTasks(project.ID, column.ID, common.Search)
		if err != nil {
			return err
		}

		fmt.Printf("%s / %s:\n", project.Name, column.Title)
		for _, t := range tasks {
			fmt.Printf("- %s\n", t.Title)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
