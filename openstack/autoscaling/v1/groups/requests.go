package groups

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type HealthCheckMethod string
type InstanceTerminatePolicy string

const (
	ELB_AUDIT  HealthCheckMethod = "ELB_AUDIT"
	NOVA_AUDIT HealthCheckMethod = "NOVA_AUDIT"

	OLD_CONFIG_OLD_INSTANCE InstanceTerminatePolicy = "OLD_CONFIG_OLD_INSTANCE"
	OLD_CONFIG_NEW_INSTANCE InstanceTerminatePolicy = "OLD_CONFIG_NEW_INSTANCE"
	OLD_INSTANCE            InstanceTerminatePolicy = "OLD_INSTANCE"
	NEW_INSTANCE            InstanceTerminatePolicy = "NEW_INSTANCE"
)

//CreateGroupBuilder is an interface from which can build the request of creating group
type CreateGroupBuilder interface {
	ToCreateGroupMap() (map[string]interface{}, error)
}

//CreateGroupOps is a struct contains the parameters of creating group
type CreateGroupOps struct {
	GroupName                 string                  `json:"scaling_group_name" required:"true"`
	ConfigurationID           string                  `json:"scaling_configuration_id,omitempty"`
	DesireInstanceNumber      int                     `json:"desire_instance_number,omitempty"`
	MinInstanceNumber         int                     `json:"min_instance_number,omitempty"`
	MaxInstanceNumber         int                     `json:"max_instance_number,omitempty"`
	CoolDownTime              int                     `json:"cool_down_time,omitempty"`
	LBInstanceID              string                  `json:"lb_listener_id,omitempty`
	AvailableZones            []string                `json:"available_zones,omitempty"`
	Networks                  []NetworkOps            `json:"networks" required:"ture"`
	SecurityGroup             []SecurityGroupOps      `json:"security_groups" required:"ture"`
	VpcID                     string                  `json:"vpc_id" required:"ture"`
	HealthPeriodicAuditMethod HealthCheckMethod       `json:"health_periodic_audit_method,omitempty"`
	HealthPeriodicAuditTime   int                     `json:"health_periodic_audit_time,omitempty"`
	InstanceTerminatePolicy   InstanceTerminatePolicy `json:"instance_terminate_policy,omitempty"`
	Notifications             []string                `json:"notifications,omitempty"`
	IsDeletePublicip          bool                    `json:"delete_publicip,omitempty"`
}

type NetworkOps struct {
	ID string `json:"id,omitempty"`
}

type SecurityGroupOps struct {
	ID string `json:"id,omitempty`
}

func (ops CreateGroupOps) ToCreateGroupMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(ops, "")
}

//CreateGroup is a method of creating group
func Create(client *gophercloud.ServiceClient, ops CreateGroupBuilder) (r CreateGroupResult) {
	b, err := ops.ToCreateGroupMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createGroupUrl(client), b, &r.Body, nil)
	return
}

//DeleteGroup is a method of deleting a group by group id
func Delete(client *gophercloud.ServiceClient, id string) (r DeleteGroupResult) {
	_, r.Err = client.Delete(deleteGroupUrl(client, id), nil)
	return
}

//GetGroup is a method of getting the detailed information of the group by id
func Get(client *gophercloud.ServiceClient, id string) (r GetGroupResult) {
	_, r.Err = client.Get(getGroupUrl(client, id), &r.Body, nil)
	return
}

type ListOpsBuilder interface {
	ToGroupListQuery() (string, error)
}

type ListOps struct {
	GroupName       string `q:"scaling_group_name"`
	ConfigurationID string `q:"scaling_configuration_id"`
	GroupStatus     string `q:"scaling_group_status"`
}

// ToGroupListQuery formats a ListOpts into a query string.
func (opts ListOps) ToGroupListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

func List(client *gophercloud.ServiceClient, ops ListOpsBuilder) pagination.Pager {
	url := listGroupUrl(client)
	if ops != nil {
		q, err := ops.ToGroupListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return GroupPage{pagination.SinglePageBase(r)}
	})
}
