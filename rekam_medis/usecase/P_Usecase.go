package usecase

import (
	"administrasi/models"
	"administrasi/rekam_medis"
	"administrasi/request"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func NewRekamUseCase(rekRepo rekam_medis.RekamRepo, redis *redis.Client) *RekamUsecase {
	return &RekamUsecase{
		rekamRepo: rekRepo,
		redis:     redis,
	}
}

type RekamUsecase struct {
	rekamRepo rekam_medis.RekamRepo
	redis     *redis.Client
}

func (rekamUC *RekamUsecase) GetAllRekamMedisUC(c *gin.Context) ([]rekam_medis.Rekam_medis, *models.Pagination, error) {
	var result []rekam_medis.Rekam_medis

	pagination, err := request.Paginate(c)
	if err != nil {
		return nil, nil, err
	}

	dataRedis, err := rekamUC.redis.Get(c, "rekam_medis").Result()
	if err != nil {
		fmt.Println("database")
		result, pagination, err := rekamUC.rekamRepo.GetAllRekamMedisRepo(pagination)
		if err != nil {
			return nil, nil, err
		}

		datajson, err := json.Marshal(result)
		if err != nil {
			return nil, nil, err
		}

		err = rekamUC.redis.Set(c, "rekam_medis", (datajson), 0).Err()
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

func (rekamUC *RekamUsecase) CreateRekamMedisUC(c *gin.Context) error {
	var result rekam_medis.Rekam_medis
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = rekamUC.rekamRepo.CreateRekamMedisRepo(&result)
	if err != nil {
		return err
	}

	rekamUC.redis.Del(c, "rekam_medis")

	return nil
}

func (rekamUC *RekamUsecase) GetDetailRekamMedisUC(c *gin.Context) (*rekam_medis.Rekam_medis, error) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return nil, err
	}

	result, err := rekamUC.rekamRepo.GetDetailRekamMedisRepo(ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (rekamUC *RekamUsecase) UpdateRekamMedisUC(c *gin.Context) error {
	var result rekam_medis.Rekam_medis
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	result.Id_RekamMedis = uint(ID)

	err = rekamUC.rekamRepo.UpdateRekamMedisRepo(&result)
	if err != nil {
		return err
	}

	return nil
}

func (rekamUC *RekamUsecase) DeleteRekamMedisUC(c *gin.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = rekamUC.rekamRepo.DeleteRekamMedisRepo(ID)
	if err != nil {
		return err
	}

	rekamUC.redis.Del(c, "rekam_medis")

	return nil
}
