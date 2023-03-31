package obat

import (
	"administrasi/models"
	"github.com/gin-gonic/gin"


)

type Obat struct {
	Id_Obat uint `gorm:"primaryKey;autoIncrement;" json:"id_obat"`
	Nama_Obat string `json:"nama_obat"`
	Stok int `json:"stok"`
}

type ObatRepo interface {
	GetAllObatRepo(pagination *models.Pagination) ([]Obat, *models.Pagination, error)
	CreateObatRepo(*Obat) error
	GetDetailObatRepo(id int) (*Obat, error)
	UpdateObatRepo(*Obat) error
	DeleteObatRepo(id int) error
}


type ObatUseCase interface {
	GetAllObatUC(*gin.Context) ([]Obat, *models.Pagination, error)
	CreateObatUC(*gin.Context) error
	GetDetailObatUC(*gin.Context) (*Obat, error)
	UpdateObatUC(*gin.Context) error
	DeleteObatUC(*gin.Context) error
}