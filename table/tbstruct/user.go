package tbstruct

import (
	"encoding/json"
	"sparrow/pkg/log/zaplog"
)

type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func (u *User) String() string {
	tb, err := json.Marshal(u)
	if err != nil {
		zaplog.LoggerSugar.Errorf("table User marshal error, err:%s, tb:%v", err.Error(), *u)
		return ""
	}

	return string(tb)
}
