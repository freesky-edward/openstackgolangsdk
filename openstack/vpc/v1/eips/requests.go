package eips

import (
	"github.com/gophercloud/gophercloud"
)

//ApplyOptsBuilder is an interface by which can build the request body of public ip
//application
type ApplyOptsBuilder interface {
	ToPublicIpApplyMap() (map[string]interface{}, error)
}

//ApplyOpts is a struct which is used to create public ip
type ApplyOpts struct {
	IP        PublicIpOpts  `json:"publicip" required:"true"`
	Bandwidth BandwidthOpts `json:"bandwidth" required:"true"`
}

type PublicIpOpts struct {
	Type    string `json:"type" required:"true"`
	Address string `json:"ip_address,omitempty"`
}

type BandwidthOpts struct {
	Name       string `json:"name" required:"true"`
	Size       int    `json:"size" required:"true"`
	ShareType  string `json:"share_type" required:"true"`
	ChargeMode string `json:"charge_mode,omitempty"`
}

func (opts ApplyOpts) ToPublicIpApplyMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

//Apply is a method by which can access to apply the public ip
func Apply(client *gophercloud.ServiceClient, opts ApplyOptsBuilder) (r ApplyResult) {
	b, err := opts.ToPublicIpApplyMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(applyURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
