package logic

import (
	"fmt"
	"time"
)

var (
	logicStop bool
	logicEnd  bool
)

func init() {
	logicStop = false
	logicEnd = false
}

func LogicStart() {
	for !logicStop {
		time.Sleep(time.Duration(3) * time.Second)
		fmt.Println("logicMain runing...")
	}
	logicEnd = true
}

func LogicEnd() {
	logicStop = true
	// 等待任务队列结束
	for !logicEnd {
	}
}
