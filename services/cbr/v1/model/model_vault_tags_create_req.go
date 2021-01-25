package model

import (
	"encoding/json"

	"strings"
)

//
type VaultTagsCreateReq struct {
	Tag *Tag `json:"tag,omitempty"`
}

func (o VaultTagsCreateReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VaultTagsCreateReq struct{}"
	}

	return strings.Join([]string{"VaultTagsCreateReq", string(data)}, " ")
}
