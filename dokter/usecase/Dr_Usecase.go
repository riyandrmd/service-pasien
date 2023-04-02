package usecase

import (
	"administrasi/dokter"
	"administrasi/models"
	"administrasi/request"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func NewDokterUC(DrRepo dokter.DokterRepo, redis *redis.Client) *DokterUC {
	return &DokterUC{
		DrRepo: DrRepo,
		redis:  redis,
	}
}

type DokterUC struct {
	DrRepo dokter.DokterRepo
	redis  *redis.Client
}

func (DokterUC *DokterUC) GetAllDokterUC(c *gin.Context) ([]dokter.Dokter, *models.Pagination, error) {
	var result []dokter.Dokter

	pagination, err := request.Paginate(c)
	if err != nil {
		return nil, nil, err
	}

	dataRedis, err := DokterUC.redis.Get(c, "dokter").Result()
	if err != nil {
		fmt.Println("database")
		result, pagination, err := DokterUC.DrRepo.GetAllDokterRepo(pagination)
		if err != nil {
			return nil, nil, err
		}

		datajson, err := json.Marshal(result)
		if err != nil {
			return nil, nil, err
		}

		err = DokterUC.redis.Set(c, "dokter", (datajson), 0).Err()
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

func (DokterUC *DokterUC) CreateDokterUC(c *gin.Context) error {
	var result dokter.Dokter
	err := c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	err = DokterUC.DrRepo.CreateDokterRepo(&result)
	if err != nil {
		return err
	}
	DokterUC.redis.Del(c, "dokter")

	return nil
}

func (DokterUC *DokterUC) GetDetailDokterUC(c *gin.Context) (*dokter.Dokter, error) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return nil, err
	}

	result, err := DokterUC.DrRepo.GetDetailDokterRepo(ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (DokterUC *DokterUC) UpdateDokterUC(c *gin.Context) error {
	var result dokter.Dokter
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = c.ShouldBindJSON(&result)
	if err != nil {
		return err
	}

	result.Id_Dokter = uint(ID)

	err = DokterUC.DrRepo.UpdateDokterRepo(&result)
	if err != nil {
		return err
	}

	DokterUC.redis.Del(c, "dokter")

	return nil
}

func (DokterUC *DokterUC) DeleteDokterUC(c *gin.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil || ID <= 0 {
		return err
	}

	err = DokterUC.DrRepo.DeleteDokterRepo(ID)
	if err != nil {
		return err
	}

	DokterUC.redis.Del(c, "dokter")

	return nil
}
