package lib

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"kb/api"
	"kb/types"
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
		out = viper.GetStringSlice("cache.columns")
		if len(out) > 0 {
			break
		}

		projectId = viper.GetInt("selected.project.id")
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
