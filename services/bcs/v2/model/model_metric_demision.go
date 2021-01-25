package model

import (
	"encoding/json"

	"strings"
)

// 指标描述信息
type MetricDemision struct {
	// 命名空间
	Namespace *string `json:"namespace,omitempty"`
	// 指标名称
	MetricName *string `json:"metricName,omitempty"`
	// 维度列表
	Dimensions *[]Dimension `json:"dimensions,omitempty"`
}

func (o MetricDemision) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MetricDemision struct{}"
	}

	return strings.Join([]string{"MetricDemision", string(data)}, " ")
}
