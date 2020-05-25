package api

import (
	"kb/types"
	"strings"
)

func GetProjects(filter *string) ([]*types.Project, error) {
	var projects []*types.Project

	err := RPC.CallFor(&projects, "getAllProjects")
	if err != nil {
		return nil, err
	}

	if filter != nil {
		filteredProjects := make([]*types.Project, 0, len(projects))

		for _, p := range projects {
			if len(*filter) <= len(p.Name) && strings.EqualFold(*filter, p.Name[:len(*filter)]) {
				filteredProjects = append(filteredProjects, p)
			}
		}

		projects = filteredProjects
	}

	return projects, nil
}
