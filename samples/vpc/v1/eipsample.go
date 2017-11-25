package main

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/eips"
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
	sc, _ := openstack.NewVpcServiceV1(client, eo)

	ip, err := eips.Apply(sc, eips.ApplyOpts{
		IP: eips.PublicIpOpts{
			Type:    "",
			Address: "",
		},
		Bandwidth: eips.BandwidthOpts{
			Name:      "",
			Size:      2,
			ShareType: "",
		},
	}).Extract()

	if err != nil {
		return
	}
	fmt.Println(ip.ID)

	err1 := eips.Delete(sc, ip.ID).ExtractErr()
	if err1 != nil {
		fmt.Println(err1)
	}
}
