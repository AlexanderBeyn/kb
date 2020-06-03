package cmd

import (
	"fmt"
	"github.com/AlexanderBeyn/kb/api"
	"github.com/AlexanderBeyn/kb/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sort"
	"strconv"
)

var showCmd = &cobra.Command{
	Use:   "show [flags] [[%%proj] %col] [/search]",
	Short: "Show tasks",

	Long: `Show tasks`,

	Example: `
# Show tasks in the default column:
kb show

# Show tasks in the "done" column containing "great task":
kb show %done "/great task"

# Show tasks from all columns containing "excellent":
kb show %* /excellent`,

	ValidArgsFunction: lib.CliSigilCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		common := lib.ParseCommonArgs(args)

		project, err := lib.PickProject("", common.Project)
		if err != nil {
			return err
		}

		var columnID int

		if common.Column == nil || *common.Column != "*" {
			column, err := lib.PickColumn("", *project, common.Column)
			if err != nil {
				return err
			}
			columnID = column.ID
		}

		tasks, err := api.GetTasks(project.ID, columnID, common.Search)
		if err != nil {
			return err
		}

		sort.Slice(tasks, func(i, j int) bool {
			return tasks[i].ColumnID < tasks[j].ColumnID
		})

		key := fmt.Sprintf("cache.columns.%d", project.ID)
		cache := viper.Get(key).(map[string]interface{})
		var oldColumn int
		var columnName string
		for _, t := range tasks {
			if oldColumn != t.ColumnID {
				cacheColumn, _ := cache[strconv.Itoa(t.ColumnID)].(map[string]interface{})
				columnName, _ = cacheColumn["title"].(string)
				fmt.Printf("%s / %s\n", project.Name, columnName)
				oldColumn = t.ColumnID
			}
			fmt.Printf("- %s\n", t.Title)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
