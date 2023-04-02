package handler

import (
	"administrasi/middleware"
	"administrasi/rekammedis"
	"administrasi/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RekamMedisRoute(rekdisUc rekammedis.RekamMedisUseCase, v1 *gin.RouterGroup) {
	uc := RekDisHandler{
		rekdisUc,
	}

	v2 := v1.Group("rekammedis")

	v2.GET("", middleware.Auth(), uc.GetAllRekamMedis)
	v2.GET(":id", middleware.Auth(), uc.GetDetailRekamMedis)
	v2.POST("", middleware.Auth(), uc.CreateRekamMedis)
	v2.PUT(":id", middleware.Auth(), uc.UpdateRekamMedis)
	v2.DELETE(":id", middleware.Auth(), uc.DeleteRekamMedis)

}

type RekDisHandler struct {
	rekdisUc rekammedis.RekamMedisUseCase
}

func (RekDisHandler *RekDisHandler) GetAllRekamMedis(c *gin.Context) {
	result, pagination, err := RekDisHandler.rekdisUc.GetAllRekamMedisUC(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Data:    result,
		Status:  "succes get all data",
		Message: "succes get all data",
		Meta:    pagination,
	})

}

func (RekDisHandler *RekDisHandler) CreateRekamMedis(c *gin.Context) {
	err := RekDisHandler.rekdisUc.CreateRekamMedisUC(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		Status:  "succes",
		Message: "succes",
	})
}

func (RekDisHandler *RekDisHandler) GetDetailRekamMedis(c *gin.Context) {
	result, err := RekDisHandler.rekdisUc.GetDetailRekamMedisUC(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Data:    result,
		Status:  "Success ",
		Message: "success",
	})
}

func (RekDisHandler *RekDisHandler) UpdateRekamMedis(c *gin.Context) {
	err := RekDisHandler.rekdisUc.UpdateRekamMedisUC(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{
		Status:  "success",
		Message: "success",
	})
}

func (RekDisHandler *RekDisHandler) DeleteRekamMedis(c *gin.Context) {
	err := RekDisHandler.rekdisUc.DeleteRekamMedisUC(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, response.Response{
		Status:  "success",
		Message: "success",
	})
}
