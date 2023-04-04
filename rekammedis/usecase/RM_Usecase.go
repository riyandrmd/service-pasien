package usecase

import (
	"administrasi/models"
	"administrasi/rekammedis"
	"administrasi/request"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func NewNewRekamMedisUC(RmRepo rekammedis.RekamMedisRepo, redis *redis.Client) *RekamMedisUC {
	return &RekamMedisUC{
		RekDisRepo: RmRepo,
		redis:      redis,
	}
}

type RekamMedisUC struct {
	RekDisRepo rekammedis.RekamMedisRepo
	redis      *redis.Client
}

func (RekamMedisUC *RekamMedisUC) GetAllRekamMedisUC(c *gin.Context) ([]rekammedis.RekamMedis, *models.Pagination, error) {
	var result []rekammedis.RekamMedis

	pagination, err := request.Paginate(c)
	if err != nil {
		return nil, nil, err
	}

	dataRedis, err := RekamMedisUC.redis.Get(c, "rekammedis").Result()
	if err != nil {
		fmt.Println("database")
		result, pagination, err := RekamMedisUC.RekDisRepo.GetAllRekamMedisRepo(pagination)
		if err != nil {
			return nil, nil, err
		}

		datajson, err := json.Marshal(result)
		if err != nil {
			return nil, nil, err
		}

		err = RekamMedisUC.redis.Set(c, "rekammedis", (datajson), 0).Err()
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

func (RekamMedisUC *RekamMedisUC) CreateRekamMedisUC(c *gin.Context) error {
	var result rekammedis.RekamMedis
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = RekamMedisUC.RekDisRepo.CreateRekamMedisRepo(&result)
	if err != nil {
		return err
	}

	RekamMedisUC.redis.Del(c, "rekammedis")

	return nil
}

func (RekamMedisUC *RekamMedisUC) GetDetailRekamMedisUC(c *gin.Context) (*rekammedis.RekamMedis, error) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return nil, err
	}

	result, err := RekamMedisUC.RekDisRepo.GetDetailRekamMedisRepo(ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (RekamMedisUC *RekamMedisUC) UpdateRekamMedisUC(c *gin.Context) error {
	var result rekammedis.RekamMedis
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	result.Id_Rekdis = uint(ID)

	err = RekamMedisUC.RekDisRepo.UpdateRekamMedisRepo(&result)
	if err != nil {
		return err
	}

	RekamMedisUC.redis.Del(c, "rekammedis")

	return nil
}

func (RekamMedisUC *RekamMedisUC) DeleteRekamMedisUC(c *gin.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = RekamMedisUC.RekDisRepo.DeleteRekamMedisRepo(ID)
	if err != nil {
		return err
	}

	RekamMedisUC.redis.Del(c, "rekammedis")

	return nil
}
