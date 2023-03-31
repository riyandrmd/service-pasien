package poli

import (
	"administrasi/models"
	"administrasi/pasien"

	"github.com/gin-gonic/gin"
)

type Poli struct {
	Id_Poli   uint            `gorm:"PrimaryKey" json:"id_poli"`
	Nama_Poli string          `json:"nama_poli"`
	Pasien    []pasien.Pasien `gorm:"foreignKey:Id_Poli" constraint:"OnUpdate:CASCADE,OnDelete:SET NULL;" json:"pasien"`
}

type PoliRepo interface {
	GetAllPoliRepo(pagination *models.Pagination) ([]Poli, *models.Pagination, error)
	CreatePoliRepo(*Poli) error
	GetDetailPoliRepo(id int) (*Poli, error)
	UpdatePoliRepo(*Poli) error
	DeletePoliRepo(id int) error
}

type PoliUseCase interface {
	GetAllPoliUC(*gin.Context) ([]Poli, *models.Pagination, error)
	CreatePoliUC(*gin.Context) error
	GetDetailPoliUC(*gin.Context) (*Poli, error)
	UpdatePoliUC(*gin.Context) error
	DeletePoliUC(*gin.Context) error
}
