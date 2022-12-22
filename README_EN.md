# Decort SDK

Decort SDK is a library, written in GO (Golang) for interact with the **DECORT** API.  
The library contents structures and methods for requesting to an user (cloudapi) and admin (cloudbroker) groups of API.  
Also the library have structures for responses.

## Contents

- [Install](#install)
- [API List](#api-list)
- [Examples](#examples)
- [Examples2](#examples2)

## Install

```bash
go get -u github.com/rudecs/decort-sdk
```

## API List

## Examples

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rudecs/decort-sdk/config"
	"github.com/rudecs/decort-sdk/pkg/cloudapi/kvmx86"
)

func main() {
	cfg := config.Config{
		AppID:     "<APPID>",
		AppSecret: "<APPSECRET>",
		SSOURL:    "https://sso.digitalenergy.online",
		DecortURL: "https://mr4.digitalenergy.online",
		Retries:   5,
	}
	client := decort.New(cfg)
	req := kvmx86.CreateRequest{
		RGID:    123,
		Name:    "compute",
		CPU:     4,
		RAM:     4096,
		ImageID: 321,
	}

	res, err := client.KVMX86().Create(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
```

## Examples2
