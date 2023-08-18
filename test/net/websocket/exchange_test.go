package websocket

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"sparrow/internal/web"
	"sparrow/internal/web/pb"
	"sparrow/pkg/log/zaplog"
	"sparrow/pkg/net/webscok"
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
		web.PACKET_GWC_VERIFY_PAK:                 web.PACKET_CGW_CLIENT_LOGIN_PAK,
		web.PACKET_PC_ENTER_WORLD_PAK:             web.PACKET_CP_ENROLL_MATCH_PAK,
		web.PACKET_PC_ENTER_MATCH_ROOM_PAK:        web.PACKET_CGW_CONNECT_GAMESERVER_PAK,
		web.PACKET_GWC_CONNECT_GAMESERVER_RET_PAK: web.PACKET_CG_LOGIN_PAK,
		web.PACKET_GC_ENTER_SCENE_PAK:             web.PACKET_CG_ENTER_SCENE_OK_PAK,
		web.PACKET_GC_MOVE_PAK:                    web.PACKET_CG_EXCHANGE_JJGOLD_PAK,
	}

	var idMsg = map[int64]interface{}{
		web.PACKET_CGW_VERIFY_PAK:             web.CgwVerify,
		web.PACKET_CGW_CLIENT_LOGIN_PAK:       web.CgwLogin,
		web.PACKET_CP_ENROLL_MATCH_PAK:        web.CpEnrollMatch,
		web.PACKET_CGW_CONNECT_GAMESERVER_PAK: web.CgwConnectGamesever,
		web.PACKET_CG_LOGIN_PAK:               web.CgLogin,
		web.PACKET_CG_ENTER_SCENE_OK_PAK:      web.CgEnterSceneOk,
		web.PACKET_CG_EXCHANGE_JJGOLD_PAK:     web.CgExchangeJJGold,
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

			var msg web.JsonSt
			err = json.Unmarshal(p, &msg)
			if err != nil {
				zaplog.LoggerSugar.Errorf("read errro, err:%s", err.Error())
				continue
			}

			id := msg.ProtocolId

			//判断是否结束
			zaplog.LoggerSugar.Infof("recv msg:%s", string(p))

			//修改发送消息
			err = func(id int64, srcMsg *web.JsonSt) error {

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

				if id == web.PACKET_PC_ENTER_MATCH_ROOM_PAK {
					var matchRoomInfo pb.PC_ENTER_MATCH_ROOM
					data, errData := json.Marshal(srcMsg.Data)
					errData = jsonpb.UnmarshalString(string(data), &matchRoomInfo)
					if errData != nil {
						return errData
					}

					web.UpdateCgwConnectGameserverReq(matchRoomInfo.GetGServerId())
					idMsg[web.PACKET_CGW_CONNECT_GAMESERVER_PAK] = web.CgwConnectGamesever
					dstMsg = web.CgwConnectGamesever

				} else if id == web.PACKET_GC_MOVE_PAK {
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
					zaplog.LoggerSugar.Errorf("send msg failed, id:%d, msg[%v]", id, msgSend)
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
	sendVerify, sendErr := json.Marshal(web.CgwVerify)
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
