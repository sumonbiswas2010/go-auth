package initializers

import "go-auth/models"

func SyncDB() {
	// DB.Migrator().DropTable(&models.User{})
	DB.AutoMigrate(&models.User{})
	
}
