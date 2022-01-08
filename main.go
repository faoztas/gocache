package main

import (
	"fmt"

	"gocache/common/job"
	"gocache/common/log"
	"gocache/common/server"
	"gocache/config"
	"gocache/di"
	"gocache/route"
	"gocache/utils/errors"
)

func main() {
	env := config.GetEnvironment()
	listenAddr := fmt.Sprintf("%s:%s", env.ApplicationHost, env.ApplicationPort)

	err := di.ContainerStorageService.Load()
	if err != nil {
		log.Fatal(err)
	}

	if env.StorageJob {
		job.Run(func() {
			if err = di.ContainerStorageService.Save(); err != nil {
				log.Error(err)
			}
		}, config.GetEnvironment().StorageSchedule)
	}

	router := di.ContainerRouter
	route.DefineAPIRouter(router)

	log.Http(fmt.Errorf(errors.ServerListen, listenAddr))
	server.ListenAndServe(listenAddr, router)
}
