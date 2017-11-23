package v1

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/rds/v1/instances"
)

//Test how to create and delete instance
func TestInstanceCreateDelete(T *testing.T) {
	client, err := clients.NewRdsV1Client()
	if err != nil {
		t.Fatalf("Unable to create a rds client: %v", err)
	}

	instance, err := CreateInstance(T, client)

	if err != nil {
		t.Fatalf("Unable to create instance: %v", err)
	}
	defer DeleteInstance(t, client, instance)

	newInstance, err := instances.Get(client, instance.ID).Extract()
	if err != nil {
		t.Errorf("Unable to retrieve instance: %v", err)
	}

	tools.PrintResource(t, newInstance)
}

func TestInstanceList(T *testing.T) {
	client, err := clients.NewRdsV1Client()
	if err != nil {
		t.Fatalf("Unable to create a rds client: %v", err)
	}

	instances, err := instances.List(client).Extract()
	if err != nil {
		t.Fatalf("Unable to extract instances: %v", err)
	}

	for _, instance := range instances {
		tools.PrintResource(t, instance)
	}
}
