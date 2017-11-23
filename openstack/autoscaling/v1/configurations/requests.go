package configurations

import (
	"github.com/gophercloud/gophercloud"
)

type CreateOptsBuilder interface {
	ToConfigurationCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	Name           string             `json:"scaling_configuration_name" required:"true"`
	InstanceConfig InstanceConfigOpts `json:"instance_config" required:"true"`
}

func (opts CreateOpts) ToConfigurationCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(&opts, "")
}

//InstanceConfigOpts is an inner struct of CreateOpts
type InstanceConfigOpts struct {
	ID          string                 `json:"instance_id,omitempty"`
	FlavorRef   string                 `jsno:"flavorRef,omitempty"`
	ImageRef    string                 `json:"imageRef,omitempty"`
	Disk        DiskOpts               `json:"disk,omitempty"`
	SSHKey      string                 `json:"key_name,omitempty"`
	Personality []PersonalityOpts      `json:"personality,omitempty"`
	PubicIp     PublicIpOpts           `json:"public_ip,omitempty"`
	UserData    string                 `json:"user_data,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"` //TODO not sure the type
}

//DiskOpts is an inner struct of InstanceConfigOpts
type DiskOpts struct {
	Size       int    `json:"size" required:"true"`
	VolumeType string `json:"volume_type" required:"true"`
	DiskType   string `json:"disk_type" required:"true"`
}

type PersonalityOpts struct {
	Path    string `json:"path" required:"true"`
	Content string `json:"content" required:"true"`
}

type PublicIpOpts struct {
	Eip EipOpts `json:"eip" required:"true"`
}

type EipOpts struct {
	IpType    string        `json:"ip_type" required:"true"`
	Bandwidth BandwidthOpts `json:"bandwidth" required:"true"`
}

type BandwidthOpts struct {
	Size         int    `json:"size" required:"true"`
	ShareType    string `json:"share_type" required:"true"`
	ChargingMode string `json:"charging_mode" required:"true"`
}

//Create is a method by which can be able to access to create a configuration
//of autoscaling
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToConfigurationCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

//Get is a method by which can be able to access to get a configuration of
//autoscaling detailed information
func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

func Delete(client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id), nil)
	return
}
