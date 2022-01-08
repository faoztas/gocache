package di

import (
	"gocache/common/router"
	"gocache/config"
	"gocache/controller"
	"gocache/repository"
	"gocache/service"
)

var (
	ContainerRouter = router.NewRouter()

	ContainerStorageRepository = repository.NewStorageRepository(&config.GetEnvironment().FilePath)
	ContainerStorageService    = service.NewStorageService(&ContainerStorageRepository, &ContainerCacheRepository)

	ContainerCacheRepository = repository.NewCacheRepository()
	ContainerCacheService    = service.NewCacheService(&ContainerCacheRepository)
	ContainerCacheController = controller.NewCacheController(&ContainerRouter, &ContainerCacheService)
)
