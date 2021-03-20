package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListApisNotBoundWithSignatureKeyV2Request struct {
	InstanceId string `json:"instance_id"`

	SignId string `json:"sign_id"`

	EnvId *string `json:"env_id,omitempty"`

	ApiId *string `json:"api_id,omitempty"`

	ApiName *string `json:"api_name,omitempty"`

	GroupId *string `json:"group_id,omitempty"`

	Offset *int64 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListApisNotBoundWithSignatureKeyV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApisNotBoundWithSignatureKeyV2Request struct{}"
	}

	return strings.Join([]string{"ListApisNotBoundWithSignatureKeyV2Request", string(data)}, " ")
}
