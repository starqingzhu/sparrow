package main

import (
	"sparrow/pkg/net/webscok/gorillaweb"
	"time"
)

func main() {
	var addr = "127.0.0.1:8000"

	gorillaweb.Init(addr)
	for {
		time.Sleep(500 * time.Millisecond)
	}
}
