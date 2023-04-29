package webservices

import (
	"database-api/app/core/contracts/dtos/carPartDtos"
	"database-api/app/core/contracts/dtos/partPhotoDtos"
	"database-api/app/core/services.abstractions/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type webService struct {
	Services services.IServiceManager
}

func NewWebservice(manager services.IServiceManager) *webService {
	return &webService{Services: manager}
}

func (w *webService) AddPart(c *gin.Context) {
	var part *carPartDtos.CarPartDto

	err := c.BindJSON(&part)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	res := w.Services.AddPart(part)
	c.JSON(http.StatusOK, res)
	return
}

func (w *webService) GetAllParts(c *gin.Context) {
	res := w.Services.GetAll()
	c.JSON(http.StatusOK, res)
	return
}

func (w *webService) AddPartPhoto(c *gin.Context) {
	var photoDto *partPhotoDtos.PartPhotoDto

	err := c.BindJSON(&photoDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	res := w.Services.AddPhoto(photoDto)
	c.JSON(http.StatusOK, res)
	return
}
