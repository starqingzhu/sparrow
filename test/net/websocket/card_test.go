package websocket

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
	"sparrow/internal/web"
	"sparrow/internal/web/pb"
	"sparrow/pkg/log/zaplog"
	"sparrow/pkg/net/webscok"
	"sync"
	"testing"
	"time"
)

func TestCard(t *testing.T) {
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
		61002:                         61012,
		302:                           307,
		351:                           354,
		web.PACKET_PC_EMAIL_AWARD_PAK: web.PACKET_CP_REQDRAW_PAK,
		//321:   568,
	}
	var idMsg = map[int64]interface{}{
		61001:                          web.CgwVerify,
		61012:                          web.CgwLogin,
		307:                            web.CpGmEmail,
		354:                            web.CpGmEmailAward,
		web.PACKET_CP_REQDRAW_PAK:      web.CardDraw1,
		web.PACKET_CP_REQBATCHDRAW_PAK: web.CardDraw10,
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

				if id == web.PACKET_PC_EMAIL_INFO_PAK {
					var emailInfo pb.PC_EMAIL_INFO
					data, errData := json.Marshal(srcMsg.Data)
					if errData != nil {
						return errData
					}
					errData = jsonpb.UnmarshalString(string(data), &emailInfo)
					if errData != nil {
						return errData
					}

					dstMsg = web.JsonSt{
						ProtocolId: web.PACKET_CP_EMAIL_AWARD_PAK,
						Data: pb.CP_EMAIL_AWARD{
							Guid: proto.String(emailInfo.Data.Head.GetGuid()),
						},
					}

				} else {
					dstMsg = msgSend
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
