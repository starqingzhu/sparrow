package websocket

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"sparrow/pkg/log/zaplog"
	"sparrow/pkg/net/webscok"
	"sparrow/test/net/websocket/pb"
	"sync"
	"testing"
	"time"
)

func TestExchange(t *testing.T) {
	var addr = "192.168.59.184:9102"
	webClient, err := webscok.Init(addr)
	if err != nil {
		panic("init failed")
	}

	defer webClient.Close()
	var wg sync.WaitGroup

	var continueFlag = true
	var moveCount = 0

	// 测试 背包
	var next = map[int64]int64{
		PACKET_GWC_VERIFY_PAK:                 PACKET_CGW_CLIENT_LOGIN_PAK,
		PACKET_PC_ENTER_WORLD_PAK:             PACKET_CP_ENROLL_MATCH_PAK,
		PACKET_PC_ENTER_MATCH_ROOM_PAK:        PACKET_CGW_CONNECT_GAMESERVER_PAK,
		PACKET_GWC_CONNECT_GAMESERVER_RET_PAK: PACKET_CG_LOGIN_PAK,
		PACKET_GC_ENTER_SCENE_PAK:             PACKET_CG_ENTER_SCENE_OK_PAK,
		PACKET_GC_MOVE_PAK:                    PACKET_CG_EXCHANGE_JJGOLD_PAK,
	}

	var idMsg = map[int64]interface{}{
		PACKET_CGW_VERIFY_PAK:             cgwVerify,
		PACKET_CGW_CLIENT_LOGIN_PAK:       cgwLogin,
		PACKET_CP_ENROLL_MATCH_PAK:        cpEnrollMatch,
		PACKET_CGW_CONNECT_GAMESERVER_PAK: cgwConnectGamesever,
		PACKET_CG_LOGIN_PAK:               cgLogin,
		PACKET_CG_ENTER_SCENE_OK_PAK:      cgEnterSceneOk,
		PACKET_CG_EXCHANGE_JJGOLD_PAK:     cgExchangeJJGold,
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

			//修改发送消息
			err = func(id int64, srcMsg *jsonSt) error {

				idSend, ok := next[id]
				if !ok {
					return nil
				}

				var dstMsg interface{}
				msgSend, okSend := idMsg[idSend]
				if !okSend {
					zaplog.LoggerSugar.Errorf("next is not exist id:%d", id)
					return errors.New(fmt.Sprintf("next is not exist id:%d", id))
				}

				if id == PACKET_PC_ENTER_MATCH_ROOM_PAK {
					var matchRoomInfo pb.PC_ENTER_MATCH_ROOM
					data, errData := json.Marshal(srcMsg.Data)
					errData = jsonpb.UnmarshalString(string(data), &matchRoomInfo)
					if errData != nil {
						return errData
					}

					UpdateCgwConnectGameserverReq(matchRoomInfo.GetGServerId())
					idMsg[PACKET_CGW_CONNECT_GAMESERVER_PAK] = cgwConnectGamesever
					dstMsg = cgwConnectGamesever

				} else if id == PACKET_GC_MOVE_PAK {
					if moveCount != 1 {
						dstMsg = nil
					} else {
						dstMsg = msgSend
					}
					moveCount++
				} else {
					dstMsg = msgSend
				}

				if dstMsg == nil {
					return nil
				}

				sendBuf, sendErr := json.Marshal(dstMsg)
				if sendErr != nil {
					return sendErr
				}
				err = webClient.WriteTextMessage(sendBuf)
				if err != nil {
					zaplog.LoggerSugar.Errorf("send msg failed, id:%d, msg[%s]", id, msgSend)
					return err
				}

				return nil
			}(id, &msg)

			if err != nil {
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
