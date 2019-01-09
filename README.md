Nest [![Coverage Status](https://coveralls.io/repos/github/jtsiros/nest/badge.svg?branch=master)](https://coveralls.io/github/jtsiros/nest?branch=master) [![GoDoc](https://godoc.org/github.com/jtsiros/nest?status.svg)](https://godoc.org/github.com/jtsiros/nest) ![Version](https://img.shields.io/badge/version-1.0.0-green.svg)
====

<p align="center">
  <img src="https://cdn.dribbble.com/users/1330537/screenshots/3878129/attachments/880649/hex_gopher_stand_.5.png" alt="Gopher Stand by: Kari Linder"/>
</p>

A Go library for Nest devices. This library provides basic support for Nest Cameras (work-in-progress), Thermostats, and SmokeCoAlarms. There is support for integrating golang OAuth2.0 support into the HTTP client and is expected when constructing a new client. 

## Installation
    go get github.com/jtsiros/nest

## Usage

### Devices

#### Existing Token
```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jtsiros/nest"
	"github.com/jtsiros/nest/auth"
	"github.com/jtsiros/nest/config"
)

func main() {
	// Interactive OAuth2 configuration
	appConfig := config.Config{
		APIURL: config.APIURL,
	}

	conf := auth.NewConfig(appConfig)
	tok, err := auth.NewConfigWithToken("[TOKEN]").Token()
	if err != nil {
		log.Fatal(err)
	}
	client := conf.Client(context.Background(), tok)

	n, err := nest.NewClient(appConfig, client)
	fmt.Println(n.Devices())
}

```

#### No existing Token
```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jtsiros/nest"
	"github.com/jtsiros/nest/auth"
	"github.com/jtsiros/nest/config"
	"golang.org/x/oauth2"
)

func main() {
	// Interactive OAuth2 configuration
	appConfig := config.Config{
		ClientID: "[CLIENT_ID]",
		Secret:   "[SECRET]",
		APIURL:   config.APIURL,
	}

	conf := auth.NewConfig(appConfig)
	url := conf.AuthCodeURL("STATE")

	fmt.Printf("Enter code from this authorization URL: %v\n", url)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	token, err := conf.Exchange(ctx, code,
		oauth2.SetAuthURLParam("client_id", appConfig.ClientID),
		oauth2.SetAuthURLParam("client_secret", appConfig.Secret),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, token)
	n, err := nest.NewClient(appConfig, client)

	fmt.Println(n.Devices())
}
```

### Thermostats
```go
thermostat, err := n.Thermostats.Get("[DEVICE_ID]")
// ... error handling
fmt.Println(thermostat.TargetTemperatureF)

n.Thermostats.SetHVACMode(thermostat.DeviceID, nest.Heat)
```

### SmokeCoAlarms
```go
smokeCoAlarm, err := n.SmokeCoAlarms.Get("[DEVICE_ID]")
// ... error handling
fmt.Println(smokeCoAlarm.LastConnection)
```

### Cameras
At this time, only read-only portion of the API is implemented. I'm planning on implementing the write calls
once I integrate with my HomeKit integration.
```go
camera, err := n.Cameras.Get("[DEVICE_ID]")
// ... error handling
fmt.Println(camera.IsStreaming)
```
## Author
Jon Tsiros

## Credits

Go Gopher Coding it up by: Kari Linder
