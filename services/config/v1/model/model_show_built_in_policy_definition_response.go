package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBuiltInPolicyDefinitionResponse Response Object
type ShowBuiltInPolicyDefinitionResponse struct {

	// 策略id
	Id *string `json:"id,omitempty"`

	// 策略名字
	Name *string `json:"name,omitempty"`

	// 策略类型
	PolicyType *string `json:"policy_type,omitempty"`

	// 策略描述
	Description *string `json:"description,omitempty"`

	// 策略语法类型
	PolicyRuleType *string `json:"policy_rule_type,omitempty"`

	// 策略规则
	PolicyRule *interface{} `json:"policy_rule,omitempty"`

	// 触发器类型，可选值：resource, period
	TriggerType *string `json:"trigger_type,omitempty"`

	// 关键词列表
	Keywords *[]string `json:"keywords,omitempty"`

	// 默认资源类型列表
	DefaultResourceTypes *[]PolicyDefinitionDefaultResourceTypes `json:"default_resource_types,omitempty"`

	// 策略参数
	Parameters     map[string]PolicyParameterDefinition `json:"parameters,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ShowBuiltInPolicyDefinitionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBuiltInPolicyDefinitionResponse struct{}"
	}

	return strings.Join([]string{"ShowBuiltInPolicyDefinitionResponse", string(data)}, " ")
}
