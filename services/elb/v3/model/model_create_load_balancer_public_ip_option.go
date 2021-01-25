package model

import (
	"encoding/json"

	"strings"
)

// 创建ELB时，新建公网IP请求参数
type CreateLoadBalancerPublicIpOption struct {
	// ip版数。 有效值：4表示IPv4，暂不支持新建IPv6
	IpVersion *int32 `json:"ip_version,omitempty"`
	// 弹性公网IP的网络类型，默认5_bgp，更多请参考弹性公网ip创建
	NetworkType string `json:"network_type"`
	// 账单信息 如果billinginfo不为空,说明是包周期的带宽
	BillingInfo *string `json:"billing_info,omitempty"`
	// 弹性公网IP的描述信息，不支持特殊字符
	Description *string                            `json:"description,omitempty"`
	Bandwidth   *CreateLoadBalancerBandwidthOption `json:"bandwidth"`
}

func (o CreateLoadBalancerPublicIpOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateLoadBalancerPublicIpOption struct{}"
	}

	return strings.Join([]string{"CreateLoadBalancerPublicIpOption", string(data)}, " ")
}
