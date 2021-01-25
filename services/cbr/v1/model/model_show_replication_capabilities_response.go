package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowReplicationCapabilitiesResponse struct {
	// 支持复制的区域列表
	Regions        *[]ProtectableReplicationCapabilitiesRespRegion `json:"regions,omitempty"`
	HttpStatusCode int                                             `json:"-"`
}

func (o ShowReplicationCapabilitiesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowReplicationCapabilitiesResponse struct{}"
	}

	return strings.Join([]string{"ShowReplicationCapabilitiesResponse", string(data)}, " ")
}
