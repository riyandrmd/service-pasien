package rekam_medis

import (
	"administrasi/models"

	"github.com/gin-gonic/gin"
)

type Rekam_medis struct {
	Id_RekamMedis uint   `gorm:"primaryKey;autoIncreament;" json:"id_rekammedis"`
	Id_pasien     uint   `json:"id_pasien"`
	Diagnosa      string `json:"diagnosa"`
	Id_dokter     uint   `json:"id_dokter"`
	Tgl_periksa   string `json:"tgl_periksa"`
	Rawat_Inap    int    `json:"rawat_inap"`
	Id_kamar      uint   `json:"id_kamar"`
	Id_obat       uint   `json:"id_obat"`
}

type RekamRepo interface {
	GetAllRekamMedisRepo(pagination *models.Pagination) ([]Rekam_medis, *models.Pagination, error)
	CreateRekamMedisRepo(*Rekam_medis) error
	GetDetailRekamMedisRepo(id int) (*Rekam_medis, error)
	UpdateRekamMedisRepo(*Rekam_medis) error
	DeleteRekamMedisRepo(id int) error
}

type RekamUseCase interface {
	GetAllRekamMedisUC(*gin.Context) ([]Rekam_medis, *models.Pagination, error)
	CreateRekamMedisUC(*gin.Context) error
	GetDetailRekamMedisUC(*gin.Context) (*Rekam_medis, error)
	UpdateRekamMedisUC(*gin.Context) error
	DeleteRekamMedisUC(*gin.Context) error
}
