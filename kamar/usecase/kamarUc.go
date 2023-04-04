package usecase

import (
	"administrasi/kamar"
	"administrasi/models"
	"administrasi/request"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func NewKamarUseCase(kamarRepo kamar.KamarRepo, redis *redis.Client) *KamarUsecase {
	return &KamarUsecase{
		KamarRepo: kamarRepo,
		redis:     redis,
	}
}

type KamarUsecase struct {
	KamarRepo kamar.KamarRepo
	redis     *redis.Client
}

func (KamarUsecase *KamarUsecase) GetAllKamarUC(c *gin.Context) ([]kamar.Kamar, *models.Pagination, error) {
	var result []kamar.Kamar

	pagination, err := request.Paginate(c)
	if err != nil {
		return nil, nil, err
	}

	dataRedis, err := KamarUsecase.redis.Get(c, "kamar").Result()
	if err != nil {
		fmt.Println("database")
		result, pagination, err := KamarUsecase.KamarRepo.GetAllKamarRepo(pagination)
		if err != nil {
			return nil, nil, err
		}

		datajson, err := json.Marshal(result)
		if err != nil {
			return nil, nil, err
		}

		err = KamarUsecase.redis.Set(c, "kamar", (datajson), 0).Err()
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

func (KamarUsecase *KamarUsecase) CreateKamarUC(c *gin.Context) error {
	var result kamar.Kamar
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = KamarUsecase.KamarRepo.CreateKamarRepo(&result)
	if err != nil {
		return err
	}

	KamarUsecase.redis.Del(c, "kamar")

	return nil
}

func (KamarUsecase *KamarUsecase) GetDetailKamarUC(c *gin.Context) (*kamar.Kamar, error) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return nil, err
	}

	result, err := KamarUsecase.KamarRepo.GetDetailKamarRepo(ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (KamarUsecase *KamarUsecase) UpdateKamarUC(c *gin.Context) error {
	var result kamar.Kamar
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	result.Id_kamar = uint(ID)

	err = KamarUsecase.KamarRepo.UpdateKamarRepo(&result)
	if err != nil {
		return err
	}

	KamarUsecase.redis.Del(c, "kamar")

	return nil
}

func (KamarUsecase *KamarUsecase) DeleteKamarUC(c *gin.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = KamarUsecase.KamarRepo.DeleteKamarRepo(ID)
	if err != nil {
		return err
	}

	KamarUsecase.redis.Del(c, "kamar")

	return nil
}
