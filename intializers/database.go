package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	// postgres://znlszlda:c_fkzA_5eaXL4QrcTbvqb05P0N-VUrDP@mahmud.db.elephantsql.com/znlszlda

	dsn := os.Getenv("DB_STRING")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to DB", DB)
	} else {
		fmt.Println("Database connected", DB)
	}
}
