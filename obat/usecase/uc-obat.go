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

func NewObatUseCase(obtRepo obat.ObatRepo, redis *redis.Client) *ObatUsecase {
	return &ObatUsecase{
		obatRepo: obtRepo,
		redis:      redis,
	}
}

type ObatUsecase struct {
	obatRepo obat.ObatRepo
	redis      *redis.Client
}

func (obatUC *ObatUsecase) GetAllObatUC(c *gin.Context) ([]obat.Obat, *models.Pagination, error) {
	var result []obat.Obat

	pagination, err := request.Paginate(c)
	if err != nil {
		return nil, nil, err
	}

	dataRedis, err := obatUC.redis.Get(c, "obat").Result()
	if err != nil {
		fmt.Println("database")
		result, pagination, err := obatUC.obatRepo.GetAllObatRepo(pagination)
		if err != nil {
			return nil, nil, err
		}

		datajson, err := json.Marshal(result)
		if err != nil {
			return nil, nil, err
		}

		err = obatUC.redis.Set(c, "obat", (datajson), 0).Err()
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

func (obatUC *ObatUsecase) CreateObatUC(c *gin.Context) error {
	var result obat.Obat
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = obatUC.obatRepo.CreateObatRepo(&result)
	if err != nil {
		return err
	}

	obatUC.redis.Del(c, "obat")

	return nil
}

func (obatUC *ObatUsecase) GetDetailObatUC(c *gin.Context) (*obat.Obat, error) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return nil, err
	}

	result, err := obatUC.obatRepo.GetDetailObatRepo(ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (obatUC *ObatUsecase) UpdateObatUC(c *gin.Context) error {
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

	err = obatUC.obatRepo.UpdateObatRepo(&result)
	if err != nil {
		return err
	}

	obatUC.redis.Del(c, "obat")

	return nil
}

func (obatUC *ObatUsecase) DeleteObatUC(c *gin.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = obatUC.obatRepo.DeleteObatRepo(ID)
	if err != nil {
		return err
	}

	obatUC.redis.Del(c, "obat")

	return nil
}
