package connection

import (
	"administrasi/dokter"
	"administrasi/pasien"
	"administrasi/poli"
	"administrasi/rekammedis"
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
		"101102",
		"rumah_sakit",
		"5432",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&poli.Poli{}, &pasien.Pasien{}, &dokter.Dokter{}, &rekammedis.RekamMedis{})

	return db
}

func ConnectToRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
	})
	return rdb
}
