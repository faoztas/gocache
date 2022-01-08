package log

import "log"

var adapter LogInterface

func init() {
	adapter = NewLogger()
}

type LogInterface interface {
	Debug(msg error)
	Info(msg error)
	Error(msg error)
	Fatal(msg error)
	Warning(msg error)
	Http(msg error)
	GetApiLog() *log.Logger
}

func Error(msg error) {
	adapter.Error(msg)
}

func Debug(msg error) {
	adapter.Debug(msg)
}

func Info(msg error) {
	adapter.Info(msg)
}

func Warning(msg error) {
	adapter.Warning(msg)
}

func Fatal(msg error) {
	adapter.Fatal(msg)
}

func Http(msg error) {
	adapter.Http(msg)
}

func GetApiLog() *log.Logger {
	return adapter.GetApiLog()
}
