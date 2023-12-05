package apollo

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func (c *Client) watchUpdate() {
	for _, name := range c.Conf.NameSpaceName {
		c.updateNotificationConf(&notification{name, defaultID})
	}
	c.watchNotificationsUpdate()
}

func request(cli *http.Client, url string) ([]byte, error) {
	resp, err := cli.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return ioutil.ReadAll(resp.Body)
	}

	// Diacard all body if status code is not 200
	io.Copy(ioutil.Discard, resp.Body)
	return nil, nil
}

func toString(m map[string]int) string {
	var notifications []*notification
	for k, v := range m {
		notifications = append(notifications, &notification{
			NamespaceName:  k,
			NotificationID: v,
		})
	}
	bts, err := json.Marshal(&notifications)
	if err != nil {
		return ""
	}

	return string(bts)
}
func (c *Client) updateNotificationConf(notification *notification) {
	c.Lock()
	defer c.Unlock()
	c.Notifications[notification.NamespaceName] = notification.NotificationID
}

func (c *Client) handleNamespaceUpdate(namespace string) error {
	if err := c.sync(namespace); err != nil {
		return err
	}
	return nil
}

func (c *Client) sync(namesapce string) error {
	c.Lock()
	defer c.Unlock()
	releaseKey := c.getReleaseKey(namesapce)
	url := configURL(c.Conf, namesapce, releaseKey)
	bts, err := request(c.Cli, url)
	if err != nil || len(bts) == 0 {
		return err
	}
	var result result
	if err := json.Unmarshal(bts, &result); err != nil {
		return err
	}
	c.Configs[namesapce] = &result
	return nil
}
func (c *Client) getReleaseKey(namesapce string) string {
	if value, ok := c.Configs[namesapce]; ok {
		return value.ReleaseKey
	}
	return ""
}

func (c *Client) watchNotificationsUpdate() {
	timer := time.NewTimer(time.Second * 2)
	flag := true //记录第一次发配置
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			c.RLock()
			size := len(c.ChMap)
			c.RUnlock()
			if size > 0 {
				if err := c.pumpUpdates(); err != nil && flag && len(c.Configs) > 0 {
					flag = false
					for name, _ := range c.ChMap {
						c.sendConfig(name)
					}
				}
			}
			timer.Reset(time.Second * 2)
		}
	}
}

func (c *Client) pumpUpdates() error {
	var ret error

	updates, err := c.poll()
	if err != nil {
		return err
	}
	if len(updates) == 0 {
		return err
	}
	for _, update := range updates {
		if err := c.handleNamespaceUpdate(update.NamespaceName); err != nil {
			ret = err
			continue
		}
		c.sendConfig(update.NamespaceName)
		c.updateNotificationConf(update)
	}
	c.dumpToFile()
	return ret
}

func (c *Client) poll() ([]*notification, error) {
	c.RLock()
	notifications := toString(c.Notifications)
	c.RUnlock()
	url := notificationURL(c.Conf, notifications)
	bts, err := request(c.Cli, url)
	if err != nil || len(bts) == 0 {
		return nil, err
	}
	var ret []*notification
	if err := json.Unmarshal(bts, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *Client) sendConfig(namesapce string) {
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()
	if c.Configs[namesapce] == nil || c.ChMap[namesapce] == nil {
		return
	}
	for path, charr := range c.ChMap[namesapce] {
		kv := make(map[string]string)
		for key, value := range c.Configs[namesapce].Configurations {
			if strings.Contains(key, path) {
				kv[key] = value
			}
		}
		// FIXME: 1、深拷贝 2、异步发送 3、上面的Contains？
		for _, ch := range charr {
			ch <- kv
		}
	}
}

func Load(filename string, v interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}
