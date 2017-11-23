package main

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/autoscaling/v1/configurations"
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

	page, err := groups.List(sc, groups.ListOpts{}).AllPages()
	if err != nil {
		return
	}

	gs, err := page.(groups.GroupPage).Extract()
	if err != nil {
		return
	}

	for _, group := range gs {
		fmt.Println(group.Name)
	}

	g, err := groups.Create(sc, groups.CreateOpts{Name: "", Networks: nil, SecurityGroup: nil, VpcID: ""}).Extract()

	if err != nil {
		return
	}
	fmt.Println(g)

	//TODO deal with the instance object

	p, err := configurations.List(sc, configurations.ListOpts{}).AllPages()
	if err != nil {
		return
	}

	cs, err := p.(configurations.ConfigurationPage).Extract()
	if err != nil {
		return
	}
	for _, config := range cs {
		fmt.Println(config.Name)
	}

	id, err := configurations.Create(sc, configurations.CreateOpts{Name: "hello", InstanceConfig: configurations.InstanceConfigOpts{}}).Extract()
	if err != nil {
		return
	}
	fmt.Println(id)
}
