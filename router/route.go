package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"administrasi/middleware"
	pasienHandler "administrasi/pasien/handler"
	pasienRepo "administrasi/pasien/repository"
	pasienUc "administrasi/pasien/usecase"
)

type Handlers struct {
	Ctx   context.Context
	DB    *gorm.DB
	R     *gin.Engine
	Redis *redis.Client
}

func (h *Handlers) Routes() {
	pasienRepo := pasienRepo.NewPasienRepo(h.DB)
	PasienUseCase := pasienUc.NewPasienUseCase(pasienRepo, h.Redis)

	middleware.Add(h.R, middleware.CORSMiddleware())

	v1 := h.R.Group("api")
	pasienHandler.PasienRoute(PasienUseCase, v1)
}
