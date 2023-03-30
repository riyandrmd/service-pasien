package usecase

import (
	"administrasi/models"
	"administrasi/pasien"
	"administrasi/request"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func NewPasienUseCase(pasRepo pasien.PasienRepo, redis *redis.Client) *PasienUsecase {
	return &PasienUsecase{
		pasienRepo: pasRepo,
		redis:      redis,
	}
}

type PasienUsecase struct {
	pasienRepo pasien.PasienRepo
	redis      *redis.Client
}

func (pasienUC *PasienUsecase) GetAllPasienUC(c *gin.Context) ([]pasien.Pasien, *models.Pagination, error) {
	var result []pasien.Pasien

	pagination, err := request.Paginate(c)
	if err != nil {
		return nil, nil, err
	}

	dataRedis, err := pasienUC.redis.Get(c, "pasien").Result()
	if err != nil {
		fmt.Println("database")
		result, pagination, err := pasienUC.pasienRepo.GetAllPasienRepo(pagination)
		if err != nil {
			return nil, nil, err
		}

		datajson, err := json.Marshal(result)
		if err != nil {
			return nil, nil, err
		}

		err = pasienUC.redis.Set(c, "pasien", (datajson), 0).Err()
		if err != nil {
			return nil, nil, err
		}

		return result, pagination, nil
	} else {
		fmt.Println("redis")
		err := json.Unmarshal([]byte(dataRedis), &result)
		if err != nil {
			return nil, nil, err
		}

		return result, pagination, nil
	}
}

func (pasienUC *PasienUsecase) CreatePasienUC(c *gin.Context) error {
	var result pasien.Pasien
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = pasienUC.pasienRepo.CreatePasienRepo(&result)
	if err != nil {
		return err
	}

	pasienUC.redis.Del(c, "pasien")

	return nil
}

func (pasienUC *PasienUsecase) GetDetailPasienUC(c *gin.Context) (*pasien.Pasien, error) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return nil, err
	}

	result, err := pasienUC.pasienRepo.GetDetailPasienRepo(ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (pasienUC *PasienUsecase) UpdatePasienUC(c *gin.Context) error {
	var result pasien.Pasien
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	result.Id_Pasien = uint(ID)

	err = pasienUC.pasienRepo.UpdatePasienRepo(&result)
	if err != nil {
		return err
	}

	return nil
}

func (pasienUC *PasienUsecase) DeletePasienUC(c *gin.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = pasienUC.pasienRepo.DeletePasienRepo(ID)
	if err != nil {
		return err
	}

	pasienUC.redis.Del(c, "pasien")

	return nil
}
