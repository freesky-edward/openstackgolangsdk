package groups

import (
	"github.com/gophercloud/gophercloud"
)

func createGroupUrl(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scaling_group")
}

func deleteGroupUrl(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("scaling_group", id)
}

func getGroupUrl(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("scaling_group", id)
}
