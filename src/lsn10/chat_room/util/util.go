package util

import (
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

func GetTodayDateTime() string {
	dateTime := time.Unix(time.Now().Unix()+8*3600, 0)
	return dateTime.Format("2006-01-02 15:04")
}

func GetFileLine() string {
	_, file, line, _ := runtime.Caller(1)
	return filepath.Base(file) + " " + strconv.Itoa(line) + " "
}
