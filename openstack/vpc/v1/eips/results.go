package eips

import (
	"github.com/gophercloud/gophercloud"
)

//ApplyResult is a struct which represents the result of apply public ip
type ApplyResult struct {
	gophercloud.Result
}

func (r ApplyResult) Extract() (PublicIp, error) {
	var ip struct {
		Ip PublicIp `json:"publicip"`
	}
	err := r.Result.ExtractInto(&ip)
	return ip.Ip, err
}

type PublicIp struct {
	ID            string `json:"id"`
	Status        string `json:"status"`
	Type          string `json:"type"`
	Address       string `json:"public_ip_address"`
	TenantID      string `json:"tenant_id"`
	CreateTime    string `json:"create_time"`
	BandwidthSize int    `json:"bandwidth_size"`
}
