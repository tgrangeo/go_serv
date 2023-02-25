package database

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/tgrangeo/go_serv/models"
	
)

type Dbinstance struct {
	DB *gorm.DB
}
var DB Dbinstance

func ConnectDb(){
	dsn := fmt.Sprintf(
        "host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=France/Paris",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

	db, err := gorm.Open(postgres.Open(dsn),&gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })

	if err != nil {
		log.Fatal("failed to connect to databse. \n", err)
		os.Exit(1)
	}

	log.Println("connected")
    db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
    db.AutoMigrate(&models.Fact{})
}