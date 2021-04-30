package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type SetBalancerSwitchResponse struct {
	// 任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetBalancerSwitchResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SetBalancerSwitchResponse struct{}"
	}

	return strings.Join([]string{"SetBalancerSwitchResponse", string(data)}, " ")
}
