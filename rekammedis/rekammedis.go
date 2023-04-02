package rekammedis

import (
	"administrasi/models"

	"github.com/gin-gonic/gin"
)

type RekamMedis struct {
	Id_Rekdis   uint   `gorm:"PrimaryKey" json:"id_rekdis"`
	Diagnosa    string `json:"diagnosa"`
	Id_dokter   int    `json:"id_dokter"`
	Tgl_periksa string `json:"tgl_periksa"`
	Rawat_Inap  int    `json:"rawat_inap"`
	Id_Pasien   uint   `json:"id_pasien"`
	Id_kamar    int    `json:"id_kamar"`
	Id_obat     int    `json:"id_obat"`
}

type DetailRekdis struct {
	Id_Rekdis   uint   `json:"id_rekdis"`
	Nama_Pasien string `json:"nama_pasien"`
	Umur        int    `json:"umur"`
	Gender      string `json:"gender"`
	Alamat      string `json:"alamat"`
	Nama_Poli   string `json:"nama_poli"`
	Tgl_periksa string `json:"tgl_periksa"`
	Diagnosa    string `json:"diagnosa"`
}

type RekamMedisRepo interface {
	GetAllRekamMedisRepo(pagination *models.Pagination) ([]RekamMedis, *models.Pagination, error)
	CreateRekamMedisRepo(*RekamMedis) error
	GetDetailRekamMedisRepo(id int) (*RekamMedis, error)
	UpdateRekamMedisRepo(*RekamMedis) error
	DeleteRekamMedisRepo(id int) error
}

type RekamMedisUseCase interface {
	GetAllRekamMedisUC(*gin.Context) ([]RekamMedis, *models.Pagination, error)
	CreateRekamMedisUC(*gin.Context) error
	GetDetailRekamMedisUC(*gin.Context) (*RekamMedis, error)
	UpdateRekamMedisUC(*gin.Context) error
	DeleteRekamMedisUC(*gin.Context) error
}
