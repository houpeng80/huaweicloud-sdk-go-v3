package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateAlarmTemplateResponse struct {
	// 自定义告警模板创建成功返回的ID，如：at1603252280799wLRyGLxnz。
	TemplateId     *string `json:"template_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAlarmTemplateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAlarmTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateAlarmTemplateResponse", string(data)}, " ")
}
