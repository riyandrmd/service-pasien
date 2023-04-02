package poli-repository

//try

import (
	"fiture-poli/poli"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type PoliRepository struct{
	db *gorm.DB
}

func NewPoliRepository(db *gorm.DB) *PoliRepository{
	return &PoliRepository{db}
}

func (poliRepo *PoliRepository) GetAll() ([]*poli.Poli){
	var result []*poli.Poli
	data := poliRepo.db.Find(&result)
	if data.Error != nil {
		return nil, nil, data.Error
	}
	return result
}

func (poliRepo *PoliRepository) CreatePoli(data *poli.Poli) error {
	data.Slug = slug.Make(data.Title)
	err := poliRepo.db.Create(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (poliRepo *PoliRepository) UpdatePoli(data *poli.Poli) error {
	err := poliRepo.db.First(&poli.Poli{}, "id = ?", data.ID).Error
	if err != nil {
		return err
	}

	data.Slug = slug.Make(data.Title)
	err = poliRepo.db.Model(&poli.Poli{}).Where("id = ?", data.ID).Updates(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (poliRepo *PoliRepository) DeletePoli(id int) error {
	err := poliRepo.db.First(&poli.Poli{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	err = poliRepo.db.Delete(&poli.Poli{}, poliRepo.db.Where("id = ?", id)).Error	
	if err != nil {
		return err
	}

	return nil
}