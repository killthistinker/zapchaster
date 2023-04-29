package server

import (
	"database-api/app/core/services.abstractions/services"
	"database-api/app/infrastructure/presentation/webservices"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitEndpoints(s services.IServiceManager) {
	r := gin.Default()
	webService := webservices.NewWebservice(s)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	r.POST("carpart/add", webService.AddPart)
	r.GET("carpart", webService.GetAllParts)
	r.POST("carphoto/add", webService.AddPartPhoto)

	r.Run()
}
