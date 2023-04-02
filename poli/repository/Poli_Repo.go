package repository

import (
	"administrasi/models"
	"administrasi/paginator"
	"administrasi/poli"

	"gorm.io/gorm"
)

type PoliRepo struct {
	db *gorm.DB
}

func NewPoliRepo(db *gorm.DB) *PoliRepo {
	return &PoliRepo{
		db: db,
	}
}

func (PoliRepo *PoliRepo) GetAllPoliRepo(pagination *models.Pagination) ([]poli.Poli, *models.Pagination, error) {
	var result []poli.Poli

	data := PoliRepo.db.Model(&poli.Poli{}).Preload("Pasien.RekamMedis").Find(&result).Limit(pagination.Limit).Offset(pagination.Offset)
	if data.Error != nil {
		return nil, nil, data.Error
	}

	count := data.RowsAffected
	pagination.Count = int(count)

	pagination = paginator.Paging(pagination)

	return result, pagination, nil
}

func (PoliRepo *PoliRepo) CreatePoliRepo(data *poli.Poli) error {
	err := PoliRepo.db.Create(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (PoliRepo *PoliRepo) GetDetailPoliRepo(id int) (*poli.Poli, error) {
	err := PoliRepo.db.First(&poli.Poli{}, "id_poli = ?", id).Error
	if err != nil {
		return nil, err
	}

	var result *poli.Poli
	PoliRepo.db.Model(&poli.Poli{}).Preload("Pasien.RekamMedis").Where("id_poli = ?", id).Find(&result)

	return result, nil
}

func (PoliRepo *PoliRepo) UpdatePoliRepo(data *poli.Poli) error {
	err := PoliRepo.db.First(&poli.Poli{}).Where("id_poli = ?", data.Id_Poli).Error
	if err != nil {
		return err
	}

	err = PoliRepo.db.Model(&poli.Poli{}).Where("id_poli = ?", data.Id_Poli).Updates(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (PoliRepo *PoliRepo) DeletePoliRepo(id int) error {
	err := PoliRepo.db.First(&poli.Poli{}, "id_poli = ?", id).Error
	if err != nil {
		return err
	}

	err = PoliRepo.db.Delete(&poli.Poli{}, PoliRepo.db.Where("id_poli = ?", id)).Error
	if err != nil {
		return err
	}

	return nil
}
