package instances

import (
	"github.com/gophercloud/gophercloud"
)

var RequestOpts gophercloud.RequestOpts = gophercloud.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
}

//CreateOpsBuilder is used for creating instance parameters.
//any struct providing the parameters should implement this interface
type CreateOpsBuilder interface {
	ToInstanceCreateMap() (map[string]interface{}, error)
}

//CreateOps is a struct that contains all the parameters.
type CreateOps struct {

	//instance name
	Name string `json:"name" required:"true"`

	//data store
	Datastore *DatastorOps `json:"datastore,omitempty"`

	FlavorRef string `json:"flavorRef" required:"true"`

	Volume *VolumeOps `json:"volume,omitempty"`

	Region string `json:"region,omitempty"`

	AvailabilityZone string `json:"availability,omitempty"`

	Vpc string `json:"vpc,omitempty"`

	Nics *NicsOps `json:"nics,omitempty"`

	SecurityGroup *SecurityGroupOps `json:"security,omitempty"`

	BackupStrategy *BackupStrategyOps `json:"backupstrategy,omitempty"`

	HA *HAOps `json:"ha,omitempty"`

	DbRtPd string `json:"dbRtPd,omitempty"`

	ReplicaOf string `json:"replicaOf" required:"true"`
}

type DatastorOps struct {
	Type    string `json:"type" required:"true"`
	Version string `json:"version" required:"true"`
}

type VolumeOps struct {
	Type string `json:"type" required:"true"`
	Size int    `json:"size" required:"true"`
}

type NicsOps struct {
	SubnetId string `json:"subnetId" required:"true"`
}

type SecurityGroupOps struct {
	ID string `json:"id" required:"true"`
}

type BackupStrategyOps struct {
	StartTime string `json:"startTime" required:"true"`
	KeepDays  int    `json:"keepDays,omitempty"`
}

type HAOps struct {
	Enable          bool   `json:"enable" required:"true"`
	ReplicationMode string `json:"replicationMode" required:"true"`
}

func (ops CreateOps) ToInstanceCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(ops, "instance")
}

//Create a instance with given parameters.
func Create(client *gophercloud.ServiceClient, ops CreateOpsBuilder) (r CreateResult) {
	b, err := ops.ToInstanceCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(createURL(client), b, &r.Body, &RequestOpts)

	return
}

//delete a instance via id
func Delete(client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id), &RequestOpts)
	return
}

//get a instance with detailed information by id
func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, &RequestOpts)
	return
}

//list all the instances
func List(client *gophercloud.ServiceClient) (r ListResult) {
	_, r.Err = client.Get(listURL(client), &r.Body, &RequestOpts)
	return
}
