package Repositories

import (
	"Safira/Domain/Entities"
	"Safira/Infra/Database"
)

type UserRepository struct{}

func (r *UserRepository) Create(user *Entities.User) {
	Database.DB.Create(user)
}
