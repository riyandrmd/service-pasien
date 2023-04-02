package fiture-doctor
import (
	"fiture-poli/poli"
	"github.com/gin-gonic/gin"
)

type Dokter struct{
	Id_dokter int `gorm:primaryKey json:"id_dokter"`
	Nama_dokter string `json:"nama_poli"`
	No_hp int `json:"no_hp"`
	Id_poli int `gorm:foreignKey json:"id_poli"`
}

type DoctorRepo interface{
	GetDoctor  ([]*Dokter)
	CreateDoctor (*Dokter) error
}

type DoctorUseCase interface{
	GetDoctor  (*gin.Context)
	CreateDoctor (*gin.Context) error
}