package server

import (
	"database-api/app/core/services.abstractions/services"
	"database-api/app/infrastructure/presentation/webservices"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitEndpoints(s services.IServiceManager) {
	r := gin.Default()
	webService := webservices.NewWebservice(s)
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8081"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	}))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	r.POST("carpart/add", webService.AddPart)
	r.POST("add/parts", webService.AddParts)
	r.GET("carpart", webService.GetAllParts)

	r.GET("detail", webService.PartDetails)

	r.POST("part-photo/add", webService.AddPartPhoto)

	r.Run()
}
