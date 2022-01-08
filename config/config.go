package config

import (
	"encoding/json"
	"os"
	"time"

	"github.com/faoztas/gocache/utils"
)

var env *environment

type environment struct {
	ApplicationHost string
	ApplicationPort string
	FilePath        string
	Timeout         time.Duration
	Trace           bool
	Silent          bool
	StorageJob      bool
	StorageSchedule time.Duration
}

func init() {
	file, err := os.ReadFile("env.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &env)
	if err != nil {
		panic(err)
	}

	if env.FilePath == "" {
		env.FilePath = utils.GenerateFilePath()
	}

	env.StorageSchedule = env.StorageSchedule * time.Minute
	env.Timeout = env.Timeout * time.Second
}

func GetEnvironment() *environment {
	return env
}
