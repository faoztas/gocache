package main

import (
	"fmt"
	"os"

	"github.com/faoztas/gocache/common/job"
	"github.com/faoztas/gocache/common/log"
	"github.com/faoztas/gocache/common/server"
	"github.com/faoztas/gocache/config"
	"github.com/faoztas/gocache/di"
	"github.com/faoztas/gocache/route"
	"github.com/faoztas/gocache/utils/errors"
)

func main() {
	env := config.GetEnvironment()
	listenAddr := fmt.Sprintf(":%s", os.Getenv("PORT"))

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
