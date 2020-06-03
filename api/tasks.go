package api

import (
	"fmt"
	"github.com/AlexanderBeyn/kb/types"
)

func GetTasks(project int, column int, filter *string) ([]*types.Task, error) {
	var tasks []*types.Task

	query := "status:open"
	if column != 0 {
		query = fmt.Sprintf("%s column:%d", query, column)
	}

	if filter != nil {
		query = fmt.Sprintf("%s %s", query, *filter)
	}

	err := RPC.CallFor(&tasks, "searchTasks", &types.SearchTasksParams{
		ProjectID: project,
		Query:     query,
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
