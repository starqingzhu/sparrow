/*
 * @Author: bsun
 * @Date: 2024-07-01 14:06:53
 * @Last Modified by: bsun
 * @Last Modified time: 2024-07-01 15:16:02
 */

package stimer

import (
	"fmt"
	"sparrow/pkg/log/zaplog"
	"sync"
	"testing"
	"time"
)

// 触发函数
func foo(args ...interface{}) {
	fmt.Printf("I am No. %d function, delay %d ms\n", args[0].(int), args[1].(int))
}

// 手动创建调度器运转时间轮
func TestNewTimerScheduler(t *testing.T) {
	var timeCount = 10
	wg := sync.WaitGroup{}
	wg.Add(timeCount)

	timerScheduler := NewTimerScheduler()
	timerScheduler.Start()

	//在scheduler中添加timer
	for i := 1; i <= timeCount; i++ {
		f := NewDelayFunc(foo, []interface{}{i, i * 3})
		tID, err := timerScheduler.CreateTimerAfter(f, time.Duration(3*i)*time.Millisecond)
		if err != nil {
			zaplog.LoggerSugar.Error("create timer error", tID, err)
			break
		}
	}

	//执行调度器触发函数
	go func() {
		delayFuncChan := timerScheduler.GetTriggerChan()
		for df := range delayFuncChan {
			df.Call()
			wg.Done()
		}

	}()

	wg.Wait()
}

// 采用自动调度器运转时间轮
func TestNewAutoExecTimerScheduler(t *testing.T) {

	autoTS := NewAutoExecTimerScheduler()

	//给调度器添加Timer
	for i := 0; i < 10; i++ {
		f := NewDelayFunc(foo, []interface{}{i, i * 3})
		tID, err := autoTS.CreateTimerAfter(f, time.Duration(3*i)*time.Millisecond)
		if err != nil {
			zaplog.LoggerSugar.Error("create timer error", tID, err)
			break
		}
	}
	select {}
}
