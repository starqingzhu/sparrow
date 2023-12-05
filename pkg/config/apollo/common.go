package apollo

import (
	"fmt"
	"io"
	"net/url"
	"sparrow/pkg/utils/ip"
)

const defaultID = -1

// var (
// 	ModifyKVurl string
// 	Releaseurl  string
// 	Deleteurl   string
// 	AddKVurl    string
// 	Token       string
// )

// type config struct {
// 	PortalAddress          string `json:"portalAddress"`
// 	Config_service_Address string `json:"config_service_Address"`
// 	Env                    string `json:"env"`
// 	AppId                  string `json:"appId"`
// 	ClusterName            string `json:"clusterName"`
// 	NamespaceName          string `json:"namespaceName"`
// 	Token                  string `json:"token"`
// }

// func init() {
// 	config := config{}
// 	Load("./config.json", &config)
// 	Token = config.Token
// 	ModifyKVurl = "http://" + config.PortalAddress + "/openapi/v1/envs/" + config.Env + "/apps/" + config.AppId + "/clusters/" + config.ClusterName + "/namespaces/" + config.NamespaceName + "/items/"
// 	Releaseurl = "http://" + config.PortalAddress + "/openapi/v1/envs/" + config.Env + "/apps/" + config.AppId + "/clusters/" + config.ClusterName + "/namespaces/" + config.NamespaceName + "/releases"
// 	Deleteurl = "http://" + config.PortalAddress + "/openapi/v1/envs/" + config.Env + "/apps/" + config.AppId + "/clusters/" + config.ClusterName + "/namespaces/" + config.NamespaceName + "/items/"
// 	AddKVurl = "http://" + config.PortalAddress + "/openapi/v1/envs/" + config.Env + "/apps/" + config.AppId + "/clusters/" + config.ClusterName + "/namespaces/" + config.NamespaceName + "/items"
// }

type Conf struct {
	AppID         string   `json:"appId,omitempty"`
	Cluster       string   `json:"cluster,omitempty"`
	NameSpaceName []string `json:"namespaceNames,omitempty"`
	IP            string   `json:"ip,omitempty"`
}

type ProtalConf struct {
	IP          string `json:"ip"`
	Env         string `json:"env"`
	AppID       string `json:"appId"`
	ClusterName string `json:"clusterName"`
	UserID      string `json:"userId"`
	Token       string `json:"token"`
}

type notification struct {
	NamespaceName  string `json:"namespaceName,omitempty"`
	NotificationID int    `json:"notificationId,omitempty"`
}
type result struct {
	// AppID          string            `json:"appId"`
	// Cluster        string            `json:"cluster"`
	NamespaceName  string            `json:"namespaceName"`
	Configurations map[string]string `json:"configurations"`
	ReleaseKey     string            `json:"releaseKey"`
}

type ReleaseInfo struct {
	AppId          string            `json:"appId"`
	Cluster        string            `json:"cluster"`
	NamespaceName  string            `json:"namespaceName"`
	Configurations map[string]string `json:"configurations"`
	ReleaseKey     string            `json:"releaseKey"`
}

type ModifyBody struct {
	Key                      string `json:"key"`
	Value                    string `json:"value"`
	Comment                  string `json:"comment"`
	DataChangeLastModifiedBy string `json:"dataChangeLastModifiedBy"`
}
type AddBody struct {
	Key                 string `json:"key"`
	Value               string `json:"value"`
	Comment             string `json:"comment"`
	DataChangeCreatedBy string `json:"dataChangeCreatedBy"`
}

type RleaseBody struct {
	ReleaseTitle   string `json:"releaseTitle"`
	ReleaseComment string `json:"releaseComment"`
	ReleasedBy     string `json:"releasedBy"`
}

type Request struct {
	Key    string
	Method string
	Body   io.Reader
	Url    string
	Token  string
}

func notificationURL(conf *Conf, notifications string) string {
	return fmt.Sprintf("http://%s/notifications/v2?appId=%s&cluster=%s&notifications=%s",
		conf.IP,
		url.QueryEscape(conf.AppID),
		url.QueryEscape(conf.Cluster),
		url.QueryEscape(notifications))
}

func configURL(conf *Conf, namespace, releaseKey string) string {
	return fmt.Sprintf("http://%s/configs/%s/%s/%s?releaseKey=%s&ip=%s",
		conf.IP,
		url.QueryEscape(conf.AppID),
		url.QueryEscape(conf.Cluster),
		url.QueryEscape(namespace),
		releaseKey,
		ip.GetLocalIP())
}
