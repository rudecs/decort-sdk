package disks

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// Request struct for get information about disk
type GetRequest struct {
	// ID of the disk
	// Required: true
	DiskID uint64 `url:"diskId"`
}

func (drq GetRequest) validate() error {
	if drq.DiskID == 0 {
		return errors.New("validation-error: field DiskID must be set")
	}

	return nil
}

// Get gets  disk details
// Notice: the devicename field is the name as it is passed to the kernel (kname in linux) for unattached disks this field has no relevant value
func (d Disks) Get(ctx context.Context, req GetRequest) (*RecordDisk, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	url := "/cloudbroker/disks/get"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordDisk{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
