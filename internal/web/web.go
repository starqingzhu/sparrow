package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"sparrow/internal/web/pb"
	"sparrow/pkg/log/zaplog"
	"sparrow/pkg/net/webscok"
	"sync"
	"sync/atomic"
	"time"
)

type WebLoginClient struct {
	Addr     string
	UserInfo string
	//WG          *sync.WaitGroup
	TotalClient *atomic.Int64
}

func (c *WebLoginClient) Run() {
	defer func() {
		if err := recover(); err != nil {
			zaplog.LoggerSugar.Errorf("panic %s is disconnect, err:%v", c.Addr, err)
		}
	}()

	//defer c.WG.Done()
	defer c.TotalClient.Add(-1)

	webClient, err := webscok.Init(c.Addr)
	if err != nil {
		zaplog.LoggerSugar.Errorf("websocket init failed, err:%s, addr:%s, userInfo:%s",
			err.Error(), c.Addr, c.UserInfo)
		return
	}

	defer webClient.Close()
	var wgClient sync.WaitGroup
	var continueFlag = true

	// 持续读取
	wgClient.Add(1)
	go func() {
		defer wgClient.Done()

		for continueFlag {

			var p []byte
			_, p, err = webClient.ReadMessage()
			if err != nil {
				zaplog.LoggerSugar.Errorf("read errro, err:%s", err.Error())
				break
			}

			var msg JsonSt
			err = json.Unmarshal(p, &msg)
			if err != nil {
				zaplog.LoggerSugar.Errorf("read errro, err:%s", err.Error())
				continue
			}

			id := msg.ProtocolId

			//判断是否结束
			zaplog.LoggerSugar.Infof("recv msg:%s", string(p))

			//修改发送消息
			err = func(id int64, srcMsg *JsonSt) error {

				idSend, ok := LoginNext[id]
				if !ok {
					if id == PACKET_PC_LOGIN_RET_PAK {
						//time.Sleep(200 * time.Millisecond)
						webClient.Conn.Close()
						zaplog.LoggerSugar.Infof("%s disconnect", c.Addr)
						return errors.New("tcp disconnect")
					}
					return nil
				}

				var dstMsg interface{}
				msgSend, okSend := LoginIdMsg[idSend]
				if !okSend {
					zaplog.LoggerSugar.Errorf("LoginNext is not exist id:%d", id)
					return errors.New(fmt.Sprintf("LoginNext is not exist id:%d", id))
				}

				if id == PACKET_GWC_VERIFY_PAK {
					var CgwLoginReq = JsonSt{
						ProtocolId: PACKET_CGW_CLIENT_LOGIN_PAK,
						Data: pb.CGW_CLIENT_LOGIN{
							GameVersion: proto.String("gameVersion=1.1.0&resVersion=1.1.0&patch=30"),
							MaxPacketId: proto.Int32(PACKET_CP_MAX),
							DeviceInfo: &pb.DeviceInfoClient{
								AppVersion: proto.String("1"),
							},
							Account: proto.String(c.UserInfo),
							Type:    proto.Int32(int32(pb.CGW_CLIENT_LOGIN_TEST_LOGIN)),
							Token:   proto.String(""),
							Uuid:    proto.String(""),
						},
					}

					dstMsg = CgwLoginReq

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
	CgwVerifySend := JsonSt{
		ProtocolId: PACKET_CGW_VERIFY_PAK,
		Data:       map[string]interface{}{"userinfo": c.UserInfo, "isreconnect": 0, "centerid": 0},
	}
	sendVerify, sendErr := json.Marshal(CgwVerifySend)
	if sendErr != nil {
		//wgClient.Done()
		return
	}
	err = webClient.WriteTextMessage(sendVerify)
	if err != nil {
		//wgClient.Done()
	}

	wgClient.Wait()
	//c.TotalClient.Add(-1)
}
