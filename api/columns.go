package api

import (
	"fmt"
	"github.com/AlexanderBeyn/kb/types"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

func GetColumns(project int, filter *string) ([]*types.Column, error) {
	var columns []*types.Column

	err := RPC.CallFor(&columns, "getColumns", []int{project})
	if err != nil {
		return nil, err
	}

	cache := make(map[string]map[string]interface{})
	for _, column := range columns {
		cache[strconv.Itoa(column.ID)] = map[string]interface{}{
			"Title":    column.Title,
			"Position": column.Position,
		}
	}
	viper.Set(fmt.Sprintf("cache.columns.%d", project), cache)
	viper.Set("cache.columns_updated", time.Now())
	err = viper.WriteConfig()
	if err != nil {
		return nil, err
	}

	if filter != nil {
		filteredColumns := make([]*types.Column, 0, len(columns))

		for _, p := range columns {
			if len(*filter) <= len(p.Title) && strings.EqualFold(*filter, p.Title[:len(*filter)]) {
				filteredColumns = append(filteredColumns, p)
			}
		}

		columns = filteredColumns
	}

	return columns, nil
}
