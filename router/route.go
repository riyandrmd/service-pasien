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

	poliHandler "administrasi/poli/handler"
	poliRepo "administrasi/poli/repository"
	poliUc "administrasi/poli/usecase"
)

type Handlers struct {
	Ctx   context.Context
	DB    *gorm.DB
	R     *gin.Engine
	Redis *redis.Client
}

func (h *Handlers) Routes() {
	pasienRepo := pasienRepo.NewPasienRepo(h.DB)
	poliRepo := poliRepo.NewPoliRepo(h.DB)

	PasienUseCase := pasienUc.NewPasienUseCase(pasienRepo, h.Redis)
	PoliUseCase := poliUc.NewPoliUseCase(poliRepo, h.Redis)

	middleware.Add(h.R, middleware.CORSMiddleware())

	v1 := h.R.Group("api")
	pasienHandler.PasienRoute(PasienUseCase, v1)
	poliHandler.PoliRoute(PoliUseCase, v1)
}
