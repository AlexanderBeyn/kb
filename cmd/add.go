package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"kb/api"
	"kb/lib"
	"kb/types"
	"strings"
)

var addCmd = &cobra.Command{
	Use:   "add [flags] [[%%proj] %col] [title] [+]",
	Short: "Add a new task",

	Long: `Add a new task

You will be prompted for a title if you don't specify it on the command line.
If the title on the command line ends with a single '+' by itself, you will
be prompted for a description for the new task.`,

	Example: `
# Add a new task to the default column:
kb add This is my new task

# Add a new task to a column "backlog", prompting for a description:
kb add %backlog This is a new backlog task +

# Add a new task, prompting for a column and title:
kb add %`,

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

		title := strings.Join(common.Args, " ")
		var description string
		needDescription := common.More
		if len(title) == 0 {
			err := survey.AskOne(&survey.Input{
				Message: "Enter a title:",
			}, &title)
			if err != nil {
				return err
			}

			if len(title) == 0 {
				fmt.Println("Cancelling!")
				return nil
			}

			needDescription = true
		}

		if needDescription {
			err := survey.AskOne(&survey.Multiline{
				Message: "Enter a description",
			}, &description)
			if err != nil {
				return err
			}
		}

		fmt.Printf("Creating new task in project %s / %s:\n", project.Name, column.Title)
		fmt.Printf("Title: %s\n", title)
		fmt.Println("-- Description --")
		if description == "" {
			fmt.Println("<empty>")
		} else {
			fmt.Println(description)
		}

		_, err = api.CreateTask(types.CreateTaskParams{
			ProjectID:   project.ID,
			ColumnID:    column.ID,
			Title:       title,
			Description: description,
		})
		if err != nil {
			return err
		}

		return nil
	},

	ValidArgsFunction: lib.CliSigilCompletions,
}

func init() {
	rootCmd.AddCommand(addCmd)
}
