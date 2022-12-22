package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for limit IO
type LimitIORequest struct {
	// ID of the disk to limit
	// Required: true
	DiskID uint64 `url:"diskId"`

	// Alias for total_iops_sec for backwards compatibility
	// Required: false
	IOPS uint64 `url:"iops,omitempty"`

	// TotalBytesSec
	// Required: false
	TotalBytesSec uint64 `url:"total_bytes_sec,omitempty"`

	// ReadBytesSec
	// Required: false
	ReadBytesSec uint64 `url:"read_bytes_sec,omitempty"`

	// WriteBytesSec
	// Required: false
	WriteBytesSec uint64 `url:"write_bytes_sec,omitempty"`

	// TotalIOPSSec
	// Required: false
	TotalIOPSSec uint64 `url:"total_iops_sec,omitempty"`

	// ReadIOPSSec
	// Required: false
	ReadIOPSSec uint64 `url:"read_iops_sec,omitempty"`

	// WriteIOPSSec
	// Required: false
	WriteIOPSSec uint64 `url:"write_iops_sec,omitempty"`

	// TotalBytesSecMax
	// Required: false
	TotalBytesSecMax uint64 `url:"total_bytes_sec_max,omitempty"`

	// ReadBytesSecMax
	// Required: false
	ReadBytesSecMax uint64 `url:"read_bytes_sec_max,omitempty"`

	// WriteBytesSecMax
	// Required: false
	WriteBytesSecMax uint64 `url:"write_bytes_sec_max,omitempty"`

	// TotalIOPSSecMax
	// Required: false
	TotalIOPSSecMax uint64 `url:"total_iops_sec_max,omitempty"`

	// ReadIOPSSecMax
	// Required: false
	ReadIOPSSecMax uint64 `url:"read_iops_sec_max,omitempty"`

	// WriteIOPSSecMax
	// Required: false
	WriteIOPSSecMax uint64 `url:"write_iops_sec_max,omitempty"`

	// SizeIOPSSec
	// Required: false
	SizeIOPSSec uint64 `url:"size_iops_sec,omitempty"`
}

func (drq LimitIORequest) validate() error {
	if drq.DiskID == 0 {
		return errors.New("validation-error: field DiskID can not be empty or equal to 0")
	}

	return nil
}

// LimitIO limit IO for a certain disk
// total and read/write options are not allowed to be combined
// see http://libvirt.org/formatdomain.html#elementsDisks iotune section for more details
func (d Disks) LimitIO(ctx context.Context, req LimitIORequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/disks/limitIO"

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
