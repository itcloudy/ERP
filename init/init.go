package init

import (
	"errors"

	. "goERP/utils"
	"path/filepath"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/howeyc/fsnotify"
)

var LangTypes []*LangType // Languages are supported.

func initLocales() {
	// Initialized language type list.
	langs := strings.Split(beego.AppConfig.String("lang::types"), "|")
	names := strings.Split(beego.AppConfig.String("lang::names"), "|")
	LangTypes = make([]*LangType, 0, len(langs))
	for i, v := range langs {
		LangTypes = append(LangTypes, &LangType{
			Lang: v,
			Name: names[i],
		})
	}

	for _, lang := range langs {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file: " + err.Error())
			return
		}
	}
}
func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

func loadtimes(t time.Time) int {
	return int(time.Now().Sub(t).Nanoseconds() / 1e6)
}

func initTemplates() {
	beego.AddFuncMap("dict", dict)
	beego.AddFuncMap("loadtimes", loadtimes)
}

func InitApp() {
	initTemplates()
	initLocales()
	// settingCompress()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic("Failed start app watcher: " + err.Error())
	}

	go func() {
		for {
			select {
			case event := <-watcher.Event:
				switch filepath.Ext(event.Name) {
				case ".ini":
					beego.Info(event)

					if err := i18n.ReloadLangs(); err != nil {
						beego.Error("Conf Reload: ", err)
					}

					beego.Info("Config Reloaded")

					// case ".json":
					// 	if event.Name == CompressConfPath {
					// 		settingCompress()
					// 		beego.Info("Beego Compress Reloaded")
					// 	}
				}
			}
		}
	}()

	if err := watcher.WatchFlags("conf", fsnotify.FSN_MODIFY); err != nil {
		beego.Error(err)
	}
}
