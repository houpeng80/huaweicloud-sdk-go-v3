package model

import (
	"encoding/json"

	"strings"
)

//
type IdentityToken struct {
	// token的ID。与请求头中的X-Auth-Token填写其一即可，若都填写，优先校验X-Auth-Token。
	Id *string `json:"id,omitempty"`
	// AK/SK和securitytoken的有效期，时间单位为秒。取值范围：15min ~ 24h ，默认为15min。
	DurationSeconds *int32 `json:"duration_seconds,omitempty"`
}

func (o IdentityToken) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "IdentityToken struct{}"
	}

	return strings.Join([]string{"IdentityToken", string(data)}, " ")
}
