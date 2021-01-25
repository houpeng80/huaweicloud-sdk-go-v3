package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneUpdateIdentityProviderRequest struct {
	Id   string                                     `json:"id"`
	Body *KeystoneUpdateIdentityProviderRequestBody `json:"body,omitempty"`
}

func (o KeystoneUpdateIdentityProviderRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneUpdateIdentityProviderRequest struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateIdentityProviderRequest", string(data)}, " ")
}
