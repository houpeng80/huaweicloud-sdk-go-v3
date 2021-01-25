package model

import (
	"encoding/json"

	"strings"
)

type ProtectableResult struct {
	// 不支持备份的错误码
	Code *string `json:"code,omitempty"`
	// 不支持备份的原因
	Reason *string `json:"reason,omitempty"`
	// 是否可备份
	Result string    `json:"result"`
	Vault  *VaultGet `json:"vault,omitempty"`
}

func (o ProtectableResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ProtectableResult struct{}"
	}

	return strings.Join([]string{"ProtectableResult", string(data)}, " ")
}
