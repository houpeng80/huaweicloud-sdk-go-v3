package model

import (
	"encoding/json"

	"strings"
)

type Resource struct {
	ExtraInfo *ResourceExtraInfo `json:"extra_info,omitempty"`
	// 待备份资源id

	Id string `json:"id"`
	// 待备份资源名称，长度限制：0-255

	Name string `json:"name"`
	// [待备份资源的类型: OS::Nova::Server, OS::Cinder::Volume, OS::Ironic::BareMetalServer, OS::Native::Server, OS::Sfs::Turbo](tag:hws,hws_hk,fcs_vm,ctc) [待备份资源的类型: OS::Nova::Server,  OS::Cinder::Volume](tag:dt,ocb,tlf,sbc)

	Type string `json:"type"`
}

func (o Resource) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
