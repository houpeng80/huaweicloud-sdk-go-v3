package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowLoadbalancersStatusResponse struct {
	Statuses       *StatusResp `json:"statuses,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowLoadbalancersStatusResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowLoadbalancersStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowLoadbalancersStatusResponse", string(data)}, " ")
}
