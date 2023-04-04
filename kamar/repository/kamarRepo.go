package repository

import (
	"administrasi/kamar"
	"administrasi/models"
	"administrasi/paginator"

	"gorm.io/gorm"
)

type KamarRepo struct {
	db *gorm.DB
}

func NewKamarRepo(db *gorm.DB) *KamarRepo {
	return &KamarRepo{
		db: db,
	}
}

func (KamarRepo *KamarRepo) GetAllKamarRepo(pagination *models.Pagination) ([]kamar.Kamar, *models.Pagination, error) {
	var result []kamar.Kamar

	data := KamarRepo.db.Find(&result).Limit(pagination.Limit).Offset(pagination.Offset)
	if data.Error != nil {
		return nil, nil, data.Error
	}

	count := data.RowsAffected
	pagination.Count = int(count)

	pagination = paginator.Paging(pagination)

	return result, pagination, nil
}

func (KamarRepo *KamarRepo) CreateKamarRepo(data *kamar.Kamar) error {
	err := KamarRepo.db.Create(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (KamarRepo *KamarRepo) GetDetailKamarRepo(id int) (*kamar.Kamar, error) {
	err := KamarRepo.db.First(&kamar.Kamar{}, "Id_kamar = ?", id).Error
	if err != nil {
		return nil, err
	}

	var result *kamar.Kamar
	KamarRepo.db.Where("Id_kamar = ?", id).Find(&kamar.Kamar{}).Scan(&result)

	return result, nil
}

func (KamarRepo *KamarRepo) UpdateKamarRepo(data *kamar.Kamar) error {
	err := KamarRepo.db.First(&kamar.Kamar{}).Where("Id_kamar = ?", data.Id_kamar).Error
	if err != nil {
		return err
	}

	err = KamarRepo.db.Model(&kamar.Kamar{}).Where("Id_kamar = ?", data.Id_kamar).Updates(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (KamarRepo *KamarRepo) DeleteKamarRepo(id int) error {
	err := KamarRepo.db.First(&kamar.Kamar{}, "Id_kamar = ?", id).Error
	if err != nil {
		return err
	}

	err = KamarRepo.db.Delete(&kamar.Kamar{}, KamarRepo.db.Where("Id_kamar = ?", id)).Error
	if err != nil {
		return err
	}

	return nil
}
