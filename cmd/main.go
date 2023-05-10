package main

import (
	"fmt"
	"sparrow/pkg/cache/redis"
)

func main() {
	fmt.Println("enter.........")
	redis.Init()
}
