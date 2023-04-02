package doctor-handler

//try

import (
	"net/http"
	"fiture-doctor/doctor"
	"response/response"

	"github.com/gin-gonic/gin"
)

func DoctorRoute(doctorUc doctor.DoctorUseCase, r *gin.RouterGroup) {
	uc := doctorHandler{
		doctorUc,
	}

	v2 := r.Group("doctor")
	v2.GET("", uc.GetAllDoctor)
	v2.POST("", uc.CreateDoctor)
}

type doctorHandler struct {
	doctorUc doctor.DoctorUseCase
}

func (doctorHandler *doctorHandler) GetAllDoctor(c *gin.Context) {
	result := doctorHandler.doctorUc.GetAllDoctor(c)
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

func (doctorHandler *doctorHandler) CreateDoctor(c *gin.Context) {
	err := doctorHandler.doctorUc.CreateDoctor(c)
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

