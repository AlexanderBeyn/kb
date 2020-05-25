package api

import (
	"fmt"
	"kb/types"
)

func GetTasks(project int, column int, filter *string) ([]*types.Task, error) {
	var tasks []*types.Task

	if filter == nil {
		filter = new(string)
	}

	err := RPC.CallFor(&tasks, "searchTasks", &types.SearchTasksParams{
		ProjectID: project,
		Query:     fmt.Sprintf("status:open column:%d %s", column, *filter),
	})
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func CreateTask(task types.CreateTaskParams) (int, error) {
	var out int
	err := RPC.CallFor(&out, "createTask", task)
	return out, err
}

func MoveTaskPosition(params types.MoveTaskPositionParams) (bool, error) {
	var out bool
	err := RPC.CallFor(&out, "moveTaskPosition", params)
	return out, err
}
