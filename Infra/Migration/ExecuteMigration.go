package Migration

import (
	"Safira/Domain/Entities"
	"Safira/Infra/Database"
)

func ExecuteMigration() {
	Database.DB.AutoMigrate(&Entities.User{})
}
