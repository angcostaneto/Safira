package UseCase

import (
	"Safira/Domain/Entities"
	"Safira/Domain/Repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUserCase(c *gin.Context) {
	var inputUser Entities.CreateUserInput
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &Entities.User{
		Name:     inputUser.Name,
		Email:    inputUser.Email,
		Password: inputUser.Password,
	}

	repository := Repositories.UserRepository{}
	repository.Create(user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
