package repository

import (
	"administrasi/models"
	"administrasi/obat"
	"administrasi/paginator"

	"gorm.io/gorm"
)

type ObatRepo struct {
	db *gorm.DB
}

func NewObatRepo(db *gorm.DB) *ObatRepo {
	return &ObatRepo{
		db: db,
	}
}

func (ObatRepo *ObatRepo) GetAllObatRepo(pagination *models.Pagination) ([]obat.Obat, *models.Pagination, error) {
	var result []obat.Obat

	data := ObatRepo.db.Find(&result).Limit(pagination.Limit).Offset(pagination.Offset)
	if data.Error != nil {
		return nil, nil, data.Error
	}

	count := data.RowsAffected
	pagination.Count = int(count)

	pagination = paginator.Paging(pagination)

	return result, pagination, nil
}

func (ObatRepo *ObatRepo) CreateObatRepo(data *obat.Obat) error {
	err := ObatRepo.db.Create(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (ObatRepo *ObatRepo) GetDetailObatRepo(id int) (*obat.Obat, error) {
	err := ObatRepo.db.First(&obat.Obat{}, "Id_obat = ?", id).Error
	if err != nil {
		return nil, err
	}

	var result *obat.Obat
	ObatRepo.db.Where("Id_obat = ?", id).Find(&obat.Obat{}).Scan(&result)

	return result, nil
}

func (ObatRepo *ObatRepo) UpdateObatRepo(data *obat.Obat) error {
	err := ObatRepo.db.First(&obat.Obat{}).Where("Id_obat = ?", data.Id_Obat).Error
	if err != nil {
		return err
	}

	err = ObatRepo.db.Model(&obat.Obat{}).Where("Id_obat = ?", data.Id_Obat).Updates(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (ObatRepo *ObatRepo) DeleteObatRepo(id int) error {
	err := ObatRepo.db.First(&obat.Obat{}, "Id_obat = ?", id).Error
	if err != nil {
		return err
	}

	err = ObatRepo.db.Delete(&obat.Obat{}, ObatRepo.db.Where("Id_obat = ?", id)).Error
	if err != nil {
		return err
	}

	return nil
}
