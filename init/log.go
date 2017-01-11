package init

import (
	. "goERP/utils"

	"github.com/astaxie/beego/logs"
)

func ConfigLog() {
	Log = logs.NewLogger()
	Log.Async()
	Log.EnableFuncCallDepth(true)

}

// Logs.SetLogger(logs.AdapterMultiFile, `{"filename":"test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
