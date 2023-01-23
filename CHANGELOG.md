## Version 1.1.0

### Features

#### CloudAPI

- Account
  - Delete "ResTypes" field in Create/Update request structs
  - Add fields "DiskSizeMax", "Shareable", "SEPs" in Get/List response structs
- BService
  - Add fields "SEPID", "SEPPool" in GroupAdd request structs
  - Add field "PoolName" in List/ListDeleted response structs
- Compute
  - Add fields "PresentTo", "Shareable" in Get/List/ListDeleted response structs
- Disks
  - Add fields "PresentTo", "Shareable", "Computes" in Get/List/ListDeleted/ListUnattached/Search response structs
  - Delete fields "ComputeID", "ComputeName" in List/ListDeleted/ListUnattached/Search response structs
  - Add Share/Unshare methods
- FLIPgroup
  - Add field "ClientNames" in Get response struct
- Image
  - Add field "PresentTo" in Get response struct
- RG
  - Delete "ResTypes" field in Create/Update request structs
  - Add fields "DiskSizeMax", "Shareable", "SEPs" in Get/List response structs

#### Cloudbroker

- Account
  - Add fields "SEPs", "ResourceTypes", "PresentTo", "DiskSizeMax", "UniqPools", "Shareable" in Get/List/ListDeleted/ListDisks/ListRG response structs
- Compute
  - Add fields "VINSConnected", "TotalDiskSize", "Shareable", "PresentTo" in Get/List/ListDeleted response structs
- Disks
  - Add fields "ReferenceID", "Shareable", "PresentTo", "Computes" in List/ListDeleted/ListUnattached/Search response structs
  - Delete fields "ComputeID", "ComputeName" in List/ListDeleted/ListUnattached/Search response structs
  - Add Share/Unshare methods
- Grid
  - Add fields "SEPs", "DiskSizeMax" in Get/List response structs
- Image
  - Add field "PresentTo" in Get response struct
- KVMX86
  - Add field "Userdata" in MassCreate request struct
  - Delete field "IPAddr" in MassCreate request struct
- KVMPPC
  - Add field "Userdata" in MassCreate request struct
  - Delete field "IPAddr" in MassCreate request struct
- RG
  - Add fields "DiskSizeMax", "Shareable", "SEPs" in Get/List response structs
