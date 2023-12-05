package apollo

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

// 支持同一key可以watch多次
type watchan chan map[string]string
type watchanArray []watchan

// Client ...
type Client struct {
	sync.RWMutex
	Notifications map[string]int //notifications 信息
	Configs       map[string]*result
	ChMap         map[string]map[string]watchanArray
	Cli           *http.Client
	Conf          *Conf
	Protal        *ProtalConf
}

// NewWriteableApolloClient 创建基于apollo的可写入config client
func NewWriteableApolloClient(conf *Conf, protal *ProtalConf) *Client {
	cli := &Client{
		Cli:           &http.Client{Timeout: 15 * time.Second},
		Conf:          conf,
		Notifications: make(map[string]int),
		Configs:       make(map[string]*result),
		ChMap:         make(map[string]map[string]watchanArray),
	}
	cli.loadFromCache()
	go cli.watchUpdate()
	cli.Protal = protal
	return cli
}

//NewApolloClient 创建基于apollo的config client
func NewApolloClient(conf *Conf) *Client {
	cli := &Client{
		Cli:           &http.Client{Timeout: 15 * time.Second},
		Conf:          conf,
		Notifications: make(map[string]int),
		Configs:       make(map[string]*result),
		ChMap:         make(map[string]map[string]watchanArray),
	}
	cli.loadFromCache()
	go cli.watchUpdate()
	return cli
}

// GetKeyValue ...
func (c *Client) GetKeyValue(path *Path) (map[string]string, error) {
	if err := c.sync(path.NameSpace); err != nil && len(c.Configs) == 0 {
		return nil, err
	}
	if value, ok := c.Configs[path.NameSpace]; ok {
		kv := make(map[string]string)
		for k, v := range value.Configurations {
			if strings.Contains(k, path.Key) {
				kv[k] = v
			}
		}
		return kv, nil
	}
	return nil, nil
}

// WatchKeyValue ...
func (c *Client) WatchKeyValue(path *Path) chan map[string]string {
	// 每次watch 动态增加namespace, 重复watch的key, idx从-1开始
	c.updateNotificationConf(&notification{path.NameSpace, -1})

	c.Lock()
	defer c.Unlock()
	ch := make(watchan)
	if _, ok := c.ChMap[path.NameSpace]; !ok {
		c.ChMap[path.NameSpace] = make(map[string]watchanArray)
	}
	arr := c.ChMap[path.NameSpace][path.Key]
	if arr == nil {
		arr = make(watchanArray, 0, 1)
	}
	arr = append(arr, ch)
	c.ChMap[path.NameSpace][path.Key] = arr
	return ch
}

// PutKeyValue ...
func (c *Client) PutKeyValue(path *Path, value string) error {
	if c.Protal == nil {
		return fmt.Errorf("readonly client")
	}

	// 先尝试修改
	r := c.Protal.makeModifyRequest(path.Key, value, path.NameSpace)
	modifyResp, err := SendRequest(r)
	if err != nil {
		return err
	}
	if len(modifyResp) != 0 {
		// 需要添加
		addReq := c.Protal.makeAddRequest(path.Key, value, path.NameSpace)
		if body, err := SendRequest(addReq); err != nil {
			return err
		} else if err = checkResponse(body); err != nil {
			return err
		}
	}
	// 发布
	relReq := c.Protal.makeReleaseRequest(path.NameSpace)
	if body, err := SendRequest(relReq); err != nil {
		return err
	} else if err = checkResponse(body); err != nil {
		return err
	}

	return nil
}

// DelKeyValue ...
func (c *Client) DelKeyValue(path *Path) error {
	if c.Protal == nil {
		return fmt.Errorf("readonly client")
	}
	// 删除key
	r := c.Protal.makeDelRequest(path.Key, path.NameSpace)
	if body, err := SendRequest(r); err != nil {
		return err
	} else if len(body) != 0 {
		return fmt.Errorf("del fail %v", string(body))
	}

	// 发布
	relReq := c.Protal.makeReleaseRequest(path.NameSpace)
	if body, err := SendRequest(relReq); err != nil {
		return err
	} else if err = checkResponse(body); err != nil {
		return err
	}

	return nil
}
