package instances

import (
	"github.com/gophercloud/gophercloud"
)

type Instance struct {
	ID               string         `json:"id"`
	Status           string         `json:"status"`
	Name             string         `json:"name"`
	Created          string         `json:"created"`
	HostName         string         `json:"hostname"`
	Type             string         `json:"type"`
	Region           string         `json:"string"`
	Updated          string         `json:"updated"`
	AvailabilityZone string         `json:"availabilityZone"`
	Vpc              string         `json:"vpc"`
	Nics             Nics           `json:"nics"`
	SecurityGroup    SecurityGroup  `json:"securityGroup"`
	Flavor           Flavor         `json:"flavor"`
	Volume           Volume         `json:"volume"`
	DbPort           int            `json:"dbPort"`
	DataStoreInfo    DataStore      `json:"dataStoreInfo"`
	ExtendParameters Extends        `json:"extendParam"`
	BackupStrategy   BackupStrategy `json:"backupStrategy"`
	SlaveId          string         `json:"slaveId"`
	HA               HA             `json:"ha"`
	ReplicaOf        string         `json:"replica_of"`
}

type HA struct {
	ReplicationMode string `json:"replicationMode"`
}

type BackupStrategy struct {
	StartTime string `json:"startTime"`
	KeepDays  int    `json:"keepDays"`
}

type Nics struct {
	SubenetID string `json:"subnetId"`
}

type Flavor struct {
	ID string `json:"id"`
}

type SecurityGroup struct {
	ID string `json:"id"`
}

type Volume struct {
	Type string `json:"type"`
	Size int    `json:"size"`
}

type DataStore struct {
	Type    string `json:"type"`
	Version string `json:"version"`
}

type Extends struct {
	Jobs []Job `json:"jobs"`
}

type Job struct {
	ID string `json:"id"`
}

// Extract will get the Volume object out of the commonResult object.
func (r commonResult) Extract() (*Instance, error) {
	var s Instance
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "instance")
}

type commonResult struct {
	gophercloud.Result
}

// CreateResult contains the response body and error from a Create request.
type CreateResult struct {
	commonResult
}

type DeleteResult struct {
	commonResult
}

type GetResult struct {
	commonResult
}

type ListResult struct {
	gophercloud.Result
}

func (lr ListResult) Extract() ([]Instance, error) {
	var a struct {
		Instances []Instance `json:"instances"`
	}
	err := lr.Result.ExtractInto(&a)
	return a.Instances, err
}
