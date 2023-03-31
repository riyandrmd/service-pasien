package repository

import (
	"administrasi/models"
	"administrasi/paginator"
	"administrasi/rekam_medis"

	"gorm.io/gorm"
)

type RekamRepo struct {
	db *gorm.DB
}

func NewRekamRepo(db *gorm.DB) *RekamRepo {
	return &RekamRepo{
		db: db,
	}
}

func (RmRepo *RekamRepo) GetAllRekamMedisRepo(pagination *models.Pagination) ([]rekam_medis.Rekam_medis, *models.Pagination, error) {
	var result []rekam_medis.Rekam_medis
	data := RmRepo.db.Find(&result).Limit(pagination.Limit).Offset(pagination.Offset)
	if data.Error != nil {
		return nil, nil, data.Error
	}

	count := data.RowsAffected
	pagination.Count = int(count)

	pagination = paginator.Paging(pagination)

	return result, pagination, nil
}

func (RmRepo *RekamRepo) CreateRekamMedisRepo(data *rekam_medis.Rekam_medis) error {
	err := RmRepo.db.Create(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (RmRepo *RekamRepo) GetDetailRekamMedisRepo(id int) (*rekam_medis.Rekam_medis, error) {
	err := RmRepo.db.First(&rekam_medis.Rekam_medis{}, "id_rekam_medis = ?", id).Error
	if err != nil {
		return nil, err
	}

	var result *rekam_medis.Rekam_medis
	RmRepo.db.Where("id_rekam_medis = ?", id).Find(&rekam_medis.Rekam_medis{}).Scan(&result)

	return result, nil
}

func (RmRepo *RekamRepo) UpdateRekamMedisRepo(data *rekam_medis.Rekam_medis) error {
	err := RmRepo.db.First(&rekam_medis.Rekam_medis{}).Where("id_rekam_medis = ?", data.Id_RekamMedis).Error
	if err != nil {
		return err
	}

	err = RmRepo.db.Model(&rekam_medis.Rekam_medis{}).Where("id_rekam_medis = ?", data.Id_RekamMedis).Updates(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (RmRepo *RekamRepo) DeleteRekamMedisRepo(id int) error {
	err := RmRepo.db.First(&rekam_medis.Rekam_medis{}, "id_rekam_medis = ?", id).Error
	if err != nil {
		return err
	}

	err = RmRepo.db.Delete(&rekam_medis.Rekam_medis{}, RmRepo.db.Where("id_rekam_medis = ?", id)).Error
	if err != nil {
		return err
	}

	return nil
}
