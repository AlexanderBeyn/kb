package lib

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"kb/api"
	"kb/types"
)

func PickProject(label string, projectFilter *string) (*types.Project, error) {
	project := &types.Project{
		ID:   viper.GetInt("selected.project.id"),
		Name: viper.GetString("selected.project.name"),
	}

	if projectFilter == nil && project.ID != 0 {
		return project, nil
	}

	projects, err := api.GetProjects(projectFilter)
	if err != nil {
		return nil, err
	}

	project, err = SelectProject(label, projects, true)
	if project == nil && err == nil {
		err = fmt.Errorf("no matching projects found")
	}
	return project, err
}

func SelectProject(label string, projects []*types.Project, autoSelect bool) (*types.Project, error) {
	if label == "" {
		label = "Select project"
	}

	if len(projects) == 0 {
		return nil, nil
	}

	if len(projects) == 1 && autoSelect {
		return projects[0], nil
	}

	var idx int
	var options []string
	for _, project := range projects {
		options = append(options, project.Name)
	}
	err := survey.AskOne(&survey.Select{
		Message: label,
		Options: options,
	}, &idx)

	if err != nil {
		return nil, err
	}

	return projects[idx], nil
}
