package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	dokterHandler "administrasi/dokter/handler"
	dokterRepo "administrasi/dokter/repository"
	dokterUc "administrasi/dokter/usecase"

	kamarHandler "administrasi/kamar/handler"
	kamarRepo "administrasi/kamar/repository"
	kamarUc "administrasi/kamar/usecase"

	"administrasi/middleware"
	pasienHandler "administrasi/pasien/handler"
	pasienRepo "administrasi/pasien/repository"
	pasienUc "administrasi/pasien/usecase"

	poliHandler "administrasi/poli/handler"
	poliRepo "administrasi/poli/repository"
	poliUc "administrasi/poli/usecase"

	rekdisHandler "administrasi/rekammedis/handler"
	rekdisRepo "administrasi/rekammedis/repository"
	rekdisUc "administrasi/rekammedis/usecase"

	obatHandler "administrasi/obat/handler"
	obatRepo "administrasi/obat/repository"
	obatUc "administrasi/obat/usecase"
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
	rekdisRepo := rekdisRepo.NewRekamMedisRepo(h.DB)
	dokterRepo := dokterRepo.NewDokterRepo(h.DB)
	kamarRepo := kamarRepo.NewKamarRepo(h.DB)
	obatRepo := obatRepo.NewObatRepo(h.DB)

	PasienUseCase := pasienUc.NewPasienUseCase(pasienRepo, h.Redis)
	PoliUseCase := poliUc.NewPoliUseCase(poliRepo, h.Redis)
	RekdisUseCase := rekdisUc.NewNewRekamMedisUC(rekdisRepo, h.Redis)
	dokterUc := dokterUc.NewDokterUC(dokterRepo, h.Redis)
	kamarUc := kamarUc.NewKamarUseCase(kamarRepo, h.Redis)
	obatUc := obatUc.NewObatUseCase(obatRepo, h.Redis)

	middleware.Add(h.R, middleware.CORSMiddleware())

	v1 := h.R.Group("api")
	pasienHandler.PasienRoute(PasienUseCase, v1)
	poliHandler.PoliRoute(PoliUseCase, v1)
	rekdisHandler.RekamMedisRoute(RekdisUseCase, v1)
	dokterHandler.DokterRoute(dokterUc, v1)
	kamarHandler.KamarRoute(kamarUc, v1)
	obatHandler.ObatRoute(obatUc, v1)

}
