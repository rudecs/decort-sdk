## Version 1.1.2

### Feature

#### Cloudbroker

- Disks
  - Add share method
  - Add unshare method

### Bug fixes

#### CloudAPI

- Account/Create
  - Change the MaxMemoryCapacity field type to int64
  - Change the MaxVDiskCapacity field type to int64
  - Change the MaxCPUCapacity field type to int64
  - Change the MaxNetworkPeerTransfer field type to int64
  - Change the MaxNumPublicIP field type to int64
  - Change the GPUUnits field type to int64
  - Change url to cloudapi
- Account/Models
  - Change the DiskSize field type to float64
  - Change the DiskSizeMax field type to uint64
- RG/Create
  - Change the MaxMemoryCapacity field type to int64
  - Change the MaxVDiskCapacity field type to int64
  - Change the MaxCPUCapacity field type to int64
  - Change the MaxNumPublicIP field type to int64
- RG/Models
  - Change the DiskSize field type to float64
  - Change the DiskSizeMax field type to uint64

#### Cloudbroker

- Account/Create
  - Change the MaxMemoryCapacity field type to int64
  - Change the MaxVDiskCapacity field type to int64
  - Change the MaxCPUCapacity field type to int64
  - Change the MaxNetworkPeerTransfer field type to int64
  - Change the MaxNumPublicIP field type to int64
  - Change the GPUUnits field type to int64
- Account/Models
  - Change the DiskSize field type to float64
  - Change the DiskSizeMax field type to uint64
- RG/Create
  - Change the MaxMemoryCapacity field type to int64
  - Change the MaxVDiskCapacity field type to int64
  - Change the MaxCPUCapacity field type to int64
  - Change the MaxNumPublicIP field type to int64
- RG/Models
  - Change the DiskSize field type to float64
  - Change the DiskSizeMax field type to uint64
