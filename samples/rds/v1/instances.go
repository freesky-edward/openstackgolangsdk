package main

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/rds/v1/instances"
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
	sc, _ := openstack.NewRdsServiceV1(client, eo)

	r := instances.Create(sc, instances.CreateOps{Name: "", FlavorRef: "", ReplicaOf: ""})
	if r.Err != nil {
		//TODO log the error
		return
	}

	var instance instances.Instance
	err3 := r.ExtractInto(&instance)
	if err3 != nil {
		//TODO deal with the error
		return
	}

	//TODO deal with the instance object
}
