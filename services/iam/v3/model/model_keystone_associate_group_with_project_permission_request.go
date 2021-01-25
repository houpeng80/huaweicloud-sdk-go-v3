package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneAssociateGroupWithProjectPermissionRequest struct {
	ProjectId string `json:"project_id"`
	GroupId   string `json:"group_id"`
	RoleId    string `json:"role_id"`
}

func (o KeystoneAssociateGroupWithProjectPermissionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneAssociateGroupWithProjectPermissionRequest struct{}"
	}

	return strings.Join([]string{"KeystoneAssociateGroupWithProjectPermissionRequest", string(data)}, " ")
}
