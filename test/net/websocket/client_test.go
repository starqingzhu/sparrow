package websocket

import (
	"sparrow/pkg/net/webscok"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	//var addr = "192.168.59.35:9102"
	var addr = "127.0.0.1:8000"
	webClient, err := webscok.Client(addr)
	if err != nil {
		panic("init failed")
	}

	defer webClient.Close()

	for {
		_, _, err = webClient.ReadMessage()
		if err != nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	//var wg sync.WaitGroup
	//
	//var continueFlag = true
	//
	////等待结束
	//go func() {
	//	select {
	//	case <-time.After(3 * time.Second):
	//		continueFlag = false
	//
	//	}
	//}()
	//
	//// 持续读取
	//go func() {
	//	wg.Add(1)
	//	defer wg.Done()
	//	for continueFlag {
	//		_, _, err = webClient.ReadMessage()
	//		if err != nil {
	//			break
	//		}
	//		zaplog.LoggerSugar.Info("go runtine contine true")
	//		time.Sleep(50 * time.Millisecond)
	//	}
	//
	//}()
	//
	////go func() {
	////	// 主程序发送
	////	var cgwVerify = "{\n    \"protocolId\":61001,\n    \"data\":{\n        \"userinfo\":\"123xx68\",\n        \"isreconnect\":0,\n        \"centerid\":7905\n    }\n}"
	////	var cgwLogin = "{\n    \"protocolId\":61012,\n    \"data\":{\n        \"GameVersion\":\"gameVersion=1.1.0&resVersion=1.1.0&patch=30\",\n        \"MaxPacketId\":564,\n        \"deviceInfo\":{\n\t\t\t\"appVersion\":\"1\"\n\t\t},\n        \"Account\":\"123xx68\",\n        \"type\":1,\n        \"token\":\"\",\n        \"uuid\":\"\"\n    }\n}\n"
	////	var cpEnrollMatch = "{\n    \"protocolId\":384,\n    \"data\":{\n        \"matchId\":\"1111\",\n        \"tbId\":1\n    }\n}"
	////	var cgwConnectGamesever = "{\n    \"protocolId\":61004,\n    \"data\":{\n        \"gameserverid\":5524,\n        \"isreconnect\":0\n    }\n}"
	////	var cgLogin = "{\n    \"protocolId\":1,\n    \"data\":{\n        \"GameVersion\":0,\n        \"ProgramVersion\":108,\n        \"MaxPacketId\":243,\n        \"Account\":\"123xx68\",\n        \"sex\":0\n    }\n}"
	////	var cgEnterSceneOk = "{\n    \"protocolId\":14,\n    \"data\":{\n        \"IsOK\":1\n    }\n}"
	////	var msgList = []string{cgwVerify, cgwLogin, cpEnrollMatch, cgwConnectGamesever, cgLogin, cgEnterSceneOk}
	////	for _, v := range msgList {
	////		webClient.WriteTextMessage([]byte(v))
	////		time.Sleep(500 * time.Millisecond)
	////	}
	////}()
	//
	//wg.Wait()
}
