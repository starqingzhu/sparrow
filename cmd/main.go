package main

import (
	"fmt"
	"sparrow/pkg/cache/myredigo"
)

func main() {
	fmt.Println("enter.........")
	myredigo.Init()
}
