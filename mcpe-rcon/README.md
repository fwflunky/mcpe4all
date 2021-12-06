# mcpe-rcon

A Minecraft RCON library written in Go.

Forked from [github.com/Kelwing/mc-rcon](https://github.com/Kelwing/mc-rcon), cleaned up messages and translated text

## Installation
```
go get github.com/fwflunky/mcpe4all/mcpe-rcon
```

## Basic usage
```go
package main

import (
    "log"

    mcpercon "github.com/fwflunky/mcpe4all/mcpe-rcon"
)

func main() {
    conn := new(mcpercon.MCConn)
    err := conn.Open("localhost:25575", "testpw")
    if err != nil {
        log.Fatalln("Open failed", err)
    }
    defer conn.Close()

    err = conn.Authenticate()
    if err != nil {
        log.Fatalln("Auth failed", err)
    }
    
    resp, err := conn.SendCommand("tps")
    if err != nil {
        log.Fatalln("Command failed", err)
    }
    log.Println(resp)
}
```
