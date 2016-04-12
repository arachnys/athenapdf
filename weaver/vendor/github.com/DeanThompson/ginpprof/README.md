ginpprof
========

[![GoDoc](https://godoc.org/github.com/DeanThompson/ginpprof?status.svg)](https://godoc.org/github.com/DeanThompson/ginpprof)
[![Build
Status](https://travis-ci.org/DeanThompson/ginpprof.svg?branch=master)](https://travis-ci.org/DeanThompson/ginpprof)

A wrapper for [golang web framework gin](https://github.com/gin-gonic/gin) to use `net/http/pprof` easily.

## Install

First install ginpprof to your GOPATH using `go get`:

```sh
go get github.com/DeanThompson/ginpprof
```

## Usage

```go
package main

import (
	"github.com/gin-gonic/gin"

	"github.com/DeanThompson/ginpprof"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// automatically add routers for net/http/pprof
	// e.g. /debug/pprof, /debug/pprof/heap, etc.
	ginpprof.Wrapper(router)

	router.Run(":8080")
}
```

Start this server, and you will see such outputs:

```text
[GIN-debug] GET   /ping                     --> main.func·001 (3 handlers)
[GIN-debug] GET   /debug/pprof/             --> github.com/DeanThompson/ginpprof.func·001 (3 handlers)
[GIN-debug] GET   /debug/pprof/heap         --> github.com/DeanThompson/ginpprof.func·002 (3 handlers)
[GIN-debug] GET   /debug/pprof/goroutine    --> github.com/DeanThompson/ginpprof.func·003 (3 handlers)
[GIN-debug] GET   /debug/pprof/block        --> github.com/DeanThompson/ginpprof.func·004 (3 handlers)
[GIN-debug] GET   /debug/pprof/threadcreate --> github.com/DeanThompson/ginpprof.func·005 (3 handlers)
[GIN-debug] GET   /debug/pprof/cmdline      --> github.com/DeanThompson/ginpprof.func·006 (3 handlers)
[GIN-debug] GET   /debug/pprof/profile      --> github.com/DeanThompson/ginpprof.func·007 (3 handlers)
[GIN-debug] GET   /debug/pprof/symbol       --> github.com/DeanThompson/ginpprof.func·008 (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```

Now visit [http://127.0.0.1:8080/debug/pprof/](http://127.0.0.1:8080/debug/pprof/) and you'll see what you want.

Have Fun.
