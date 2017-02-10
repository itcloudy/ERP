package utils

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/session"
)

//LangType 语言
type LangType struct {
	Lang, Name string
}

var DefaultPageLimit = 20

//GlobalSessions 全局session
var GlobalSessions *session.Manager

// consoleLogs开发模式下日志
var consoleLogs *logs.BeeLogger

// fileLogs 生产环境下日志
var fileLogs *logs.BeeLogger
var runmode string

func init() {
	consoleLogs = logs.NewLogger(1)
	consoleLogs.SetLogger(logs.AdapterConsole)
	fileLogs = logs.NewLogger(10000)
	fileLogs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/goERP.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"],"level":7,"daily":true,"maxdays":10}`)
	runmode = strings.TrimSpace(strings.ToLower(beego.AppConfig.String("runmode")))
	if runmode == "" {
		runmode = "dev"
	}
}

//LogOut 输出日志
// @Title LogOut
// @Param	body		body 	models.AccountAccountTag	true		"body for AccountAccountTag content"
func LogOut(level, v interface{}) {
	format := "%s"
	if level == "" {
		level = "debug"
	}
	if runmode == "dev" {
		switch level {
		case "emergency":
			fileLogs.Emergency(format, v)
		case "alert":
			fileLogs.Alert(format, v)
		case "critical":
			fileLogs.Critical(format, v)
		case "error":
			fileLogs.Error(format, v)
		case "warning":
			fileLogs.Warning(format, v)
		case "notice":
			fileLogs.Notice(format, v)
		case "informational":
			fileLogs.Informational(format, v)
		case "debug":
			fileLogs.Debug(format, v)
		case "warn":
			fileLogs.Warn(format, v)
		case "info":
			fileLogs.Info(format, v)
		case "trace":
			fileLogs.Trace(format, v)
		default:
			fileLogs.Debug(format, v)
		}
	}
	switch level {
	case "emergency":
		consoleLogs.Emergency(format, v)
	case "alert":
		consoleLogs.Alert(format, v)
	case "critical":
		consoleLogs.Critical(format, v)
	case "error":
		consoleLogs.Error(format, v)
	case "warning":
		consoleLogs.Warning(format, v)
	case "notice":
		consoleLogs.Notice(format, v)
	case "informational":
		fileLogs.Informational(format, v)
	case "debug":
		consoleLogs.Debug(format, v)
	case "warn":
		consoleLogs.Warn(format, v)
	case "info":
		consoleLogs.Info(format, v)
	case "trace":
		consoleLogs.Trace(format, v)
	default:
		consoleLogs.Debug(format, v)
	}

}
