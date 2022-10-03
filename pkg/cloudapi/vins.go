package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/vins"
)

func (ca *CloudApi) VINS() *vins.VINS {
	return vins.New(ca.client)
}
