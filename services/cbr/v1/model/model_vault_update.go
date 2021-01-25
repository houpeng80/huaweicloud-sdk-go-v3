package model

import (
	"encoding/json"

	"strings"
)

type VaultUpdate struct {
	Billing *BillingUpdate `json:"billing,omitempty"`
	// 存储库名称
	Name *string `json:"name,omitempty"`
	// 是否支持自动挂载
	AutoBind  *bool           `json:"auto_bind,omitempty"`
	BindRules *VaultBindRules `json:"bind_rules,omitempty"`
	// 是否自动扩容。按需存储库支持自动扩容，包周期存储库不支持扩容。
	AutoExpand *bool `json:"auto_expand,omitempty"`
}

func (o VaultUpdate) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VaultUpdate struct{}"
	}

	return strings.Join([]string{"VaultUpdate", string(data)}, " ")
}
