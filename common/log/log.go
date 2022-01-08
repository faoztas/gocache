package log

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"time"

	"gocache/config"
	"gocache/utils"
)

type logger struct {
	api *log.Logger
	log *log.Logger
}

func NewLogger() *logger {
	return &logger{
		api: log.New(os.Stdout, "http: ", log.Ldate|log.Lmicroseconds|log.LUTC),
		log: log.New(os.Stdout, "log: ", log.Lmsgprefix),
	}
}

func (s *logger) set(msg string, level string) interface{} {
	p, file, line, _ := runtime.Caller(4)
	data := map[string]interface{}{
		"time":      time.Now().Format(utils.TimeFormat),
		"timestamp": time.Now().Unix(),
		"level":     level,
		"class":     fmt.Sprintf("file: %s - line: %d", file, line),
		"function":  runtime.FuncForPC(p).Name(),
		"data":      msg,
	}

	if config.GetEnvironment().Trace {
		data["trace"] = string(debug.Stack())
	}

	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func (s *logger) Error(msg error) {
	if !config.GetEnvironment().Silent {
		s.log.Println(s.set(msg.Error(), "ERROR"))
	}
}

func (s *logger) Debug(msg error) {
	if !config.GetEnvironment().Silent {
		s.log.Println(s.set(msg.Error(), "DEBUG"))
	}
}

func (s *logger) Info(msg error) {
	if !config.GetEnvironment().Silent {
		s.log.Println(s.set(msg.Error(), "INFO"))
	}
}

func (s *logger) Fatal(msg error) {
	s.log.Fatal(s.set(msg.Error(), "FATAL"))
}

func (s *logger) Warning(msg error) {
	if !config.GetEnvironment().Silent {
		s.log.Println(s.set(msg.Error(), "WARNING"))
	}
}

func (s *logger) Http(msg error) {
	s.api.Println(msg.Error())
}

func (s *logger) GetApiLog() *log.Logger {
	return s.api
}
