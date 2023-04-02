package dokter

import (
	"administrasi/models"

	"github.com/gin-gonic/gin"
)

type Dokter struct {
	Id_Dokter   uint   `gorm:"PrimaryKey" json:"id_dokter"`
	Nama_Dokter string `json:"nama_dokter"`
	No_Telp     string `json:"no_telp"`
	Id_Poli     uint   `json:"id_poli"`
}

type DokterRepo interface {
	GetAllDokterRepo(pagination *models.Pagination) ([]Dokter, *models.Pagination, error)
	CreateDokterRepo(*Dokter) error
	GetDetailDokterRepo(id int) (*Dokter, error)
	UpdateDokterRepo(*Dokter) error
	DeleteDokterRepo(id int) error
}

type DokterUseCase interface {
	GetAllDokterUC(*gin.Context) ([]Dokter, *models.Pagination, error)
	CreateDokterUC(*gin.Context) error
	GetDetailDokterUC(*gin.Context) (*Dokter, error)
	UpdateDokterUC(*gin.Context) error
	DeleteDokterUC(*gin.Context) error
}
