package clients

//NewRdsV1Client returns a *ServiceClient for making calls
// to the Huawei Cloud RDS V1 API
func NewRdsV1Client() (*gophercloud.ServiceCient, error) {
	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	client, err := openstack.AuthenticatedClient(ao)
	if err != nil {
		return nil, err
	}

	openstack.NewRdsServiceV1(client, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
}
