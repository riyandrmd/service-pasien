package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"administrasi/middleware"
	rekamHandler "administrasi/rekam_medis/handler"
	rekamRepo "administrasi/rekam_medis/repository"
	rekamUc "administrasi/rekam_medis/usecase"
)

type Handlers struct {
	Ctx   context.Context
	DB    *gorm.DB
	R     *gin.Engine
	Redis *redis.Client
}

func (h *Handlers) Routes() {
	RekamRepo := rekamRepo.NewRekamRepo(h.DB)
	RekamUseCase := rekamUc.NewRekamUseCase(RekamRepo, h.Redis)

	middleware.Add(h.R, middleware.CORSMiddleware())

	v1 := h.R.Group("api")
	rekamHandler.RekamRoute(RekamUseCase, v1)
}
