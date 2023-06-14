package gorillaweb

import (
	"errors"
	"github.com/gorilla/websocket"
	"sparrow/pkg/log/zaplog"
	"sync"
)

type (
	Connection struct {
		Conn *websocket.Conn

		inChan    chan []byte
		outChan   chan []byte
		closeChan chan string

		mutex    sync.Mutex
		isClosed bool
	}
)

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
