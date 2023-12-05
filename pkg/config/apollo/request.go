package apollo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sparrow/pkg/log/zaplog"
	"time"

	"io"
	"io/ioutil"
	"net/http"
)

func (p *ProtalConf) makeAddRequest(key, value, namespace string) *Request {
	var tmp = AddBody{Key: key, Value: value, Comment: "", DataChangeCreatedBy: p.UserID}
	data, _ := json.Marshal(tmp)
	body := bytes.NewBuffer([]byte(data))
	url := fmt.Sprintf("http://%s/openapi/v1/envs/%s/apps/%s/clusters/%s/namespaces/%s/items",
		p.IP,
		p.Env,
		p.AppID,
		p.ClusterName,
		namespace)
	req := newRequest("", "POST", url, body, p.Token)
	return req

}

func (p *ProtalConf) makeModifyRequest(key, value, namespace string) *Request {
	var tmp = ModifyBody{Key: key, Value: value, Comment: "", DataChangeLastModifiedBy: p.UserID}
	data, _ := json.Marshal(tmp)
	body := bytes.NewBuffer([]byte(data))
	url := fmt.Sprintf("http://%s/openapi/v1/envs/%s/apps/%s/clusters/%s/namespaces/%s/items/%s",
		p.IP,
		p.Env,
		p.AppID,
		p.ClusterName,
		namespace,
		key)
	req := newRequest("", "PUT", url, body, p.Token)
	return req
}

func (p *ProtalConf) makeDelRequest(key, namespace string) *Request {
	url := fmt.Sprintf("http://%s/openapi/v1/envs/%s/apps/%s/clusters/%s/namespaces/%s/items/%s?operator=%s",
		p.IP,
		p.Env,
		p.AppID,
		p.ClusterName,
		namespace,
		key,
		p.UserID)
	req := newRequest("", "DELETE", url, nil, p.Token)
	return req
}

func (p *ProtalConf) makeReleaseRequest(namespace string) *Request {
	var tmp = RleaseBody{
		ReleaseTitle:   fmt.Sprintf("release by %s at %s", p.UserID, time.Now().Format("2006-01-02 15:04:05")),
		ReleaseComment: "",
		ReleasedBy:     p.UserID,
	}
	data, _ := json.Marshal(tmp)
	body := bytes.NewBuffer([]byte(data))
	url := fmt.Sprintf("http://%s/openapi/v1/envs/%s/apps/%s/clusters/%s/namespaces/%s/releases",
		p.IP,
		p.Env,
		p.AppID,
		p.ClusterName,
		namespace)
	req := newRequest("", "POST", url, body, p.Token)
	return req
}

func checkResponse(respBody []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(respBody, &m)
	if err != nil {
		return err
	}
	if _, ok := m["status"]; ok {
		return fmt.Errorf("resp %v", m)
	}
	return nil
}

func newRequest(key string, method, url string, body io.Reader, token string) *Request {
	return &Request{Key: key, Method: method, Body: body, Url: url, Token: token}
}

func (r *Request) GetKey() string {
	return r.Key
}

func (r *Request) GetMethod() string {
	return r.Method
}

func (r *Request) GetBody() io.Reader {
	return r.Body
}

func (r *Request) GetUrl() string {
	return r.Url
}

func (r *Request) GetToken() string {
	return r.Token
}

func SendRequest(r *Request) (respBody []byte, err error) {
	httpClient := http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest(r.GetMethod(), r.GetUrl(), r.GetBody())
	req.Header.Set("Authorization", r.GetToken())
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if err != nil {
		zaplog.LoggerSugar.Errorf("create request failed, err:%s", err.Error())
		return nil, err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		zaplog.LoggerSugar.Errorf("request failed, err:%s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		zaplog.LoggerSugar.Errorf("read response body failed, err:%s", err.Error())
	}
	return respBody, err
}
