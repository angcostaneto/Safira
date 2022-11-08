package Presentation

import (
	"Safira/UseCase"
	"github.com/gin-gonic/gin"
)

func CreateUserController(c *gin.Context) {
	UseCase.CreateUserCase(c)
}
