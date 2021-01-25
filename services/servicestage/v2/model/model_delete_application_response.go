package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteApplicationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteApplicationResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteApplicationResponse struct{}"
	}

	return strings.Join([]string{"DeleteApplicationResponse", string(data)}, " ")
}
