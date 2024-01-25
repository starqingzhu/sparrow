package assetfile

import (
	//"github.com/recastnavigation/recast"
	"testing"
)

func TestAssetFile(t *testing.T) {
	//data, err := ioutil.ReadFile("./NavMesh.asset")
	//if err != nil {
	//	t.Errorf("open file failed, err:%s", err.Error())
	//	return
	//}
	//
	//mapDataStr := string(data)
	//mapDataRowStrs := strings.FieldsFunc(mapDataStr, func(r rune) bool {
	//	return r == '\n'
	//})
	//mapData := make([][]int, len(mapDataRowStrs))
	//for i, rowStr := range mapDataRowStrs {
	//	rowStr = strings.TrimSpace(rowStr)
	//	if rowStr == "" {
	//		continue
	//	}
	//	rowFields := strings.FieldsFunc(rowStr, func(r rune) bool {
	//		return r == ' '
	//	})
	//	mapData[i] = make([]int, len(rowFields))
	//	for j, fieldStr := range rowFields {
	//		field, err1 := strconv.Atoi(fieldStr)
	//		if err1 != nil {
	//			// 处理地图数据解析错误
	//		}
	//		mapData[i][j] = field
	//	}
	//}
	//
	//t.Logf("sucess........\n")
	//navMesh := &recast.NavMesh{}
	//success, err := navMesh.Decode(data)
	//if err != nil {
	//	// 处理NavMesh数据解析错误
	//}
	//if !success {
	//	// 处理NavMesh数据解码失败
	//}
}
