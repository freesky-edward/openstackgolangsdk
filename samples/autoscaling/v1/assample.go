package main

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/autoscaling/v1/groups"
)

func main() {
	opts := gophercloud.AuthOptions{ //TODO change all the parameters
		IdentityEndpoint: "",
		Username:         "",
		UserID:           "",
		Password:         "",
		DomainID:         "",
		DomainName:       "",
		AllowReauth:      true,
	}

	client, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return
	}

	eo := gophercloud.EndpointOpts{Region: "ReionOne"} //TODO change the region

	sc, _ := openstack.NewAutoScalingService(client, eo)

	g, err := groups.Create(sc, groups.CreateGroupOps{GroupName: "", Networks: nil, SecurityGroup: nil, VpcID: ""}).Extract()

	if err != nil {
		return
	}
	fmt.Println(g)

	//TODO deal with the instance object
}
