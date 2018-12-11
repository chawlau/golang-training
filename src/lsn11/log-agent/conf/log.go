package conf

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

func ConvertLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}

func InitLogger() (err error) {
	config := make(map[string]interface{})
	config["fileName"] = AppConfig.LogPath
	config["level"] = logs.LevelDebug

	configStr, err := json.Marshal(config)

	if err != nil {
		fmt.Println("marshal failed err ", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))
	logs.SetLogFuncCall(true)
	logs.SetLogFuncCallDepth(3)
	return
}
