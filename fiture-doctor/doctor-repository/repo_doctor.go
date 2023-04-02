package doctor-repository

//try

import (
	"fiture-doctor/doctor"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type DoctorRepository struct{
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) *DoctorRepository{
	return &DoctorRepository{db}
}

func (dokterRepo *DoctorRepository) GetDoctor() ([]*doctor.Dokter){
	var result []*doctor.Dokter
	data := dokterRepo.db.Find(&result)
	if data.Error != nil {
		return nil, nil, data.Error
	}
	return result
}

func (dokterRepo *DoctorRepository) CreateDoctor(data *docter.Dokter) error {
	data.Slug = slug.Make(data.Title)
	err := dokterRepo.db.Create(data).Error
	if err != nil {
		return err
	}

	return nil
}
