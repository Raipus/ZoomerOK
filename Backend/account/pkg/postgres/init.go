package postgres

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Init() {
	// Подключение к базе данных
	Instance, dbError = gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
			config.Config.PostgresUser,
			config.Config.PostgresPassword,
			config.Config.PostgresHost,
			strconv.Itoa(config.Config.PostgresPort),
			config.Config.PostgresDb),
		PreferSimpleProtocol: true, // отключает использование неявных подготовленных операторов
	}), &gorm.Config{})

	if dbError != nil {
		panic("failed to connect database: " + dbError.Error())
	}

	log.Println("Connected to Database!")
}

func Migrate() {
	if err := Instance.AutoMigrate(&User{}); err != nil {
		log.Fatalf("Error migrating User: %v", err)
	}
	if err := Instance.AutoMigrate(&Friend{}); err != nil {
		log.Fatalf("Error migrating Friend: %v", err)
	}
	log.Println("Database Migration Completed!")
}

func initPostgres() *gorm.DB {
	if gin.Mode() == gin.ReleaseMode {
		Init()
		Migrate()

		if Instance == nil {
			panic("Database not initialized")
		} else {
			log.Println("Database initialized")
			return Instance
		}
	} else {
		return nil
	}
}
