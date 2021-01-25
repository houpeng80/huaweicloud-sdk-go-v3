package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateScriptRequest struct {
	Body *ScriptInfo `json:"body,omitempty"`
}

func (o CreateScriptRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateScriptRequest struct{}"
	}

	return strings.Join([]string{"CreateScriptRequest", string(data)}, " ")
}
