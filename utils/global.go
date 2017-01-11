package utils

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/session"
)

//语言
type LangType struct {
	Lang, Name string
}

//日志
var Log *logs.BeeLogger

//全局session
var GlobalSessions *session.Manager
