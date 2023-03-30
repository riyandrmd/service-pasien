package repository

import (
	"administrasi/models"
	"administrasi/paginator"
	"administrasi/pasien"

	"gorm.io/gorm"
)

type PasienRepo struct {
	db *gorm.DB
}

func NewPasienRepo(db *gorm.DB) *PasienRepo {
	return &PasienRepo{
		db: db,
	}
}

func (PasienRepo *PasienRepo) GetAllPasienRepo(pagination *models.Pagination) ([]pasien.Pasien, *models.Pagination, error) {
	var result []pasien.Pasien

	data := PasienRepo.db.Find(&result).Limit(pagination.Limit).Offset(pagination.Offset)
	if data.Error != nil {
		return nil, nil, data.Error
	}

	count := data.RowsAffected
	pagination.Count = int(count)

	pagination = paginator.Paging(pagination)

	return result, pagination, nil
}

func (PasienRepo *PasienRepo) CreatePasienRepo(data *pasien.Pasien) error {
	err := PasienRepo.db.Create(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (PasienRepo *PasienRepo) GetDetailPasienRepo(id int) (*pasien.Pasien, error) {
	err := PasienRepo.db.First(&pasien.Pasien{}, "id_pasien = ?", id).Error
	if err != nil {
		return nil, err
	}

	var result *pasien.Pasien
	PasienRepo.db.Where("id_pasien = ?", id).Find(&pasien.Pasien{}).Scan(&result)

	return result, nil
}

func (PasienRepo *PasienRepo) UpdatePasienRepo(data *pasien.Pasien) error {
	err := PasienRepo.db.First(&pasien.Pasien{}).Where("id_pasien = ?", data.Id_Pasien).Error
	if err != nil {
		return err
	}

	err = PasienRepo.db.Model(&pasien.Pasien{}).Where("id_pasien = ?", data.Id_Pasien).Updates(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (PasienRepo *PasienRepo) DeletePasienRepo(id int) error {
	err := PasienRepo.db.First(&pasien.Pasien{}, "id_pasien = ?", id).Error
	if err != nil {
		return err
	}

	err = PasienRepo.db.Delete(&pasien.Pasien{}, PasienRepo.db.Where("id_pasien = ?", id)).Error
	if err != nil {
		return err
	}

	return nil
}
