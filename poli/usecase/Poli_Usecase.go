package usecase

import (
	"administrasi/models"
	"administrasi/poli"
	"administrasi/request"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func NewPoliUseCase(polRepo poli.PoliRepo, redis *redis.Client) *PoliUsecase {
	return &PoliUsecase{
		PoliRepo: polRepo,
		redis:    redis,
	}
}

type PoliUsecase struct {
	PoliRepo poli.PoliRepo
	redis    *redis.Client
}

func (PoliUC *PoliUsecase) GetAllPoliUC(c *gin.Context) ([]poli.Poli, *models.Pagination, error) {
	var result []poli.Poli

	pagination, err := request.Paginate(c)
	if err != nil {
		return nil, nil, err
	}

	dataRedis, err := PoliUC.redis.Get(c, "poli").Result()
	if err != nil {
		fmt.Println("database")
		result, pagination, err := PoliUC.PoliRepo.GetAllPoliRepo(pagination)
		if err != nil {
			return nil, nil, err
		}

		datajson, err := json.Marshal(result)
		if err != nil {
			return nil, nil, err
		}

		err = PoliUC.redis.Set(c, "poli", (datajson), 0).Err()
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

func (PoliUC *PoliUsecase) CreatePoliUC(c *gin.Context) error {
	var result poli.Poli
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = PoliUC.PoliRepo.CreatePoliRepo(&result)
	if err != nil {
		return err
	}

	PoliUC.redis.Del(c, "pasien")
	PoliUC.redis.Del(c, "poli")
	PoliUC.redis.Del(c, "rekammedis")

	return nil
}

func (PoliUC *PoliUsecase) GetDetailPoliUC(c *gin.Context) (*poli.Poli, error) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return nil, err
	}

	result, err := PoliUC.PoliRepo.GetDetailPoliRepo(ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (PoliUC *PoliUsecase) UpdatePoliUC(c *gin.Context) error {
	var result poli.Poli
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	result.Id_Poli = uint(ID)

	err = PoliUC.PoliRepo.UpdatePoliRepo(&result)
	if err != nil {
		return err
	}

	PoliUC.redis.Del(c, "pasien")
	PoliUC.redis.Del(c, "poli")
	PoliUC.redis.Del(c, "rekammedis")

	return nil
}

func (PoliUC *PoliUsecase) DeletePoliUC(c *gin.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = PoliUC.PoliRepo.DeletePoliRepo(ID)
	if err != nil {
		return err
	}

	PoliUC.redis.Del(c, "pasien")
	PoliUC.redis.Del(c, "poli")
	PoliUC.redis.Del(c, "rekammedis")

	return nil
}
