package handler

import (
	"administrasi/middleware"
	"administrasi/obat"
	"administrasi/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ObatRoute(obatUc obat.ObatUseCase, v1 *gin.RouterGroup) {
	uc := obatHandler{
		obatUc,
	}

	v2 := v1.Group("obat")
	v2.GET("", middleware.Auth(), uc.GetAllObat)
	v2.GET(":id", middleware.Auth(), uc.GetDetail)
	v2.POST("", middleware.Auth(), uc.CreateObat)
	v2.PUT(":id", middleware.Auth(), uc.UpdateObat)
	v2.DELETE(":id", middleware.Auth(), uc.DeleteObat)

}

type obatHandler struct {
	obatUc pasien.ObatUseCase
}


func (obatHandler *obatHandler) GetAllObat(c *gin.Context) {
	result, pagination, err := obatHandler.obatUc.GetAllObatUC(c)
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

func (obatHandler *obatHandler) CreateObat(c *gin.Context) {
	err := obatHandler.obatUc.CreateObatUC(c)
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

func (obatHandler *obatHandler) GetDetail(c *gin.Context) {
	result, err := obatHandler.obatUc.GetDetailObatUC(c)
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

func (obatHandler *obatHandler) UpdateObat(c *gin.Context) {
	err := obatHandler.obatUc.UpdateObatUC(c)
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

func (obatHandler *obatHandler) DeleteObat(c *gin.Context) {
	err := obatHandler.obatUc.DeleteObatUC(c)
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
