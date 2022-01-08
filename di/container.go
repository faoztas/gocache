package di

import (
	"github.com/faoztas/gocache/common/router"
	"github.com/faoztas/gocache/config"
	"github.com/faoztas/gocache/controller"
	"github.com/faoztas/gocache/repository"
	"github.com/faoztas/gocache/service"
)

var (
	ContainerRouter = router.NewRouter()

	ContainerStorageRepository = repository.NewStorageRepository(&config.GetEnvironment().FilePath)
	ContainerStorageService    = service.NewStorageService(&ContainerStorageRepository, &ContainerCacheRepository)

	ContainerCacheRepository = repository.NewCacheRepository()
	ContainerCacheService    = service.NewCacheService(&ContainerCacheRepository)
	ContainerCacheController = controller.NewCacheController(&ContainerRouter, &ContainerCacheService)
)
