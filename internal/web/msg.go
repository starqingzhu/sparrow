package web

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"sparrow/internal/web/pb"
)

/*
	//gate认证相关
	61001	PACKET_CGW_VERIFY_PAK
	61002	PACKET_GWC_VERIFY_PAK
	61012	PACKET_CGW_CLIENT_LOGIN_PAK

	//登录相关
	301 PACKET_CP_LOGIN_PAK
	302 PACKET_PC_LOGIN_RET_PAK
	303 PACKET_PC_ENTER_WORLD_PAK

	//匹配相关
	384 PACKET_CP_ENROLL_MATCH_PAK
	385 PACKET_PC_ENROLL_MATCH_RET_PAK
	386 PACKET_PC_ENTER_MATCH_ROOM_PAK
	61004 PACKET_CGW_CONNECT_GAMESERVER_PAK
	61005 PACKET_GWC_CONNECT_GAMESERVER_RET_PAK
	1	PACKET_CG_LOGIN_PAK
	2	PACKET_GC_LOGIN_RET_PAK
	13	PACKET_GC_ENTER_SCENE_PAK
	14	PACKET_CG_ENTER_SCENE_OK_PAK

	//背包相关
	320 PACKET_PC_ITEM_LIST_PAK
	321 PACKET_PC_ITEM_CHANGE_PAK
	568	PACKET_CP_ONEKEYFUSION_PAK //一键融合

	//gm相关
	307	PACKET_CP_GM_COMMAND_PAK
	308	PACKET_PC_GM_COMMAND_PAK

	//邮件相关
	350	PACKET_PC_EMAIL_LIST_PAK
	351 PACKET_PC_EMAIL_INFO_PAK
	352 PACKET_CP_EMAIL_READ_PAK
	353	PACKET_PC_EMAIL_READ_PAK
	354 PACKET_CP_EMAIL_AWARD_PAK
	355 PACKET_PC_EMAIL_AWARD_PAK
	356 PACKET_CP_EMAIL_DEL_PAK
	357 PACKET_PC_EMAIL_DEL_PAK

*/

type (
	JsonStRet struct {
		ProtocolId int64  `json:"protocolId"`
		Data       string `json:"data"`
	}
)

// "{\n    \"protocolId\":61001,\n    \"data\":{\n        \"userinfo\":\"123xx69\",\n        \"isreconnect\":0,\n        \"centerid\":0\n    }\n}"
var userInfo = "123xx70"
var CgwVerify = JsonSt{
	ProtocolId: PACKET_CGW_VERIFY_PAK,
	Data:       map[string]interface{}{"userinfo": userInfo, "isreconnect": 0, "centerid": 0},
}

// "{\n    \"protocolId\":61012,\n    \"data\":{\n        \"GameVersion\":\"gameVersion=1.1.0&resVersion=1.1.0&patch=30\",\n        \"MaxPacketId\":575,\n        \"deviceInfo\":{\n\t\t\t\"appVersion\":\"1\"\n\t\t},\n        \"Account\":\"123xx69\",\n        \"type\":1,\n        \"token\":\"\",\n        \"uuid\":\"\"\n    }\n}\n"
var CgwLogin = JsonSt{
	ProtocolId: PACKET_CGW_CLIENT_LOGIN_PAK,
	Data: pb.CGW_CLIENT_LOGIN{
		GameVersion: proto.String("gameVersion=1.1.0&resVersion=1.1.0&patch=30"),
		MaxPacketId: proto.Int32(PACKET_CP_MAX),
		DeviceInfo: &pb.DeviceInfoClient{
			AppVersion: proto.String("1"),
		},
		Account: proto.String(userInfo),
		Type:    proto.Int32(int32(pb.CGW_CLIENT_LOGIN_TEST_LOGIN)),
		Token:   proto.String(""),
		Uuid:    proto.String(""),
	},
}

// "{\n    \"protocolId\":384,\n    \"data\":{\n        \"matchId\":\"1111\",\n        \"tbId\":1\n    }\n}"
var CpEnrollMatch = JsonSt{
	ProtocolId: PACKET_CP_ENROLL_MATCH_PAK,
	Data: pb.CP_ENROLL_MATCH{
		MatchId: proto.String(fmt.Sprintf("%d", 111)),
		TbId:    proto.Int32(1),
	},
}

// "{\n    \"protocolId\":61004,\n    \"data\":{\n        \"gameserverid\":5524,\n        \"isreconnect\":0\n    }\n}"
var gameServerId int32 = 18021
var CgwConnectGamesever = JsonSt{
	ProtocolId: PACKET_CGW_CONNECT_GAMESERVER_PAK,
	Data: pb.CGW_CONNECT_GAMESERVER{
		Gameserverid: proto.Int32(gameServerId),
		Isreconnect:  proto.Int32(0),
	},
}

func GetCgwConnectGameserverReq(id int32) *JsonSt {
	return &JsonSt{
		ProtocolId: PACKET_CGW_CONNECT_GAMESERVER_PAK,
		Data: pb.CGW_CONNECT_GAMESERVER{
			Gameserverid: proto.Int32(id),
			Isreconnect:  proto.Int32(1),
		},
	}
}

func UpdateCgwConnectGameserverReq(id int32) {
	CgwConnectGamesever = *GetCgwConnectGameserverReq(id)
}

// "{\n    \"protocolId\":1,\n    \"data\":{\n        \"GameVersion\":0,\n        \"ProgramVersion\":108,\n        \"MaxPacketId\":253,\n        \"Account\":\"123xx69\",\n        \"sex\":0\n    }\n}"
var gameVersion int32 = 0
var programVersion int32 = 108

var CgLogin = JsonSt{
	ProtocolId: PACKET_CG_LOGIN_PAK,
	Data: pb.CG_LOGIN{
		GameVersion:    proto.Int32(gameVersion),
		ProgramVersion: proto.Int32(programVersion),
		MaxPacketId:    proto.Int32(PACKET_CG_MAX),
		Account:        proto.String(userInfo),
		Sex:            proto.Int32(0),
	},
}

// "{\n    \"protocolId\":14,\n    \"data\":{\n        \"IsOK\":1\n    }\n}"
var CgEnterSceneOk = JsonSt{
	ProtocolId: PACKET_CG_ENTER_SCENE_OK_PAK,
	Data: pb.CG_ENTER_SCENE_OK{
		IsOK: proto.Int32(1),
	},
}

// "{\n    \"data\":{\n        \"command\":\"additem,1001,99\",\n        \"name\":\"123xx69\",\n        \"type\":4\n    },\n    \"protocolId\":307\n}"
var cpGmItem = JsonSt{
	ProtocolId: PACKET_CP_GM_COMMAND_PAK,
	Data: pb.CP_GM_COMMAND{
		Command: proto.String("additem,41010,55"),
		Name:    proto.String(userInfo),
		Type:    proto.Int32(4),
	},
}

// "{\n    \"data\":{\n        \"command\":\"sendmail,title,content,1002,1\",\n        \"name\":\"123xx69\",\n        \"type\":1\n    },\n    \"protocolId\":307\n}"
var CpGmEmail = JsonSt{
	ProtocolId: PACKET_CP_GM_COMMAND_PAK,
	Data: pb.CP_GM_COMMAND{
		Command: proto.String("sendmail,title,content,41010,20"),
		Name:    proto.String(userInfo),
		Type:    proto.Int32(1),
	},
}

var CpGmEmailAward = JsonSt{
	ProtocolId: PACKET_CP_EMAIL_AWARD_PAK,
	Data: pb.CP_EMAIL_AWARD{
		Guid: proto.String(fmt.Sprintf("%d", 0)),
	},
}

// "{\n    \"data\":{\n        \"equipType\":-1\n    },\n    \"protocolId\":568\n}"
var cpOnekeyFusion = JsonSt{
	ProtocolId: PACKET_CP_ONEKEYFUSION_PAK,
	Data: pb.CP_OneKeyFusion{
		EquipType: proto.Int32(-1),
	},
}

// 奖池
var CardDraw1 = JsonSt{
	ProtocolId: PACKET_CP_REQDRAW_PAK,
	Data: pb.CP_ReqDraw{
		Id: proto.Int32(1),
	},
}
var CardDraw10 = JsonSt{
	ProtocolId: PACKET_CP_REQBATCHDRAW_PAK,
	Data: pb.CP_ReqBatchDraw{
		Id: proto.Int32(1),
	},
}

// 签到
var SignReq = JsonSt{
	ProtocolId: PACKET_CP_SEVENLOGIN_AWARD_PAK,
	Data: pb.CP_SEVENLOGIN_AWARD{
		Number: proto.Int32(0),
	},
}

var SignAfterReq = JsonSt{
	ProtocolId: PACKET_CP_AFTER_SIGNIN_PAK,
	Data: pb.CP_AFTER_SIGNIN{
		Number: proto.Int32(0),
	},
}

// 兑换
var CgExchangeJJGold = JsonSt{
	ProtocolId: PACKET_CG_EXCHANGE_JJGOLD_PAK,
	Data: pb.CG_Exchange_JJGold{
		Type: proto.Int32(1),
		Num:  proto.Int32(100),
	},
}
