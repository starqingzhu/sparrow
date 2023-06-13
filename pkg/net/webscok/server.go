package webscok

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sparrow/pkg/log/zaplog"
	"sync"
	"time"
)

var Main *Server

type (
	Server struct {
		Upgrader   *websocket.Upgrader
		Router     *mux.Router
		Conns      sync.Map //map[string]*Connection
		CCloseChan chan string
	}
	Connection struct {
		Conn *websocket.Conn

		inChan    chan []byte
		outChan   chan []byte
		closeChan chan string

		mutex    sync.Mutex
		isClosed bool
	}
)

const (
	InitConnNum  = 20
	ChanCloseLen = 30
)

func Init(addr string) {
	if Main == nil {
		Main = &Server{}
	}

	var upgrader = &websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	r := mux.NewRouter()

	Main.Upgrader = upgrader
	Main.Router = r
	Main.CCloseChan = make(chan string, ChanCloseLen)

	//注册 回调
	r.HandleFunc("/ws", Main.CallBackWs)

	go func() {
		zaplog.LoggerSugar.Infof("listenandserve addr:%s", addr)
		err := http.ListenAndServe(addr, r)
		if err != nil {
			zaplog.LoggerSugar.Errorf("listen and serve failed, err:%s\n", err.Error())
			return
		}
	}()
}

func InitConnection(c *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		Conn:      c,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan string, 1),
	}
	go conn.readLoop()
	go conn.writeLoop()
	return
}

func (s *Server) CallBackWs(w http.ResponseWriter, r *http.Request) {

	//建立连接
	connWs, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	zaplog.LoggerSugar.Infof("websocket recept success, client:%s, local:%s", connWs.RemoteAddr(), connWs.LocalAddr())

	//端实例对象
	var conn *Connection
	conn, err = InitConnection(connWs)

	go func() {
		for {
			if errWrite := conn.WriteMessage([]byte("heartbeat")); errWrite != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		data, errRead := conn.ReadMessage()
		if errRead != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}

	//添加连接关闭回调
	//client.Conn.SetCloseHandler(client.HandlerClose)
	//conn.Conn.SetCloseHandler(func(code int, text string) error {
	//	zaplog.LoggerSugar.Infof("Connection closed with code %d and text %s", code, text)
	//	return nil
	//})
	//添加连接管理
	//s.AddConns(conn.RemoteAddr().String(), client)

ERR:
	conn.Close()

}

func (s *Server) HandleMessage(conn *websocket.Conn, messageType int, message []byte) {
	// 处理来自客户端的消息
	fmt.Printf("Received: %s\n", message)

	// 将消息发送回客户端
	err := conn.WriteMessage(messageType, message)
	if err != nil {
		log.Println(err)
		return
	}
}

/*
-----------------连接管理----------------
*/
func (s *Server) AddConns(key string, val interface{}) {
	s.Conns.Store(key, val)
	zaplog.LoggerSugar.Infof("server AddConns %s", key)
}

func (s *Server) DelConns(key string) {
	s.Conns.Delete(key)
	zaplog.LoggerSugar.Infof("server DelConns %s", key)
}

func (s *Server) GetConn(key string) *Connection {
	val, b := s.Conns.Load(key)
	if !b {
		return nil
	}

	ret, ok := val.(*Connection)
	if !ok {
		zaplog.LoggerSugar.Errorf("GetConn key:%s failed", key)
		return nil
	}
	zaplog.LoggerSugar.Infof("server GetConn %s", key)
	return ret
}

/*
----------------------------连接管理------------------------
*/
func (c *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-c.inChan:
	case <-c.closeChan:
		err = errors.New("connection is closed")

	}
	zaplog.LoggerSugar.Infof("read data:%s, addr:%s", string(data), c.Conn.RemoteAddr().String())
	return
}

func (c *Connection) WriteMessage(data []byte) (err error) {
	select {
	case c.outChan <- data:
	case <-c.closeChan:
		err = errors.New("connection is closed")
	}
	zaplog.LoggerSugar.Infof("write data:%s, addr:%s", string(data), c.Conn.RemoteAddr().String())
	return
}

func (c *Connection) Close() {
	// 线程安全，可多次调用
	c.Conn.Close()
	// 利用标记，让closeChan只关闭一次
	c.mutex.Lock()
	if !c.isClosed {
		close(c.closeChan)
		c.isClosed = true
	}
	c.mutex.Unlock()
}

func (c *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = c.Conn.ReadMessage(); err != nil {
			goto ERR
		}

		select {
		case c.inChan <- data:
		case <-c.closeChan:
			goto ERR
		}
	}

ERR:
	c.Close()
}

func (c *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)

	for {
		select {
		case data = <-c.outChan:
		case <-c.closeChan:
			goto ERR
		}

		if err = c.Conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}

ERR:
	c.Close()
}
