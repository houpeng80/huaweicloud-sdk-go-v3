package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateAnimatedGraphicsTaskResponse struct {
	// 任务ID
	TaskId *string `json:"task_id,omitempty"`
	// 任务状态
	Status *string `json:"status,omitempty"`
	// 任务创建时间
	CreateTime *string     `json:"create_time,omitempty"`
	Output     *ObsObjInfo `json:"output,omitempty"`
	// 解析文件名称
	OutputFileName *string `json:"output_file_name,omitempty"`
	// 任务描述，如当任务异常时，此字段为异常的具体信息
	Description    *string   `json:"description,omitempty"`
	Metadata       *MetaData `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateAnimatedGraphicsTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAnimatedGraphicsTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateAnimatedGraphicsTaskResponse", string(data)}, " ")
}
