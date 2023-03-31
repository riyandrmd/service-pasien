package handlers

import (
	"net/http"
	"administrasi/kamar"
	"github.com/gin-gonic/gin"
	"administrasi/response"
	"administrasi/middleware"
)

func KamarRoute(kamarUc kamar.KamarUseCase, v1 *gin.RouterGroup) {
	uc := kamarHandler{
		kamarUc,
	}

	v2 := r.Group("kamar")
	v2.GET("",middleware.Auth(), uc.GetAllKamar)
	v2.GET(":id", middleware.Auth(), uc.GetDetail)
	v2.POST("",middleware.Auth(), uc.CreateKamar)
	v2.PUT(":id",middleware.Auth(), uc.UpdateKamar)
	v2.DELETE(":id",middleware.Auth(), uc.DeleteKamar)

}

type kamarHandler struct {
	kamarUc kamar.KamarUseCase
}

func (kamarHandler *kamarHandler) GetAllKamar(c *gin.Context) {
	result,pagination, err := kamarHandler.kamarUc.GetAllKamarUC(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":err,
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		Data: result,
		Status: "success get all data",
		Message: "success get all data",
	})
}

func (kamarHandler *kamarHandler) CreateKamar(c *gin.Context) {
	err := kamarHandler.kamarUc.CreateKamarUC(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		Status:  "success",
		Message: "success",
	})
}

func (kamarHandler *kamarHandler) GetDetail(c *gin.Context) {
	result, err := kamarHandler.kamarUc.GetDetailKamarUC(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Data:    result,
		Status:  "success",
		Message: "success",
	})
}

func (kamarHandler *kamarHandler) UpdateKamar(c *gin.Context) {
	err := kamarHandler.kamarUc.UpdateKamarUC(c)
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

func (kamarHandler *kamarHandler) DeleteKamar(c *gin.Context) {
	err := kamarHandler.kamarUc.DeleteKamarUC(c)
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
