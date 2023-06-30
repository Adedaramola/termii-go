# Termii-go
Golang SDK for termii

## Installation
```bash
go get -u github.com/adedaramola/termii-go
```

## Usage
```go
package main

import (
    "context"
    "log",

    "github.com/adedaramola/termii-go"
)

func main() {
    client, err := termii.NewClient(context.Background(), &termii.Config{
        ApiKey: "dummy-api-key",
    })
    if err != nil {
        log.Println(err)
    }

    token, err := client.SendToken(&termii.SendTokenOptions{})
    if err != nil {
        log.Println(err)
    }

    fmt.Println("Pin ID: ", token.PinID)
}
```