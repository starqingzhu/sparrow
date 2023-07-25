package websocket

import (
	"encoding/json"
	"sparrow/pkg/log/zaplog"
	"sparrow/pkg/net/webscok"
	"sync"
	"testing"
	"time"
)

type jsonSt struct {
	ProtocolId int64       `json:"protocolId"`
	Data       interface{} `json:"data"`
}

func TestClient(t *testing.T) {
	var addr = "192.168.59.35:9102"
	webClient, err := webscok.Init(addr)
	if err != nil {
		panic("init failed")
	}

	defer webClient.Close()
	var wg sync.WaitGroup

	var continueFlag = true

	var cgwVerify = "{\n    \"protocolId\":61001,\n    \"data\":{\n        \"userinfo\":\"123xx69\",\n        \"isreconnect\":0,\n        \"centerid\":7905\n    }\n}"
	var cgwLogin = "{\n    \"protocolId\":61012,\n    \"data\":{\n        \"GameVersion\":\"gameVersion=1.1.0&resVersion=1.1.0&patch=30\",\n        \"MaxPacketId\":575,\n        \"deviceInfo\":{\n\t\t\t\"appVersion\":\"1\"\n\t\t},\n        \"Account\":\"123xx69\",\n        \"type\":1,\n        \"token\":\"\",\n        \"uuid\":\"\"\n    }\n}\n"
	var cpEnrollMatch = "{\n    \"protocolId\":384,\n    \"data\":{\n        \"matchId\":\"1111\",\n        \"tbId\":1\n    }\n}"
	var cgwConnectGamesever = "{\n    \"protocolId\":61004,\n    \"data\":{\n        \"gameserverid\":5524,\n        \"isreconnect\":0\n    }\n}"
	var cgLogin = "{\n    \"protocolId\":1,\n    \"data\":{\n        \"GameVersion\":0,\n        \"ProgramVersion\":108,\n        \"MaxPacketId\":253,\n        \"Account\":\"123xx69\",\n        \"sex\":0\n    }\n}"
	var cgEnterSceneOk = "{\n    \"protocolId\":14,\n    \"data\":{\n        \"IsOK\":1\n    }\n}"
	/*
		385 PACKET_PC_ENROLL_MATCH_RET_PAK
		386 PACKET_PC_ENTER_MATCH_ROOM_PAK
		61004 PACKET_CGW_CONNECT_GAMESERVER_PAK
		61005 PACKET_GWC_CONNECT_GAMESERVER_RET_PAK
		1	PACKET_CG_LOGIN_PAK
		2	PACKET_GC_LOGIN_RET_PAK
		13	PACKET_GC_ENTER_SCENE_PAK
		14	PACKET_CG_ENTER_SCENE_OK_PAK
	*/
	var next = map[int64]int64{
		61002: 61012,
		303:   384,
		386:   61004,
		61005: 1,
		13:    14,
	}
	var idMsg = map[int64]string{
		61001: cgwVerify,
		61012: cgwLogin,
		384:   cpEnrollMatch,
		61004: cgwConnectGamesever,
		1:     cgLogin,
		14:    cgEnterSceneOk,
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
			//if endId == id {
			//	zaplog.LoggerSugar.Infof("game over!!!")
			//	break
			//}
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

			err = webClient.WriteTextMessage([]byte(msgSend))
			if err != nil {
				zaplog.LoggerSugar.Errorf("send msg failed, id:%d, msg[%s]", id, msgSend)
				break
			}

			//zaplog.LoggerSugar.Info("go runtine contine true")
			time.Sleep(50 * time.Millisecond)
		}

	}()

	// 主程序发送
	err = webClient.WriteTextMessage([]byte(cgwVerify))
	if err != nil {
		wg.Done()
	}

	wg.Wait()
}
