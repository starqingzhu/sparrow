package websocket

import (
	"encoding/json"
	"sparrow/pkg/log/zaplog"
	"sparrow/pkg/net/webscok"
	"sync"
	"testing"
	"time"
)

//func TestClient(t *testing.T) {
//	var addr = "192.168.59.184:9102"
//	webClient, err := webscok.Init(addr)
//	if err != nil {
//		panic("init failed")
//	}
//
//	defer webClient.Close()
//	var wg sync.WaitGroup
//
//	var continueFlag = true
//
//	// //测试主流程
//	//var next = map[int64]int64{
//	//	61002: 61012,
//	//	//303:   384,
//	//	//386:   61004,
//	//	//61005: 1,
//	//	//13:    14,
//	//}
//	//var idMsg = map[int64]string{
//	//	61001: cgwVerify,
//	//	61012: cgwLogin,
//	//	//384:   cpEnrollMatch,
//	//	//61004: cgwConnectGamesever,
//	//	//1:     cgLogin,
//	//	//14:    cgEnterSceneOk,
//	//}
//
//	// 测试 背包
//	var next = map[int64]int64{
//		61002: 61012,
//		303:   307,
//		321:   568,
//	}
//	var idMsg = map[int64]string{
//		61001: cgwVerify,
//		61012: cgwLogin,
//		307:   cpGmEmail,
//		//568:   cpOnekeyFusion,
//	}
//	//var endId int64 = 0
//
//	// 持续读取
//	go func() {
//		wg.Add(1)
//		defer wg.Done()
//		for continueFlag {
//
//			var p []byte
//			_, p, err = webClient.ReadMessage()
//			if err != nil {
//				zaplog.LoggerSugar.Errorf("read errro, err:%s", err.Error())
//				break
//			}
//
//			var msg jsonSt
//			err = json.Unmarshal(p, &msg)
//			if err != nil {
//				zaplog.LoggerSugar.Errorf("read errro, err:%s", err.Error())
//				continue
//			}
//
//			id := msg.ProtocolId
//
//			//判断是否结束
//			zaplog.LoggerSugar.Infof("recv msg:%s", string(p))
//
//			idSend, ok := next[id]
//			if !ok {
//				continue
//			}
//			msgSend, okSend := idMsg[idSend]
//			if !okSend {
//				zaplog.LoggerSugar.Errorf("next is not exist id:%d", id)
//				break
//			}
//
//			err = webClient.WriteTextMessage([]byte(msgSend))
//			if err != nil {
//				zaplog.LoggerSugar.Errorf("send msg failed, id:%d, msg[%s]", id, msgSend)
//				break
//			}
//
//			//zaplog.LoggerSugar.Info("go runtine contine true")
//			time.Sleep(50 * time.Millisecond)
//		}
//
//	}()
//
//	// 主程序发送
//	err = webClient.WriteTextMessage([]byte(cgwVerify))
//	if err != nil {
//		wg.Done()
//	}
//
//	wg.Wait()
//}

func TestClient2(t *testing.T) {
	var addr = "192.168.59.184:9102"
	webClient, err := webscok.Init(addr)
	if err != nil {
		panic("init failed")
	}

	defer webClient.Close()
	var wg sync.WaitGroup

	var continueFlag = true

	// 测试 背包
	var next = map[int64]int64{
		61002: 61012,
		302:   307,
		//321:   568,
	}
	var idMsg = map[int64]interface{}{
		61001: cgwVerify,
		61012: cgwLogin,
		307:   cpGmEmail,
		//568:   cpOnekeyFusion,
	}
	//var endId int64 = 0

	// 持续读取
	go func() {
		wg.Add(1)
		defer wg.Done()
		for continueFlag {

			var p []byte
			_, p, err = webClient.ReadMessage()
			if err != nil {
				zaplog.LoggerSugar.Errorf("read errro, err:%s", err.Error())
				break
			}

			var msg jsonSt
			err = json.Unmarshal(p, &msg)
			if err != nil {
				zaplog.LoggerSugar.Errorf("read errro, err:%s", err.Error())
				continue
			}

			id := msg.ProtocolId

			//判断是否结束
			zaplog.LoggerSugar.Infof("recv msg:%s", string(p))

			idSend, ok := next[id]
			if !ok {
				continue
			}
			msgSend, okSend := idMsg[idSend]
			if !okSend {
				zaplog.LoggerSugar.Errorf("next is not exist id:%d", id)
				break
			}

			sendBuf, sendErr := json.Marshal(msgSend)
			if sendErr != nil {
				break
			}

			err = webClient.WriteTextMessage(sendBuf)
			if err != nil {
				zaplog.LoggerSugar.Errorf("send msg failed, id:%d, msg[%s]", id, msgSend)
				break
			}

			//zaplog.LoggerSugar.Info("go runtine contine true")
			time.Sleep(50 * time.Millisecond)
		}

	}()

	// 主程序发送
	sendVerify, sendErr := json.Marshal(cgwVerify)
	if sendErr != nil {
		wg.Done()
		return
	}
	err = webClient.WriteTextMessage(sendVerify)
	if err != nil {
		wg.Done()
	}

	wg.Wait()
}
