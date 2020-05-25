package api

import (
	"github.com/spf13/viper"
	"kb/types"
	"strings"
	"time"
)

func GetColumns(project int, filter *string) ([]*types.Column, error) {
	var columns []*types.Column

	err := RPC.CallFor(&columns, "getColumns", []int{project})
	if err != nil {
		return nil, err
	}

	var cache []string
	for _, column := range columns {
		cache = append(cache, column.Title)
	}
	viper.Set("cache.columns", cache)
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
