package usecase

import (
	"administrasi/models"
	"administrasi/kamar"
	"administrasi/request"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func NewKamarUseCase(kamRepo kamar.PasienRepo, redis *redis.Client) *KamarUseCase {
	return &KamarUseCase{
		kamarRepo: kamRepo,
		redis:     redis,
	}
}

type KamarUseCase struct {
	kamarRepo kamar.KamarRepo
	redis     *redis.Client
}

func (kamarUC *KamarUsecase) GetAllPasienUC(c *gin.Context) ([]kamar.Kamar, *models.Pagination, error) {
	var result []kamar.Kamar

	pagination, err := request.Paginate(c)
	if err != nil {
		return nil, nil, err
	}

	dataRedis, err := kamarUC.redis.Get(c, "kamar").Result()
	if err != nil {
		fmt.Println("database")
		result, pagination, err := kamarUC.kamarRepo.GetAllKamarRepo(pagination)
		if err != nil {
			return nil, nil, err
		}

		datajson, err := json.Marshal(result)
		if err != nil {
			return nil, nil, err
		}

		err = kamarUC.redis.Set(c, "kamar", (datajson), 0).Err()
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

func (kamarUC *KamarUseCase) CreateKamarUC(c *gin.Context) error {
	var result kamar.Kamar
	err := c.ShouldBindJSON(&result)

	if err != nil {
		return err
	}

	err = kamarUC.kamarRepo.CreateKamarRepo(&result)

	if err != nil {
		return err
	}
	kamarUC.redis.Del(c, "kamar")

	return nil
}

func (kamarUC *KamarUseCase) GetDetailKamarUC(c *gin.Context) (*kamar.Kamar, error) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return nil, err
	}

	result, err := kamarUC.kamarRepo.GetDetailKamarRepo(ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (kamarUC *KamarUseCase) UpdateKamarUC(c *gin.Context) error {
	var result kamar.Kamar
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	result.Id_Kamar = uint(ID)

	err = kamarUC.kamarRepo.UpdateKamarRepo(&result)
	if err != nil {
		return err
	}

	return nil
}

func (kamarUC *KamarUseCase) DeleteKamarUC(c *gin.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = kamarUC.kamarRepo.DeleteKamarRepo(ID)
	if err != nil {
		return err
	}

	kamarUC.redis.Del(c, "kamar")

	// return nil
}
