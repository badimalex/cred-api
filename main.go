package main

import (
	"cred-api/internal/api"
	"cred-api/internal/db"
	"cred-api/internal/logging"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	logging.InitLogger()

	db.InitDB()

	err = db.MigrateDB()
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	r := gin.Default()

	api.RegisterRoutes(r)

	r.Run()
}
