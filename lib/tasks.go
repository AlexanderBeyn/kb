package lib

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"kb/api"
	"kb/types"
)

func PickTask(label string, project types.Project, column types.Column, filter *string) (*types.Task, error) {
	tasks, err := api.GetTasks(project.ID, column.ID, filter)
	if err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, fmt.Errorf("no matching tasks found in column %v", column.Title)
	}

	return SelectTask(label, tasks, false)
}

func SelectTask(label string, tasks []*types.Task, autoSelect bool) (*types.Task, error) {
	if label == "" {
		label = "Select task"
	}

	if len(tasks) == 0 {
		return nil, nil
	}

	if len(tasks) == 1 && autoSelect {
		return tasks[0], nil
	}

	var idx int
	var options []string
	for _, task := range tasks {
		options = append(options, task.Title)
	}
	err := survey.AskOne(&survey.Select{
		Message: label,
		Options: options,
	}, &idx)

	if err != nil {
		return nil, err
	}

	return tasks[idx], nil
}
