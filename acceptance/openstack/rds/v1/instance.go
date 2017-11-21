package v1

import (
	"testing"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/rds/v1/instances"
)

func CreateInstance(T *testing.T, client *gophercloud.ServiceClient) (*Instance, error) {
	instanceName := tools.RandomString("ACPTTEST", 5)
	t.Logf("Attempting to create instance: %s", instanceName)

	createOps := instances.CreateOps{
		Name:      instanceName,
		FlavorRef: "",
		ReplicaOf: "",
	}

	instance, err := instances.Create(client, createOps).Extract()
	if err != nil {
		return nil, err
	}

	return instance, err
}

func DeleteInstance(T *testing.T, client *gophercloud.ServiceClient, instance Instance) error {
	err := instances.Delete(client, instance.ID).ExtractErr()
	if err != nil {
		t.Fatalf("Unable to delete volume %s: %v", instance.ID, err)
	}
	t.Logf("Deleted instance: %s", instance.ID)
}
