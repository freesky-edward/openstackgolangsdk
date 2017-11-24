package policies

import (
	"github.com/gophercloud/gophercloud"
)

//CreateOptsBuilder is an interface by which can serialize the create parameters
type CreateOptsBuilder interface {
	ToPolicyCreateMap() (map[string]interface{}, error)
}

//CreateOpts is a struct which will be used to create a policy
type CreateOpts struct {
	Name           string             `json:"scaling_policy_name" required:"true"`
	ID             string             `json:"scaling_group_id" required:"true"`
	Type           string             `json:"scaling_policy_type" required:"true"`
	AlarmID        string             `json:"alarm_id,omitempty"`
	SchedulePolicy SchedulePolicyOpts `json:"scheduled_policy,omitempty"`
	Action         ActionOpts         `json:"scaling_policy_action,omitempty"`
	CoolDownTime   int                `json:"cool_down_time,omitempty"`
}

type SchedulePolicyOpts struct {
	LaunchTime      string `json:"launch_time" required:"true"`
	RecurrenceType  string `json:"recurrence_type,omitempty"`
	RecurrenceValue string `json:"recurrence_value,omitempty"`
	StartTime       string `json:"start_time,omitempty"`
	EndTime         string `json:"end_time,omitempty"`
}

type ActionOpts struct {
	Operation   string `json:"operation,omitempty"`
	InstanceNum int    `json:"instance_number,omitempty"`
}

func (opts CreateOpts) ToPolicyCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

//Create is a method which can be able to access to create the policy of autoscaling
//service.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToPolicyCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(createURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

//Delete is a method which can be able to access to delete a policy of autoscaling
func Delete(client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id), nil)
	return
}

//Get is a method which can be able to access to get a policy detailed information
func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}
