package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"sparrow/pkg/log/zaplog"
	"testing"
)

func TestExcel(t *testing.T) {
	f, err := excelize.OpenFile("E:\\project\\go\\sparrow\\output\\config\\ActivityTask.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err = f.Close(); err != nil {
			zaplog.LoggerSugar.Errorf("close excel err:%s", err.Error())
		}
	}()

	l := f.GetSheetList()
	for _, v := range l {
		rows, err1 := f.GetRows(v)
		if err1 != nil {
			zaplog.LoggerSugar.Errorf("GetRows failed, err:%s", err.Error())
			return
		}

		for _, row := range rows {
			zaplog.LoggerSugar.Debugf("%v\n", row)
		}
	}

}
