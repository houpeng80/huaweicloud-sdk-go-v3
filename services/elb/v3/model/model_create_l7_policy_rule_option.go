package model

import (
	"encoding/json"

	"strings"
)

// 转发规则匹配策略
type CreateL7PolicyRuleOption struct {
	// 转发规则的管理状态；取值范围： true/false。该字段为预留字段，暂未启用。默认为true。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 转发规则的匹配类型。取值范围：HOST_NAME：匹配请求中的域名；PATH：匹配请求中的路径；同一个转发策略下转发规则的type不能重复。
	Type string `json:"type"`
	// 转发匹配方式：type为HOST_NAME时，取值范围：EQUAL_TO：精确匹配；type为PATH时，取值范围：REGEX：正则匹配；STARTS_WITH：前缀匹配；EQUAL_TO：精确匹配。
	CompareType string `json:"compare_type"`
	// 是否反向匹配；取值范围：true/false。默认值：false；该字段为预留字段，暂未启用。
	Invert *bool `json:"invert,omitempty"`
	// 匹配内容的键值。默认为null。该字段为预留字段，暂未启用。
	Key *string `json:"key,omitempty"`
	// 匹配内容的值。不能包含空格。当type为HOST_NAME时，取值范围：String (100)，字符串只能包含英文字母、数字、“-”或“.”，且必须以字母或数字开头。当type为PATH时，取值范围：String (128)。当转发规则的compare_type为STARTS_WITH、EQUAL_TO时，字符串只能包含英文字母、数字、_~';@^-%#&$.*+?,=!:| /()[]{}，且必须以\"/\"开头。
	Value string `json:"value"`
}

func (o CreateL7PolicyRuleOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateL7PolicyRuleOption struct{}"
	}

	return strings.Join([]string{"CreateL7PolicyRuleOption", string(data)}, " ")
}
