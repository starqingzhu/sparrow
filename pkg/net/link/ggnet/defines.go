package ggnet

/*
------------------------common-------------------------------------
*/
// 网络类型
type NetworkType string

const (
	Tcp       = "link"
	Udp       = "udp"
	WebSocket = "websocket"
)

var (
	invalidIpString = "-"
)

/*
-----------------------client--------------------------------------
*/

/*
-----------------------server--------------------------------------
*/
var globalSessionId IDUint64 = 0

/*
-------------------------------protocol------------------------------------
*/
// Protocol format:
//
// * 0                                   4
// * +---------------+-------------------+
// * |  	       body len              |
// * +-----------+-----------+-----------+
// * |                                   |
// * +                                   +
// * |           body bytes              |
// * +                                   +
// * |            ... ...                |
// * +-----------------------------------+.

const (
	TRANS_HEAD_LEN = 4
)
