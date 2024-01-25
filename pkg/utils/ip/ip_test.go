package ip

import "testing"

func TestGetIp(t *testing.T) {
	ip := GetLocalIP()
	addr := "http://169.254.169.254/latest/meta-data/public-ipv4"
	pubIp := GetPublicIp(addr)

	t.Logf("ip:%s, pubIp:%s", ip, pubIp)
}
