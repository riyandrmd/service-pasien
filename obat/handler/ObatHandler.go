package handler

import (
	"administrasi/obat"
	"administrasi/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ObatRoute(ObatUc obat.ObatUseCase, v1 *gin.RouterGroup) {
	uc := ObatHandler{
		ObatUc,
	}

	v2 := v1.Group("obat")

	v2.GET("", uc.GetAllObat)
	v2.GET(":id", uc.GetDetailObat)
	v2.POST("", uc.CreateObat)
	v2.PUT(":id", uc.UpdateObat)
	v2.DELETE(":id", uc.DeleteObat)

}

type ObatHandler struct {
	ObatUc obat.ObatUseCase
}

func (ObatHandler *ObatHandler) GetAllObat(c *gin.Context) {
	result, pagination, err := ObatHandler.ObatUc.GetAllObatUC(c)
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

func (ObatHandler *ObatHandler) CreateObat(c *gin.Context) {
	err := ObatHandler.ObatUc.CreateObatUC(c)
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

func (ObatHandler *ObatHandler) GetDetailObat(c *gin.Context) {
	result, err := ObatHandler.ObatUc.GetDetailObatUC(c)
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

func (ObatHandler *ObatHandler) UpdateObat(c *gin.Context) {
	err := ObatHandler.ObatUc.UpdateObatUC(c)
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

func (ObatHandler *ObatHandler) DeleteObat(c *gin.Context) {
	err := ObatHandler.ObatUc.DeleteObatUC(c)
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
