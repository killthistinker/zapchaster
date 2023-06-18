package webservices

import (
	"database-api/app/core/contracts/dtos/carPartDtos"
	"database-api/app/core/contracts/dtos/partPhotoDtos"
	"database-api/app/core/contracts/helpers/formdata"
	"database-api/app/core/contracts/mapper"
	"database-api/app/core/services.abstractions/services"
	"github.com/gin-gonic/gin"
	"mime/multipart"
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
	var photoDto []partPhotoDtos.PartPhotoDto
	var fileHeaders []multipart.FileHeader
	form, err := c.MultipartForm()
	files := form.File["images"]
	for _, file := range files {
		fileHeaders = append(fileHeaders, *file)
	}

	base64stringImg, err := formdata.MapToBase64(&fileHeaders)
	data := form.Value["photos"]
	err = formdata.MapToCarParts(data, &photoDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	for i := 0; i < len(photoDto); i++ {
		photoDto[i].PartPhoto = base64stringImg[i]
	}

	res := w.Services.AddPhoto(photoDto)
	c.JSON(http.StatusOK, res)
	return
}

func (w *webService) PartDetails(c *gin.Context) {
	var partDetail carPartDtos.PartDetailDto
	pId := c.Query("partId")
	if pId == "" {
		c.JSON(http.StatusBadRequest, pId)
		return
	}

	partRes, err := w.Services.GetPart(pId)
	photoRes, err := w.Services.GetPartPhotosFromId(pId)
	if err != nil {
		c.JSON(http.StatusBadRequest, pId)
		return
	}
	partDetail = mapper.MapToPartDetail(partRes, photoRes)
	c.JSON(http.StatusOK, partDetail)
	return
}

func (w *webService) AddParts(c *gin.Context) {
	var parts []carPartDtos.CarPartDto
	var fileHeaders []multipart.FileHeader
	form, err := c.MultipartForm()

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	files := form.File["images"]
	for _, file := range files {
		fileHeaders = append(fileHeaders, *file)
	}

	base64stringImg, err := formdata.MapToBase64(&fileHeaders)

	data := form.Value["parts"]
	err = formdata.MapToCarParts(data, &parts)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	for i := 0; i < len(parts); i++ {
		parts[i].MainPhoto = base64stringImg[i]
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res := w.Services.AddParts(&parts)
	if res != 0 {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
