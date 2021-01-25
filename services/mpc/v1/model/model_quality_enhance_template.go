package model

import (
	"encoding/json"

	"strings"
)

type QualityEnhanceTemplate struct {
	// 模板名称。
	TemplateName string `json:"template_name"`
	// 模板描述，查询预置模板时才会返回。
	TemplateDescription *string              `json:"template_description,omitempty"`
	Video               *QualityEnhanceVideo `json:"video,omitempty"`
}

func (o QualityEnhanceTemplate) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QualityEnhanceTemplate struct{}"
	}

	return strings.Join([]string{"QualityEnhanceTemplate", string(data)}, " ")
}
