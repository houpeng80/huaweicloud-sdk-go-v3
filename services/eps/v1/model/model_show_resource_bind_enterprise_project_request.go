package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowResourceBindEnterpriseProjectRequest struct {
	EnterpriseProjectId string         `json:"enterprise_project_id"`
	Body                *ResqEpResouce `json:"body,omitempty"`
}

func (o ShowResourceBindEnterpriseProjectRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowResourceBindEnterpriseProjectRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceBindEnterpriseProjectRequest", string(data)}, " ")
}
