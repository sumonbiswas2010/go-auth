package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	// postgres://eozhjvlo:Q-8we2OJGOxaiFvi9n2YUnRU4yHsttfq@tiny.db.elephantsql.com/eozhjvlo
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}
