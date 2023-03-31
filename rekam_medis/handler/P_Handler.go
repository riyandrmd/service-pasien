package handler

import (
	"administrasi/middleware"
	"administrasi/rekam_medis"
	"administrasi/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RekamRoute(rekamUc rekam_medis.RekamUseCase, v1 *gin.RouterGroup) {
	uc := rekamHandler{
		rekamUc,
	}

	v2 := v1.Group("rekammedis")
	v2.GET("", middleware.Auth(), uc.GetAllRekamMedis)
	v2.GET(":id", middleware.Auth(), uc.GetDetailRekamMedis)
	v2.POST("", middleware.Auth(), uc.CreateRekamMedis)
	v2.PUT(":id", middleware.Auth(), uc.UpdateRekamMedis)
	v2.DELETE(":id", middleware.Auth(), uc.DeleteRekamMedis)

}

type rekamHandler struct {
	rekamUc rekam_medis.RekamUseCase
}

func (rekamHandler *rekamHandler) GetAllRekamMedis(c *gin.Context) {
	result, pagination, err := rekamHandler.rekamUc.GetAllRekamMedisUC(c)
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

func (rekamHandler *rekamHandler) CreateRekamMedis(c *gin.Context) {
	err := rekamHandler.rekamUc.CreateRekamMedisUC(c)
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

func (rekamHandler *rekamHandler) GetDetailRekamMedis(c *gin.Context) {
	result, err := rekamHandler.rekamUc.GetDetailRekamMedisUC(c)
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

func (rekamHandler *rekamHandler) UpdateRekamMedis(c *gin.Context) {
	err := rekamHandler.rekamUc.UpdateRekamMedisUC(c)
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

func (rekamHandler *rekamHandler) DeleteRekamMedis(c *gin.Context) {
	err := rekamHandler.rekamUc.DeleteRekamMedisUC(c)
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
