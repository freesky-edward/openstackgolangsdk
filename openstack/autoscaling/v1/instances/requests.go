package instances

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

//ListOptsBuilder is an interface by which can be able to build the query string
//of the list function
type ListOptsBuilder interface {
	ToInstancesListQuery() (string, error)
}

type ListOpts struct {
	LifeCycleStatus string `q:"life_cycle_state"`
	HealthStatus    string `q:"health_status"`
}

func (opts ListOpts) ToInstancesListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

//List is a method by which can be able to access the list function that can get
//instances of a group
func List(client *gophercloud.ServiceClient, groupID string, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client, groupID)
	if opts != nil {
		q, err := opts.ToInstancesListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += q
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return InstancePage{pagination.SinglePageBase(r)}
	})
}

//DeleteOptsBuilder is an interface by whick can be able to build the query string
//of instance deletion
type DeleteOptsBuilder interface {
	ToInstanceDeleteQuery() (string, error)
}

type DeleteOpts struct {
	DeleteInstance bool `q:"instance_delete"`
}

func (opts DeleteOpts) ToInstanceDeleteQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

//Delete is a method by which can be able to delete an instance from a group
func Delete(client *gophercloud.ServiceClient, id string, opts DeleteOptsBuilder) (r DeleteResult) {
	url := deleteURL(client, id)
	if opts != nil {
		q, err := opts.ToInstanceDeleteQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += q
	}
	_, r.Err = client.Delete(url, nil)
	return
}
