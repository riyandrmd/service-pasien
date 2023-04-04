package handler

import (
	"administrasi/kamar"
	"administrasi/middleware"
	"administrasi/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func KamarRoute(KamarUc kamar.KamarUseCase, v1 *gin.RouterGroup) {
	uc := KmHandler{
		KamarUc,
	}

	v2 := v1.Group("kamar")

	v2.GET("", middleware.Auth(), uc.GetAllKamar)
	v2.GET(":id", middleware.Auth(), uc.GetDetailKamar)
	v2.POST("", middleware.Auth(), uc.CreateKamar)
	v2.PUT(":id", middleware.Auth(), uc.UpdateKamar)
	v2.DELETE(":id", middleware.Auth(), uc.DeleteKamar)

}

type KmHandler struct {
	KamarUc kamar.KamarUseCase
}

func (KmHandler *KmHandler) GetAllKamar(c *gin.Context) {
	result, pagination, err := KmHandler.KamarUc.GetAllKamarUC(c)
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

func (KmHandler *KmHandler) CreateKamar(c *gin.Context) {
	err := KmHandler.KamarUc.CreateKamarUC(c)
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

func (KmHandler *KmHandler) GetDetailKamar(c *gin.Context) {
	result, err := KmHandler.KamarUc.GetDetailKamarUC(c)
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

func (KmHandler *KmHandler) UpdateKamar(c *gin.Context) {
	err := KmHandler.KamarUc.UpdateKamarUC(c)
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

func (KmHandler *KmHandler) DeleteKamar(c *gin.Context) {
	err := KmHandler.KamarUc.DeleteKamarUC(c)
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
