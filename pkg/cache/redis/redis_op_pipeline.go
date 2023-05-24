package redis

import (
	"errors"
	"fmt"
	"log"
)

func (r *Redis) Pipeline(cmdList []interface{}) error {
	conn := r.pool.Get()
	defer conn.Close()

	if len(cmdList) == 0 {
		return errors.New("cmdList is empty")
	}
	conn.Send("MULTI")
	for i, v := range cmdList {
		paramList, ok := v.([]interface{})
		if !ok {
			return errors.New(fmt.Sprintf("paramList:%v, is not slice", paramList))
		}
		if len(paramList) == 0 {
			return errors.New(fmt.Sprintf("%d paramList is empty", i))
		}

		cmd, cmdOk := paramList[0].(string)
		if !cmdOk {
			return errors.New(fmt.Sprintf("cmd:%v is not string", paramList[0]))
		}

		if err := conn.Send(cmd, paramList[1:]...); checkNil(err) != nil {
			return errors.New(fmt.Sprintf("cmd:%v send failed, err:%s", paramList, err.Error()))
		}
		log.Printf("Pipeline [%v] Send success\n", paramList)
	}

	reply, err := conn.Do("EXEC")
	log.Printf("Pipeline cmdList[%v] err:%v, ret:%v", cmdList, err, reply)
	if err != nil {
		return err
	}
	//err := conn.Flush()
	//if checkNil(err) != nil {
	//	return errors.New(fmt.Sprintf("Pipeline Flush failed, err:%v", err))
	//}

	return err
}
