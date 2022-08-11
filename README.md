# Decort SDK

```go
package main

import (
	"context"
	"fmt"
	"log"

	decort "github.com/rudecs/decort-sdk"
	"github.com/rudecs/decort-sdk/kvmx86"
)

func main() {
	cfg := decort.Config{
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
