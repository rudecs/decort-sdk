package disks

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type LimitIORequest struct {
	DiskId           uint64 `url:"diskId"`
	IOps             uint64 `url:"iops"`
	TotalBytesSec    uint64 `url:"total_bytes_sec"`
	ReadBytesSec     uint64 `url:"read_bytes_sec"`
	WriteBytesSec    uint64 `url:"write_bytes_sec"`
	TotalIOpsSec     uint64 `url:"total_iops_sec"`
	ReadIOpsSec      uint64 `url:"read_iops_sec"`
	WriteIOpsSec     uint64 `url:"write_iops_sec"`
	TotalBytesSecMax uint64 `url:"total_bytes_sec_max"`
	ReadBytesSecMax  uint64 `url:"read_bytes_sec_max"`
	WriteBytesSecMax uint64 `url:"write_bytes_sec_max"`
	TotalIOpsSecMax  uint64 `url:"total_iops_sec_max"`
	ReadIOpsSecMax   uint64 `url:"read_iops_sec_max"`
	WriteIOpsSecMax  uint64 `url:"write_iops_sec_max"`
	SizeIOpsSec      uint64 `url:"size_iops_sec"`
}

func (drq LimitIORequest) Validate() error {
	if drq.DiskId == 0 {
		return errors.New("validation-error: field DiskId can not be empty or equal to 0")
	}

	return nil
}

func (d Disks) LimitIO(ctx context.Context, req LimitIORequest, options ...opts.DecortOpts) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/disks/limitIO"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	res, err := d.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil

}
