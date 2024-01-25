package config

import (
	"github.com/beevik/etree"
	"sparrow/pkg/log/zaplog"
)

type ServerCfg struct {
	addr string
}

var GServerCfg = &ServerCfg{}

func init() {
	GServerCfg.Load()
}

func (s *ServerCfg) Load() bool {
	doc := etree.NewDocument()
	fileName := "E:\\project\\go\\sparrow\\output\\server.xml"
	err := doc.ReadFromFile(fileName)

	if err != nil {
		zaplog.LoggerSugar.Panicf("server config load failed, err:%s, filename:%s", err.Error(), fileName)
		return false
	}

	// 获取根元素
	//root := doc.Root()
	root := doc.SelectElement("Server")
	if root == nil {
		zaplog.LoggerSugar.Panicf("server config load element failed, not find Server")
		return false
	}
	childNet := root.SelectElement("net")
	if childNet == nil {
		zaplog.LoggerSugar.Panicf("server config load element failed, not find net")
		return false
	}
	GServerCfg.addr = childNet.SelectAttrValue("wbsockAddr", "")

	return true
}

func (s *ServerCfg) Addr() string {
	return s.addr
}
