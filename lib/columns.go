package lib

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"kb/api"
	"kb/types"
)

func PickColumn(label string, project types.Project, columnFilter *string) (*types.Column, error) {
	column := &types.Column{
		ID:        viper.GetInt("selected.column.id"),
		Title:     viper.GetString("selected.column.title"),
		ProjectID: project.ID,
	}

	if columnFilter == nil && column.ID != 0 {
		return column, nil
	}

	columns, err := api.GetColumns(project.ID, columnFilter)
	if err != nil {
		return nil, err
	}

	column, err = SelectColumn(label, columns, true)
	if column == nil && err == nil {
		err = fmt.Errorf("no matching columns found in project %s", project.Name)
	}
	return column, err
}

func SelectColumn(label string, columns []*types.Column, autoSelect bool) (*types.Column, error) {
	if label == "" {
		label = "Select column"
	}

	if len(columns) == 0 {
		return nil, nil
	}

	if len(columns) == 1 && autoSelect {
		return columns[0], nil
	}

	var idx int
	var options []string
	for _, column := range columns {
		options = append(options, column.Title)
	}
	err := survey.AskOne(&survey.Select{
		Message: label,
		Options: options,
	}, &idx)

	if err != nil {
		return nil, err
	}

	return columns[idx], nil
}
