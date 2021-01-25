package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CopyImageInRegionRequest struct {
	ImageId string                        `json:"image_id"`
	Body    *CopyImageInRegionRequestBody `json:"body,omitempty"`
}

func (o CopyImageInRegionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CopyImageInRegionRequest struct{}"
	}

	return strings.Join([]string{"CopyImageInRegionRequest", string(data)}, " ")
}
