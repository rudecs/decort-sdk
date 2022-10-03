package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type LimitIORequest struct {
	DiskID           uint64 `url:"diskId"`
	IOPS             uint64 `url:"iops,omitempty"`
	TotalBytesSec    uint64 `url:"total_bytes_sec,omitempty"`
	ReadBytesSec     uint64 `url:"total_bytes_sec,omitempty"`
	WriteBytesSec    uint64 `url:"write_bytes_sec,omitempty"`
	TotalIOPSSec     uint64 `url:"total_iops_sec,omitempty"`
	ReadIOPSSec      uint64 `url:"read_iops_sec,omitempty"`
	WriteIOPSSec     uint64 `url:"write_iops_sec,omitempty"`
	TotalBytesSecMax uint64 `url:"total_bytes_sec_max,omitempty"`
	ReadBytesSecMax  uint64 `url:"read_bytes_sec_max,omitempty"`
	WriteBytesSecMax uint64 `url:"write_bytes_sec_max,omitempty"`
	TotalIOPSSecMax  uint64 `url:"total_iops_sec_max,omitempty"`
	ReadIOPSSecMax   uint64 `url:"total_iops_sec_max,omitempty"`
	WriteIOPSSecMax  uint64 `url:"write_iops_sec_max,omitempty"`
	SizeIOPSSec      uint64 `url:"size_iops_sec,omitempty"`
}

func (drq LimitIORequest) Validate() error {
	if drq.DiskID == 0 {
		return errors.New("validation-error: field DiskID must be set")
	}

	return nil
}

func (d Disks) LimitIO(ctx context.Context, req LimitIORequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/disks/limitIO"
	
	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}