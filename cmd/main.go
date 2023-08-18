package main

import (
	"flag"
	"fmt"
	"runtime"
	"sparrow/internal/web"
	"sync/atomic"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	//f, err := os.Create("trace.out")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()

	//启动trace goruntine
	//err = trace.Start(f)
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer trace.Stop()

	//debug.SetMaxThreads(10)

	addr := flag.String("addr", "10.110.55.212:9102", "enter addr eg:192.168.59.184:9102")
	clientCount := flag.Int64("clientCount", 2, "enter clientCount eg:1")
	flag.Parse()

	fmt.Printf("addr:%s, clientCount:%d\n", *addr, *clientCount)

	//var wg sync.WaitGroup
	//var clientCount = 2

	var totalClient atomic.Int64

	var clientNum int64 = 0

	for {

		if totalClient.Load() < *clientCount {
			var clientUser = fmt.Sprintf("user_%d", clientNum)
			client := web.WebLoginClient{
				Addr:     *addr,
				UserInfo: clientUser,
				//WG:          &wg,
				TotalClient: &totalClient,
			}
			totalClient.Add(1)

			go client.Run()

			fmt.Printf("current client count %d, totalClient:%d\n", clientNum, totalClient.Load())
			clientNum++

		} else {
			time.Sleep(500 * time.Millisecond)
		}

	}

	//for clientNum < *clientCount {
	//	var clientUser = fmt.Sprintf("user%d", clientNum)
	//	client := web.WebLoginClient{
	//		Addr:     *addr,
	//		UserInfo: clientUser,
	//		//WG:          &wg,
	//		TotalClient: &totalClient,
	//	}
	//	totalClient.Add(1)
	//
	//	go client.Run()
	//
	//	fmt.Printf("current client count %d, totalClient:%d\n", clientNum, totalClient.Load())
	//	clientNum++
	//
	//}

	//wg.Wait()

	//redis.Init()
	time.Sleep(1000 * time.Second)
}
