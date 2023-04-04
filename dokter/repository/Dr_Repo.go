package repository

import (
	"administrasi/dokter"
	"administrasi/models"
	"administrasi/paginator"

	"gorm.io/gorm"
)

type DokterRepo struct {
	db *gorm.DB
}

func NewDokterRepo(db *gorm.DB) *DokterRepo {
	return &DokterRepo{
		db: db,
	}
}

func (DokterRepo *DokterRepo) GetAllDokterRepo(pagination *models.Pagination) ([]dokter.Dokter, *models.Pagination, error) {
	var result []dokter.Dokter

	data := DokterRepo.db.Model(&dokter.Dokter{}).Preload("RekamMedis").Find(&result).Limit(pagination.Limit).Offset(pagination.Offset)

	if data.Error != nil {
		return nil, nil, data.Error
	}

	count := data.RowsAffected
	pagination.Count = int(count)

	pagination = paginator.Paging(pagination)

	return result, pagination, nil
}

func (DokterRepo *DokterRepo) CreateDokterRepo(data *dokter.Dokter) error {
	err := DokterRepo.db.Create(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (DokterRepo *DokterRepo) GetDetailDokterRepo(id int) (*dokter.Dokter, error) {
	err := DokterRepo.db.First(&dokter.Dokter{}, "Id_Dokter = ?", id).Error
	if err != nil {
		return nil, err
	}

	var result *dokter.Dokter
	DokterRepo.db.Where("Id_Dokter = ?", id).Find(&dokter.Dokter{}).Scan(&result)

	return result, nil
}

func (DokterRepo *DokterRepo) UpdateDokterRepo(data *dokter.Dokter) error {
	err := DokterRepo.db.First(&dokter.Dokter{}).Where("Id_Dokter = ?", data.Id_Dokter).Error
	if err != nil {
		return err
	}

	err = DokterRepo.db.Model(&dokter.Dokter{}).Where("Id_Dokter = ?", data.Id_Dokter).Updates(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (DokterRepo *DokterRepo) DeleteDokterRepo(id int) error {
	err := DokterRepo.db.First(&dokter.Dokter{}, "Id_Dokter = ?", id).Error
	if err != nil {
		return err
	}

	err = DokterRepo.db.Delete(&dokter.Dokter{}, DokterRepo.db.Where("Id_Dokter = ?", id)).Error
	if err != nil {
		return err
	}

	return nil
}
