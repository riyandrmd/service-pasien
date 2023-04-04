package handler

import (
	"administrasi/middleware"
	"administrasi/pasien"
	"administrasi/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PasienRoute(pasienUc pasien.PasienUseCase, v1 *gin.RouterGroup) {
	uc := pasienHandler{
		pasienUc,
	}

	v2 := v1.Group("pasien")
	v2.GET("", middleware.Auth(), uc.GetAllPasien)
	v2.GET(":id", middleware.Auth(), uc.GetDetail)
	v2.POST("", middleware.Auth(), uc.CreatePasien)
	v2.PUT(":id", middleware.Auth(), uc.UpdatePasien)
	v2.DELETE(":id", middleware.Auth(), uc.DeletePasien)

}

type pasienHandler struct {
	pasienUc pasien.PasienUseCase
}

func (pasienHandler *pasienHandler) GetAllPasien(c *gin.Context) {
	result, pagination, err := pasienHandler.pasienUc.GetAllPasienUC(c)
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

func (pasienHandler *pasienHandler) CreatePasien(c *gin.Context) {
	err := pasienHandler.pasienUc.CreatePasienUC(c)
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

func (pasienHandler *pasienHandler) GetDetail(c *gin.Context) {
	result, err := pasienHandler.pasienUc.GetDetailPasienUC(c)
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

func (pasienHandler *pasienHandler) UpdatePasien(c *gin.Context) {
	err := pasienHandler.pasienUc.UpdatePasienUC(c)
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

func (pasienHandler *pasienHandler) DeletePasien(c *gin.Context) {
	err := pasienHandler.pasienUc.DeletePasienUC(c)
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
