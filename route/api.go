package route

import (
	"github.com/faoztas/gocache/common/router"
	"github.com/faoztas/gocache/di"
)

func DefineAPIRouter(router router.Router) {
	cacheController := di.ContainerCacheController

	router.GET("/get", cacheController.Get)
	router.POST("/set", cacheController.Set)
	router.DELETE("/delete", cacheController.Delete)
	router.DELETE("/flush", cacheController.Flush)
}
