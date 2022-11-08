package main

import (
	"Safira/Infra/Database"
	"Safira/Infra/Migration"
	"Safira/Presentation"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	Database.ConnectDatabase()
	Migration.ExecuteMigration()

	r.POST("/user", Presentation.CreateUserController)
	r.Run()
}
