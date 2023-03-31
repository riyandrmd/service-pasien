package connection

import (
	"administrasi/rekam_medis"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost",
		"postgres",
		"hasanudin203",
		"rumah_sakit",
		"5432",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&rekam_medis.Rekam_medis{})

	return db
}

func ConnectToRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
	})
	return rdb
}
