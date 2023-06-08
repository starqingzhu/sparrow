package webscok

import (
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"sparrow/pkg/log/zaplog"
)

type (
	WebClient struct {
		Conn *websocket.Conn
		Res  *http.Response
	}
)

var GWebClient *WebClient

func init() {
	GWebClient = &WebClient{}
}

// addr:"localhost:8080"
func Init(addr string) error {
	u := url.URL{Scheme: "ws", Host: addr, Path: "/ws"}
	var err error
	GWebClient.Conn, GWebClient.Res, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		zaplog.LoggerSugar.Fatal("net client failed", u.String())
		return nil
	}

	zaplog.LoggerSugar.Infof("%s open success, res:%v", GWebClient.Conn.LocalAddr().String(), *GWebClient.Res)

	return nil
}

func (w *WebClient) Close() {
	if w == nil || w.Conn == nil {
		return
	}

	zaplog.LoggerSugar.Infof("%s close success", w.Conn.LocalAddr().String())
	w.Conn.Close()
	w.Conn = nil
}

func (w *WebClient) WriteTextMessage(messageBody []byte) error {
	err := w.Conn.WriteMessage(websocket.TextMessage, messageBody)
	zaplog.LoggerSugar.Infof("send msg[%s], err:%v", string(messageBody), err)
	return err
}

func (w *WebClient) WriteBinaryMessage(messageBody []byte) error {
	err := w.Conn.WriteMessage(websocket.BinaryMessage, messageBody)
	zaplog.LoggerSugar.Infof("send msg:[%p], err:%v", err)
	return err
}

func (w *WebClient) ReadMessage() (messageType int, p []byte, err error) {
	messageType, p, err = w.Conn.ReadMessage()

	zaplog.LoggerSugar.Infof("recv messageType:%d, msg:[%s], err:%v", messageType, p, err)
	return
}
