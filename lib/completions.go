package lib

import (
	"fmt"
	"github.com/AlexanderBeyn/kb/api"
	"github.com/AlexanderBeyn/kb/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

func CliSigilCompletions(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var out []string
	var projectId int
	var projects []*types.Project
	var columns []*types.Column
	var err error

	sigil, toComplete := StripSigil(toComplete)
	switch sigil {
	case ProjectSigil:
		projects, err = api.GetProjects(&toComplete)
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		for _, project := range projects {
			out = append(out, project.Name)
		}
	case ColumnSigil, FromColumnSigil:
		projectId = viper.GetInt("selected.project.id")

		cache := viper.GetStringMap(fmt.Sprintf("cache.columns.%d", projectId))
		if len(cache) > 0 {
			for _, column := range cache {
				title, ok := column.(map[string]interface{})["title"].(string)
				if !ok {
					continue
				}
				out = append(out, title)
			}
			break
		}

		columns, err = api.GetColumns(projectId, &toComplete)
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		for _, column := range columns {
			out = append(out, column.Title)
		}
	}

	for i, o := range out {
		out[i] = strings.ReplaceAll(
			strings.ToLower(sigil+o), " ", `\\\ `,
		)
	}

	return out, cobra.ShellCompDirectiveNoFileComp
}
