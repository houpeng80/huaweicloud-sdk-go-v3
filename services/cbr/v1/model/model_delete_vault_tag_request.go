package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteVaultTagRequest struct {
	Key string `json:"key"`

	VaultId string `json:"vault_id"`
}

func (o DeleteVaultTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteVaultTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteVaultTagRequest", string(data)}, " ")
}
