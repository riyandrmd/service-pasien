package kamar



import (
	"administrasi/models"

	"github.com/gin-gonic/gin"
)

type Kamar struct {
	Id_Kamar int `gorm:"primaryKey;autoIncrement" json:"id_kamar"`
	Nama_Kamar string `json:"nama_kamar" binding:"required"`

} 


type KamarRepo interface {
	GetAllKamarRepo(pagination *models.Pagination) ([]Kamar, *models.Pagination, error)
	CreateKamarRepo(*Kamar) error
	GetDetailKamarRepo(id int) (*Kamar, error)
	UpdateKamarRepo(*Kamar) error
	DeleteKamarRepo(id int) error

}

type KamarUseCase interface {
	GetAllKamarUC(*gin.Context) ([]Kamar, *models.Pagination, error)
	CreateKamarUC(*gin.Context) error
	GetDetailKamarUc(*gin.Context) (*Kamar, error)
	UpdateKamarUC(*gin.Context) error
	DeleteKamarUC(*gin.Context) error
}
