package redis

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func boolError(msg string) (bool, error) {
	return false, errors.New(msg)
}

func intError(msg string) (int, error) {
	return -1, errors.New(msg)
}

func int64Error(msg string) (int64, error) {
	return -1, errors.New(msg)
}

func stringError(msg string) (string, error) {
	return "", errors.New(msg)
}

func stringsError(msg string) ([]string, error) {
	return nil, errors.New(msg)
}

func mapError(msg string) (map[string]string, error) {
	return nil, errors.New(msg)
}

func sliceError(msg string) ([]interface{}, error) {
	return nil, errors.New(msg)
}

func float64Error(msg string) (float64, error) {
	return 0, errors.New(msg)
}

func toStringSlice(arr []interface{}) []string {
	if arr == nil {
		return make([]string, 0)
	}
	result := make([]string, len(arr))
	for i, v := range arr {
		if v == nil {
			result[i] = ""
			continue
		}
		result[i] = v.(string)
	}
	return result
}

/*
-------------------------------------------check----------------------------------------------
*/
func checkKey(key string) error {
	if key == "" {
		return errors.New("key cannot be an empty string")
	}
	return nil
}

func checkKeys(keys ...string) error {
	for _, k := range keys {
		if err := checkKey(k); err != nil {
			return err
		}
	}

	return nil
}

func checkParamString(mem string) error {
	if mem == "" {
		return errors.New("param cannot be an empty string")
	}
	return nil
}

func checkChannel(channel string) error {
	if channel == "" {
		return errors.New("channel or message cannot be empty string")
	}
	return nil
}

// 功能：检查函数个数是偶数，且不能为0
func checkParamEven(params ...interface{}) (int64, error) {
	pLen := int64(len(params))
	if pLen == 0 || pLen%2 != 0 {
		return int64Error(fmt.Sprintf("wrong number of params"))
	}
	return pLen, nil
}

func checkNil(err error) error {
	if err != nil && err != redis.ErrNil {
		return err
	}
	return nil
}

func checkOKResponse(cmd, rtn string, err error) error {
	if err = checkNil(err); err != nil {
		return err
	} else if rtn != "OK" {
		return fmt.Errorf("%s did not get OK response: %s", cmd, rtn)
	}
	return nil
}
