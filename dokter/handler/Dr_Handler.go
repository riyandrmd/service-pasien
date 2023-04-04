package handler

import (
	"administrasi/dokter"
	"administrasi/middleware"
	"administrasi/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DokterRoute(DokterUc dokter.DokterUseCase, v1 *gin.RouterGroup) {
	uc := DrHandler{
		DokterUc,
	}

	v2 := v1.Group("dokter")

	v2.GET("", middleware.Auth(), uc.GetAllDokter)
	v2.GET(":id", middleware.Auth(), uc.GetDetailDokter)
	v2.POST("", middleware.Auth(), uc.CreateDokter)
	v2.PUT(":id", middleware.Auth(), uc.UpdateDokter)
	v2.DELETE(":id", middleware.Auth(), uc.DeleteDokter)

}

type DrHandler struct {
	DokterUc dokter.DokterUseCase
}

func (DrHandler *DrHandler) GetAllDokter(c *gin.Context) {
	result, pagination, err := DrHandler.DokterUc.GetAllDokterUC(c)
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

func (DrHandler *DrHandler) CreateDokter(c *gin.Context) {
	err := DrHandler.DokterUc.CreateDokterUC(c)
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

func (DrHandler *DrHandler) GetDetailDokter(c *gin.Context) {
	result, err := DrHandler.DokterUc.GetDetailDokterUC(c)
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

func (DrHandler *DrHandler) UpdateDokter(c *gin.Context) {
	err := DrHandler.DokterUc.UpdateDokterUC(c)
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

func (DrHandler *DrHandler) DeleteDokter(c *gin.Context) {
	err := DrHandler.DokterUc.DeleteDokterUC(c)
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
