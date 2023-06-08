package websocket

import (
	"sparrow/pkg/log/zaplog"
	"sparrow/pkg/net/webscok"
	"sync"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	var addr = "192.168.59.35:9102"
	err := webscok.Init(addr)
	if err != nil {
		panic("init failed")
	}

	defer webscok.GWebClient.Close()
	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			_, _, err = webscok.GWebClient.ReadMessage()
			if err != nil {
				break
			}
			zaplog.LoggerSugar.Info("go runtine contine true")
			time.Sleep(50 * time.Millisecond)
		}
	}()

	var cgwVerify = "{\n    \"protocolId\":61001,\n    \"data\":{\n        \"userinfo\":\"123xx68\",\n        \"isreconnect\":0,\n        \"centerid\":7905\n    }\n}"
	var cgLogin = "{\n    \"protocolId\":61012,\n    \"data\":{\n        \"GameVersion\":\"gameVersion=1.1.0&resVersion=1.1.0&patch=30\",\n        \"MaxPacketId\":564,\n        \"deviceInfo\":{\n\t\t\t\"appVersion\":\"1\"\n\t\t},\n        \"Account\":\"123xx68\",\n        \"type\":1,\n        \"token\":\"\",\n        \"uuid\":\"\"\n    }\n}\n"
	var cpEnrollMatch = "{\n    \"protocolId\":384,\n    \"data\":{\n        \"matchId\":\"1111\",\n        \"tbId\":1\n    }\n}"
	var msgList = []string{cgwVerify, cgLogin, cpEnrollMatch}

	for _, v := range msgList {
		webscok.GWebClient.WriteTextMessage([]byte(v))
	}

	wg.Wait()
}
