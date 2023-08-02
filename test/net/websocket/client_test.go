package websocket

import (
	"encoding/json"
	"sparrow/pkg/log/zaplog"
	"sparrow/pkg/net/webscok"
	"sync"
	"testing"
	"time"
)

/*
	//gate认证相关
	61001	PACKET_CGW_VERIFY_PAK
	61002	PACKET_GWC_VERIFY_PAK
	61012	PACKET_CGW_CLIENT_LOGIN_PAK

	//登录相关
	301 PACKET_CP_LOGIN_PAK
	302 PACKET_PC_LOGIN_RET_PAK
	303 PACKET_PC_ENTER_WORLD_PAK

	//匹配相关
	384 PACKET_CP_ENROLL_MATCH_PAK
	385 PACKET_PC_ENROLL_MATCH_RET_PAK
	386 PACKET_PC_ENTER_MATCH_ROOM_PAK
	61004 PACKET_CGW_CONNECT_GAMESERVER_PAK
	61005 PACKET_GWC_CONNECT_GAMESERVER_RET_PAK
	1	PACKET_CG_LOGIN_PAK
	2	PACKET_GC_LOGIN_RET_PAK
	13	PACKET_GC_ENTER_SCENE_PAK
	14	PACKET_CG_ENTER_SCENE_OK_PAK

	//背包相关
	320 PACKET_PC_ITEM_LIST_PAK
	321 PACKET_PC_ITEM_CHANGE_PAK
	568	PACKET_CP_ONEKEYFUSION_PAK //一键融合

	//gm相关
	307	PACKET_CP_GM_COMMAND_PAK
	308	PACKET_PC_GM_COMMAND_PAK

	//邮件相关
	350	PACKET_PC_EMAIL_LIST_PAK
	351 PACKET_PC_EMAIL_INFO_PAK
	352 PACKET_CP_EMAIL_READ_PAK
	353	PACKET_PC_EMAIL_READ_PAK
	354 PACKET_CP_EMAIL_AWARD_PAK
	355 PACKET_PC_EMAIL_AWARD_PAK
	356 PACKET_CP_EMAIL_DEL_PAK
	357 PACKET_PC_EMAIL_DEL_PAK

*/

type jsonSt struct {
	ProtocolId int64       `json:"protocolId"`
	Data       interface{} `json:"data"`
}

var cgwVerify = "{\n    \"protocolId\":61001,\n    \"data\":{\n        \"userinfo\":\"123xx69\",\n        \"isreconnect\":0,\n        \"centerid\":0\n    }\n}"
var cgwLogin = "{\n    \"protocolId\":61012,\n    \"data\":{\n        \"GameVersion\":\"gameVersion=1.1.0&resVersion=1.1.0&patch=30\",\n        \"MaxPacketId\":575,\n        \"deviceInfo\":{\n\t\t\t\"appVersion\":\"1\"\n\t\t},\n        \"Account\":\"123xx69\",\n        \"type\":1,\n        \"token\":\"\",\n        \"uuid\":\"\"\n    }\n}\n"

var cpEnrollMatch = "{\n    \"protocolId\":384,\n    \"data\":{\n        \"matchId\":\"1111\",\n        \"tbId\":1\n    }\n}"
var cgwConnectGamesever = "{\n    \"protocolId\":61004,\n    \"data\":{\n        \"gameserverid\":5524,\n        \"isreconnect\":0\n    }\n}"
var cgLogin = "{\n    \"protocolId\":1,\n    \"data\":{\n        \"GameVersion\":0,\n        \"ProgramVersion\":108,\n        \"MaxPacketId\":253,\n        \"Account\":\"123xx69\",\n        \"sex\":0\n    }\n}"
var cgEnterSceneOk = "{\n    \"protocolId\":14,\n    \"data\":{\n        \"IsOK\":1\n    }\n}"

var cpGmItem = "{\n    \"data\":{\n        \"command\":\"additem,1001,99\",\n        \"name\":\"123xx69\",\n        \"type\":4\n    },\n    \"protocolId\":307\n}"
var cpGmEmail = "{\n    \"data\":{\n        \"command\":\"sendmail,title,content,1002,1\",\n        \"name\":\"123xx69\",\n        \"type\":1\n    },\n    \"protocolId\":307\n}"

var cpOnekeyFusion = "{\n    \"data\":{\n        \"equipType\":-1\n    },\n    \"protocolId\":568\n}"

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
		303:   307,
		321:   568,
	}
	var idMsg = map[int64]string{
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
