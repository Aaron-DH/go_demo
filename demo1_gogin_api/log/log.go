package log

import (
	"github.com/jeanphorn/log4go"
)

// 用于记录用户日志
var UserLog *log4go.Filter
// 用于记录系统运行日志
var RunLog *log4go.Filter

func init() {
	log4go.LoadConfiguration("./conf/log4go.json")
	UserLog = log4go.LOGGER("USER")
	RunLog = log4go.LOGGER("RUN")
}
