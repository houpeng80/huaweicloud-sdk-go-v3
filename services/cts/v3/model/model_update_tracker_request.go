package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateTrackerRequest struct {
	Body *UpdateTrackerRequestBody `json:"body,omitempty"`
}

func (o UpdateTrackerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTrackerRequest struct{}"
	}

	return strings.Join([]string{"UpdateTrackerRequest", string(data)}, " ")
}
