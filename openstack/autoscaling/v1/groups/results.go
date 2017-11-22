package groups

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

//CreateGroupResult is a struct retured by CreateGroup request
type CreateGroupResult struct {
	gophercloud.Result
}

//Extract the create group result as a string type.
func (r CreateGroupResult) Extract() (string, error) {
	var a struct {
		GroupID string `json:"scaling_group_id"`
	}
	err := r.Result.ExtractInto(a)
	return a.GroupID, err
}

//DeleteGroupResult contains the body of the deleting group request
type DeleteGroupResult struct {
	gophercloud.ErrResult
}

//GetGroupResult contains the body of getting detailed group request
type GetGroupResult struct {
	gophercloud.Result
}

//Extract method will parse the result body into Group struct
func (r GetGroupResult) Extract() (*Group, error) {
	var a Group
	err := r.Result.ExtractInto(&a)
	return &a, err
}

//Group represents the struct of one autoscaling group
type Group struct {
	GroupName                 string          `json:"scaling_group_name"`
	GroupID                   string          `json:"scaling_group_id"`
	GroupStatus               string          `json:"scaling_group_status"`
	ConfigurationID           string          `json:"scaling_configuration_id"`
	ConfigurationName         string          `json:"scaling_configuration_name"`
	ActualInstanceNumber      int             `json:"current_instance_number"`
	DesireInstanceNumber      int             `json:"desire_instance_number"`
	MinInstanceNumber         int             `json:"min_instance_number"`
	MaxInstanceNumber         int             `json:"max_instance_number"`
	CoolDownTime              int             `json:"cool_down_time"`
	LBListenerID              string          `json:"lb_listener_id"`
	LBaaSListeners            []interface{}   `json:"lbaas_listeners"`
	AvailableZones            []string        `json:"available_zones"`
	Networks                  []Network       `json:"networks"`
	SecurityGroups            []SecurityGroup `json:"security_groups"`
	CreateTime                string          `json:"create_time"`
	VpcID                     string          `json:"vpc_id"`
	Detail                    string          `json:"detail"`
	IsScaling                 bool            `json:"is_scaling"`
	HealthPeriodicAuditMethod string          `json:"health_periodic_audit_method"`
	HealthPeriodicAuditTime   int             `json:"health_periodic_audit_time"`
	InstanceTerminatePolicy   string          `json:"instance_terminate_policy"`
	Notifications             []string        `json:"notifications"`
	DeletePublicip            bool            `json:"delete_publicip"`
	CloudLocationID           string          `json:"cloud_location_id"`
}

type Network struct {
	ID string `json:"id"`
}

type SecurityGroup struct {
	ID string `json:"id"`
}

type GroupPage struct {
	pagination.SinglePageBase
}

// IsEmpty returns true if a ListResult contains no Volumes.
func (r GroupPage) IsEmpty() (bool, error) {
	groups, err := r.Extract()
	return len(groups) == 0, err
}

func (r GroupPage) Extract() ([]Group, error) {
	var gs []Group
	err := r.Result.ExtractIntoSlicePtr(gs, "scaling_groups")
	return gs, err
}

//this is the action result which is the result of enable or disable operations
type ActionResult struct {
	gophercloud.ErrResult
}
