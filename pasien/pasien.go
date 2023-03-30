package pasien

import (
	"administrasi/models"

	"github.com/gin-gonic/gin"
)

type Pasien struct {
	Id_Pasien   uint   `gorm:"primaryKey;autoIncreament;" json:"id_pasien"`
	Nama_Pasien string `json:"nama_pasien"`
	Umur        int    `json:"umur"`
	Gender      string `json:"gender"`
	Alamat      string `json:"alamat"`
	No_Telp     string `json:"no_telp"`
	Id_Poli     int64  `json:"id_poli"`
}

type PasienRepo interface {
	GetAllPasienRepo(pagination *models.Pagination) ([]Pasien, *models.Pagination, error)
	CreatePasienRepo(*Pasien) error
	GetDetailPasienRepo(id int) (*Pasien, error)
	UpdatePasienRepo(*Pasien) error
	DeletePasienRepo(id int) error
}

type PasienUseCase interface {
	GetAllPasienUC(*gin.Context) ([]Pasien, *models.Pagination, error)
	CreatePasienUC(*gin.Context) error
	GetDetailPasienUC(*gin.Context) (*Pasien, error)
	UpdatePasienUC(*gin.Context) error
	DeletePasienUC(*gin.Context) error
}
