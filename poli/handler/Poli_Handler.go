package handler

import (
	"administrasi/middleware"
	"administrasi/poli"
	"administrasi/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PoliRoute(PoliUc poli.PoliUseCase, v1 *gin.RouterGroup) {
	uc := poliHandler{
		PoliUc,
	}

	v2 := v1.Group("poli")
	v2.GET("", middleware.Auth(), uc.GetAllPoli)
	v2.GET(":id", middleware.Auth(), uc.GetDetail)
	v2.POST("", middleware.Auth(), uc.CreatePoli)
	v2.PUT(":id", middleware.Auth(), uc.UpdatePoli)
	v2.DELETE(":id", middleware.Auth(), uc.DeletePoli)

}

type poliHandler struct {
	PoliUC poli.PoliUseCase
}

func (poliHandler *poliHandler) GetAllPoli(c *gin.Context) {
	result, pagination, err := poliHandler.PoliUC.GetAllPoliUC(c)
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

func (poliHandler *poliHandler) CreatePoli(c *gin.Context) {
	err := poliHandler.PoliUC.CreatePoliUC(c)
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

func (poliHandler *poliHandler) GetDetail(c *gin.Context) {
	result, err := poliHandler.PoliUC.GetDetailPoliUC(c)
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

func (poliHandler *poliHandler) UpdatePoli(c *gin.Context) {
	err := poliHandler.PoliUC.UpdatePoliUC(c)
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

func (poliHandler *poliHandler) DeletePoli(c *gin.Context) {
	err := poliHandler.PoliUC.DeletePoliUC(c)
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
