package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"kb/api"
	"kb/lib"
	"kb/types"
)

var moveCmd = &cobra.Command{
	Use:   "move [flags] [%%proj] [^%source_col] [%target_col] [/search]",
	Short: "Move a task to a new column",

	Long: `Move a task to a new column

Tasks will be moved from the default column, if a source column is not
specified. If a target column is not specified, you will be prompted to
select one.`,

	Example: `
# Find a task matching "awesome" in the default column and move it to "done":
kb move %done /awesome

# Display all tasks in "backlog" and select one to move to "today":
kb move ^%backlog %today`,

	RunE: func(cmd *cobra.Command, args []string) error {
		common := lib.ParseCommonArgs(args)

		project, err := lib.PickProject("", common.Project)
		if err != nil {
			return err
		}

		var fromColumn = &types.Column{
			ID:    viper.GetInt("selected.column.id"),
			Title: viper.GetString("selected.column.title"),
		}

		if common.FromColumn != nil {
			fromColumn, err = lib.PickColumn("Select source column", *project, common.FromColumn)
			if err != nil {
				return err
			}
		}

		task, err := lib.PickTask("", *project, *fromColumn, common.Search)
		if err != nil {
			return err
		}

		if common.Column == nil {
			common.Column = new(string)
		}
		column, err := lib.PickColumn("Select destination column", *project, common.Column)
		if err != nil {
			return err
		}

		fmt.Printf("Moving task in project %s from %s to %s:\n", project.Name, fromColumn.Title, column.Title)
		fmt.Printf("Title: %s\n", task.Title)

		success, err := api.MoveTaskPosition(types.MoveTaskPositionParams{
			TaskID:     task.ID,
			ProjectID:  column.ProjectID,
			ColumnID:   column.ID,
			Position:   1,
			SwimlaneID: task.SwimlaneID,
		})

		if !success && err == nil {
			err = fmt.Errorf("unable to move task")
		}

		return err
	},
}

func init() {
	rootCmd.AddCommand(moveCmd)
}
