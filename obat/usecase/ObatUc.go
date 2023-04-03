package usecase

import (
	"administrasi/models"
	"administrasi/obat"
	"administrasi/request"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func NewObatUseCase(obatRepo obat.ObatRepo, redis *redis.Client) *ObatUseCase {
	return &ObatUseCase{
		obatRepo: obatRepo,
		redis:    redis,
	}
}

type ObatUseCase struct {
	obatRepo obat.ObatRepo
	redis    *redis.Client
}

func (ObatUseCase *ObatUseCase) GetAllObatUC(c *gin.Context) ([]obat.Obat, *models.Pagination, error) {
	var result []obat.Obat

	pagination, err := request.Paginate(c)
	if err != nil {
		return nil, nil, err
	}

	dataRedis, err := ObatUseCase.redis.Get(c, "obat").Result()
	if err != nil {
		fmt.Println("database")
		result, pagination, err := ObatUseCase.obatRepo.GetAllObatRepo(pagination)
		if err != nil {
			return nil, nil, err
		}

		datajson, err := json.Marshal(result)
		if err != nil {
			return nil, nil, err
		}

		err = ObatUseCase.redis.Set(c, "obat", (datajson), 0).Err()
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

func (ObatUseCase *ObatUseCase) CreateObatUC(c *gin.Context) error {
	var result obat.Obat
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = ObatUseCase.obatRepo.CreateObatRepo(&result)
	if err != nil {
		return err
	}

	ObatUseCase.redis.Del(c, "obat")

	return nil
}

func (ObatUseCase *ObatUseCase) GetDetailObatUC(c *gin.Context) (*obat.Obat, error) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return nil, err
	}

	result, err := ObatUseCase.obatRepo.GetDetailObatRepo(ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ObatUseCase *ObatUseCase) UpdateObatUC(c *gin.Context) error {
	var result obat.Obat
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	result.Id_Obat = uint(ID)

	err = ObatUseCase.obatRepo.UpdateObatRepo(&result)
	if err != nil {
		return err
	}

	ObatUseCase.redis.Del(c, "obat")

	return nil
}

func (ObatUseCase *ObatUseCase) DeleteObatUC(c *gin.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = ObatUseCase.obatRepo.DeleteObatRepo(ID)
	if err != nil {
		return err
	}

	ObatUseCase.redis.Del(c, "obat")

	return nil
}
