package cloudapi

import "github.com/rudecs/decort-sdk/pkg/cloudapi/bservice"

// Accessing the BService method group
func (ca *CloudAPI) BService() *bservice.BService {
	return bservice.New(ca.client)
}
