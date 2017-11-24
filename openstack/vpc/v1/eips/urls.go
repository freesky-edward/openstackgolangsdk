package eips

import (
	"github.com/gophercloud/gophercloud"
)

//applyURL is a method which will build the url of apply public ip
func applyURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("publicips")
}
