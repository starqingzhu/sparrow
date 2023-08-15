package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

func main() {
	runtime.GOMAXPROCS(4)
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//启动trace goruntine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	defer trace.Stop()

	//debug.SetMaxThreads(10)
	fmt.Println("enter.........")

	//redis.Init()
	//time.Sleep(100 * time.Second)
}
