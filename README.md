
# graceful-listener

[![GoDoc](https://godoc.org/github.com/ernestoalejo/graceful-listener?status.svg)](https://godoc.org/github.com/ernestoalejo/graceful-listener)

Graceful shutdown for any kind of net.Listener in Go.


## Installation

```shell
go get github.com/ernestoalejo/graceful-listener
```


### Usage

```go
package main

import (
  "log"

  "github.com/ernestoalejo/graceful-listener"
)

func main() {
  listener, err := graceful.NewGracefulListener("localhost:3000")
  if err != nil {
    log.Fatal(err)
  }

  // ... use listener
}
```
