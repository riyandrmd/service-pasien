package poli-handler


import (
	"net/http"
	"fiture-poli/poli"
	"response/response"

	"github.com/gin-gonic/gin"
)

func PoliRoute(poliUc poli.PoliUseCase, r *gin.RouterGroup) {
	uc := poliHandler{
		poliUc,
	}

	v2 := r.Group("blog")
	v2.GET("", uc.GetAllPoli)
	v2.POST("", uc.CreatePoli)
	v2.PUT(":id", uc.UpdatePoli)
	v2.DELETE(":id", uc.DeletePoli)
}

type blogPoli struct {
	poliUc poli.PoliUseCase
}

func (poliHandler *poliHandler) GetAllPoli(c *gin.Context) {
	result := poliHandler.poliUc.GetAllBlog(c)
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

func (poliHandler *poliHandler) CreatePoli(c *gin.Context) {
	err := poliHandler.poliUc.CreatePoli(c)
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


func (poliHandler *poliHandler) UpdatePoli(c *gin.Context) {
	err := poliHandler.poliUc.UpdatePoli(c)
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
	err := poliHandler.poliUc.DeletePoli(c)
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