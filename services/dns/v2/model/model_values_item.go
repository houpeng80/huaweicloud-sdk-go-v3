package model

import (
	"encoding/json"

	"strings"
)

// 版本对象。
type ValuesItem struct {
	// 所有版本列表。
	Values *[]ListApiVersionsItem `json:"values,omitempty"`
}

func (o ValuesItem) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ValuesItem struct{}"
	}

	return strings.Join([]string{"ValuesItem", string(data)}, " ")
}
