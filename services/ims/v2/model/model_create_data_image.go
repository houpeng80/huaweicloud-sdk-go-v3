package model

import (
	"encoding/json"

	"strings"
)

// 数据盘信息
type CreateDataImage struct {
	// 数据盘镜像名称。
	Name string `json:"name"`
	// 数据盘ID。
	VolumeId string `json:"volume_id"`
	// 数据盘描述。
	Description *string `json:"description,omitempty"`
	// 数据盘镜像标签。
	Tags *[]string `json:"tags,omitempty"`
}

func (o CreateDataImage) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateDataImage struct{}"
	}

	return strings.Join([]string{"CreateDataImage", string(data)}, " ")
}
