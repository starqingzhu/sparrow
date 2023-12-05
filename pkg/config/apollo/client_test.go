package apollo

import (
	"sparrow/pkg/log/zaplog"
	"testing"
)

func TestNewWriteableApolloClient(t *testing.T) {
	appid := "poker"
	cli := NewWriteableApolloClient(
		&Conf{
			AppID:         appid,
			Cluster:       "default",
			NameSpaceName: []string{"application"},
			IP:            "127.0.0.1:8080",
		},
		&ProtalConf{
			IP:          "127.0.0.1:8070",
			Env:         "DEV",
			AppID:       appid,
			ClusterName: "default",
			UserID:      "sunbin",
			Token:       "", //"c955889febd9577b0568314feb9602f67f2f68a4",
		},
	)
	testk1 := &Path{appid, "application", "systemInfo", 0}
	if r, err := cli.GetKeyValue(testk1); err != nil {
		zaplog.LoggerSugar.Errorf("get fail, err:%s, appid:%s", err.Error(), appid)
		return
	} else {
		zaplog.LoggerSugar.Infof("get kv, kv:%s, appid:%s", r, appid)
	}

	type testSt struct {
		Key1 string `json:"key1"`
		Key2 string `json:"key2"`
	}

	testk2 := &Path{appid, "application", "test", 0}
	//var testValue2 = testSt{
	//	Key1: "hello",
	//	Key2: "world",
	//}
	//val, err1 := json.Marshal(testValue2)
	//if err1 != nil {
	//	zaplog.LoggerSugar.Errorf("marshal failed, val:%+v", testValue2)
	//	return
	//}

	//var val = "{\"key1\":\"hello\",\"key2\":\"world\"}"
	//var valStr = string(val)
	//if err := cli.PutKeyValue(testk2, valStr); err != nil {
	//	zaplog.LoggerSugar.Errorf("put fail, err:%s, appid:%s", err.Error(), appid)
	//	return
	//}

	if r, err := cli.GetKeyValue(testk2); err != nil {
		zaplog.LoggerSugar.Errorf("get fail, err:%s, appid:%s", err.Error(), appid)
		return
	} else {
		zaplog.LoggerSugar.Infof("get kv, kv:%s, appid:%s", r, appid)
	}
	//
	//if err := cli.DelKeyValue(testk1); err != nil {
	//	zaplog.LoggerSugar.Errorf("del fail, err:%s, appid:%s", err.Error(), appid)
	//	return
	//}
	//if r, err := cli.GetKeyValue(testk1); err != nil {
	//	zaplog.LoggerSugar.Errorf("get fail, err:%s, appid:%s", err.Error(), appid)
	//	return
	//} else {
	//	zaplog.LoggerSugar.Infof("get kv, kv:%s, appid:%s", r, appid)
	//}
	//
	//for iloop := 0; iloop < 20; iloop++ {
	//	go func() {
	//		for i := 0; i < 10; i++ {
	//			key := fmt.Sprintf("key%d", i)
	//			// if err := cli.DelKeyValue(&Path{appid, "application", key, 0}); err != nil {
	//			// 	log.Error().Str("appid", appid).Str("key", key).Err(err).Msg("del fail")
	//			// }
	//			v := fmt.Sprintf("value%d", i)
	//			if err := cli.PutKeyValue(&Path{appid, "application", key, 0}, v); err != nil {
	//				zaplog.LoggerSugar.Errorf("put fail, err:%s, appid:%s", err.Error(), appid)
	//				return
	//			}
	//			if r, err := cli.GetKeyValue(&Path{appid, "application", key, 0}); err != nil {
	//				zaplog.LoggerSugar.Errorf("get fail, err:%s, appid:%s", err.Error(), appid)
	//				return
	//			} else {
	//				if gv, ok := r[key]; !ok {
	//					zaplog.LoggerSugar.Errorf("get testk fail, key:%s, appid:%s", key, appid)
	//
	//				} else if gv != v {
	//					zaplog.LoggerSugar.Errorf("get testk fail 2, get:%s, set:%s, appid:%s", gv, v, appid)
	//				}
	//			}
	//		}
	//	}()
	//}

	//for _, v := range []string{"testValue", "testValue2", "testValue3",
	//	"testValue4", "testValue5", "testValue6", "testValue7",
	//	"testValue8", "testValue9", "testValue10", "testValue11"} {
	//	if err := cli.PutKeyValue(testk1, v); err != nil {
	//		zaplog.LoggerSugar.Errorf("put fail, err:%s, appid:%s", err.Error(), appid)
	//		return
	//	}
	//	if r, err := cli.GetKeyValue(testk1); err != nil {
	//		zaplog.LoggerSugar.Errorf("get fail, err:%s, appid:%s", err.Error(), appid)
	//		return
	//	} else {
	//		zaplog.LoggerSugar.Infof("get kv, kv:%s, appid:%s", r, appid)
	//		if gv, ok := r[testk1.Key]; !ok {
	//			zaplog.LoggerSugar.Errorf("get testk fail, appid:%s", appid)
	//		} else if gv != v {
	//			zaplog.LoggerSugar.Errorf("get testk fail 2, set:%s, get:%s, appid:%s", v, gv, appid)
	//		}
	//	}
	//}

	for v := range cli.WatchKeyValue(&Path{appid, "application", "systemInfo", 0}) {
		zaplog.LoggerSugar.Infof("get kv, kv:%s, appid:%s", v, appid)
	}
}
