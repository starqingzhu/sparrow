package main

import (
	"sparrow/pkg/net/webscok"
	"time"
)

func main() {
	var addr = "127.0.0.1:8000"

	webscok.Init(addr)
	for {
		time.Sleep(500 * time.Millisecond)
	}
}
