package ip

import (
	"io/ioutil"
	"net"
	"net/http"
)

// GetLocalIP 获取本地ip地址
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, a := range addrs {
		if ip4 := toIP4(a); ip4 != nil {
			return ip4.String()
		}
	}
	return ""
}

func toIP4(addr net.Addr) net.IP {
	if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
		return ipnet.IP.To4()
	}
	return nil
}

func GetPublicIp(addr string) string {
	resp, err := http.Get(addr)
	if err != nil {
		return GetLocalIP()
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == 200 {
		return string(body)
	} else {
		return GetLocalIP()
	}
}
