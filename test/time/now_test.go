package time

import (
	"github.com/jinzhu/now"
	"testing"
)

func TestNow(t *testing.T) {
	tm, err := now.Parse("2023-11-30 17:00:00")
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("time:%d, date:%s", tm.UnixMilli(), tm.String())

	day := now.BeginningOfDay()
	t.Logf("day:%s", day.String())

	minute := now.BeginningOfMinute()
	t.Logf("minute:%s", minute.String())

	//now.Between()
}
