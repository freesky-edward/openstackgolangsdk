package eips

import (
	"github.com/gophercloud/gophercloud"
)

//applyURL is a method which will build the url of apply public ip
func applyURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("publicips")
}

//getURL is a method which will build the url of public up query
func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("publicips", id)
}

//deleteURL is a method which will build the url of private ip deletion
func deleteURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("privateips", id)
}

//updateURL is a method which will build the url of public ip udpation
func updateURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("publicips", id)
}
