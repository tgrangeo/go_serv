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
	Db *gorm.DB
}
var DB Dbinstance

func ConnectDb(){
	dsn := fmt.Sprintf(
        "host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
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
    db.AutoMigrate(&models.Todolist{})

	DB = Dbinstance{
        Db: db,
    }
}

func CreateDB() {
	DB.Db.Exec(`CREATE TABLE IF NOT EXISTS todolist (
		id SERIAL PRIMARY KEY,
		name TEXT,
		description TEXT
	);`)
}

func InsertDB(){ //name string, description string
	DB.Db.Exec(`INSERT INTO todolist (name, description)
	VALUES('very important','do my homework !!!!!');`)
}

// func (m Todolist) getfirst()  error {
// 	return db.QueryRow("SELECT name, price FROM products WHERE id=$1",
// 		  p.ID).Scan(&p.Name, &p.Price)
// }