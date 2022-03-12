# config
⚙️ Simple configuration parser

# Usage

## Parse config files
```go
package main

import (
	"github.com/vitego/boilerplate"
	"github.com/vitego/config"
	"log"
)

func main() {
	err := config.Parse("config", boilerplate.Config)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Get value
```go
package test

import (
	"github.com/vitego/config"
)

func test() string {
	return config.Get("app.name")
}
```
