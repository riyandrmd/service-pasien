package repository

import (
	"administrasi/models"
	"administrasi/paginator"
	"administrasi/rekammedis"

	"gorm.io/gorm"
)

type RekDisRepo struct {
	db *gorm.DB
}

func NewRekamMedisRepo(db *gorm.DB) *RekDisRepo {
	return &RekDisRepo{
		db: db,
	}
}

func (RekDisRepo *RekDisRepo) GetAllRekamMedisRepo(pagination *models.Pagination) ([]rekammedis.RekamMedis, *models.Pagination, error) {
	var result []rekammedis.RekamMedis

	data := RekDisRepo.db.Model(&rekammedis.RekamMedis{}).Preload("Kamar").Preload("Obat").Find(&result).Limit(pagination.Limit).Offset(pagination.Offset)
	if data.Error != nil {
		return nil, nil, data.Error
	}

	count := data.RowsAffected
	pagination.Count = int(count)

	pagination = paginator.Paging(pagination)

	return result, pagination, nil
}

func (RekDisRepo *RekDisRepo) CreateRekamMedisRepo(data *rekammedis.RekamMedis) error {
	err := RekDisRepo.db.Create(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (RekDisRepo *RekDisRepo) GetDetailRekamMedisRepo(id int) (*rekammedis.RekamMedis, error) {
	err := RekDisRepo.db.First(&rekammedis.RekamMedis{}, "Id_Rekdis = ?", id).Error
	if err != nil {
		return nil, err
	}

	var result *rekammedis.RekamMedis
	RekDisRepo.db.Model(&rekammedis.RekamMedis{}).Preload("Kamar").Where("Id_Rekdis = ?", id).Find(&result)

	return result, nil
}

func (RekDisRepo *RekDisRepo) UpdateRekamMedisRepo(data *rekammedis.RekamMedis) error {
	err := RekDisRepo.db.First(&rekammedis.RekamMedis{}).Where("Id_Rekdis = ?", data.Id_Rekdis).Error
	if err != nil {
		return err
	}

	err = RekDisRepo.db.Model(&rekammedis.RekamMedis{}).Where("Id_Rekdis = ?", data.Id_Rekdis).Updates(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (RekDisRepo *RekDisRepo) DeleteRekamMedisRepo(id int) error {
	err := RekDisRepo.db.First(&rekammedis.RekamMedis{}, "Id_Rekdis = ?", id).Error
	if err != nil {
		return err
	}

	err = RekDisRepo.db.Delete(&rekammedis.RekamMedis{}, RekDisRepo.db.Where("Id_Rekdis = ?", id)).Error
	if err != nil {
		return err
	}

	return nil
}
