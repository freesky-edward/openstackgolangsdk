package instances

import (
	"github.com/gophercloud/gophercloud"
)

//getURL will build the querystring by which can be able to query all the instances
//of group
func listURL(client *gophercloud.ServiceClient, groupID string) string {
	return client.ServiceURL("scaling_group_instance", groupID, "list")
}

//deleteURL will build the query url by which can be able to delete an instance from
//the group
func deleteURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("scaling_group_instance", id)
}

//batchURL will build the query url by which can be able to batch add or delete
//instances
func batchURL(client *gophercloud.ServiceClient, groupID string) string {
	return client.ServiceURL("scaling_group_instance", groupID, "action")
}
