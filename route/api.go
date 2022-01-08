package route

import (
	"gocache/common/router"
	"gocache/di"
)

func DefineAPIRouter(router router.Router) {
	cacheController := di.ContainerCacheController

	router.GET("/get", cacheController.Get)
	router.POST("/set", cacheController.Set)
	router.DELETE("/delete", cacheController.Delete)
	router.DELETE("/flush", cacheController.Flush)
}
