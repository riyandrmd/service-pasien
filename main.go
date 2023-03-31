package main

import (
	"administrasi/connection"
	"administrasi/router"
	"context"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := connection.ConnectToDb()
	ctx := context.Background()
	redis := connection.ConnectToRedis()
	rh := &router.Handlers{
		Ctx:   ctx,
		DB:    db,
		R:     r,
		Redis: redis,
	}
	rh.Routes()

	r.Run(":8082")
}
