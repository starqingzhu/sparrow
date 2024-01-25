package main

import (
	"sparrow/config"
	"sparrow/pkg/net/webscok/gorillaweb"
	"time"
)

func main() {
	//var addr = flag.String("addr", "127.0.0.1:8080", "eg:-addr=127.0.0.1:8080")
	//flag.Parse()

	gorillaweb.Init(config.GServerCfg.Addr())
	for {
		time.Sleep(500 * time.Millisecond)
	}

	////参数解析
	//var addr = flag.String("addr", "127.0.0.1:8080", "eg:-addr=127.0.0.1:8080")
	//flag.Parse()
	//
	////link 初始化
	//var protocol = ggnet.NewFramedProtocol(4, 64*1024*1024, 64*1024*1024, binary.LittleEndian)
	//opts := ggnet.NewServerSpecWithOpts(ggnet.WithAddr(*addr), ggnet.WithNetwork(ggnet.Tcp), ggnet.WithSendChanSize(10))
	//
	//var handler ggnet.HandlerFunc = func(sess ggnet.Session) {}
	//var server = ggnet.NewTcpServer(opts, protocol, handler)
	//err := server.Run()
	//if err != nil {
	//	zaplog.LoggerSugar.Errorf("server init failed, %s", addr)
	//	return
	//}
	//zaplog.LoggerSugar.Infof("server init success, %s", addr)
}
