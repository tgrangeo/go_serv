package database

import (
	"fmt"
	"log"
	"os"

	"github.com/tgrangeo/go_serv/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDb() {
	var err error

	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
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
}

func DropTable(){
	db.Migrator().DropTable(&models.Todolist{})
}

func InsertDB() {
	create("t","tt")
	create("e","ee")
	create("f","ff")
	create("g","gg")
	create("3","33")
	create("h","hh")
	create("b","bb")
	update("b", "newb","newbb")
	printAll()
}

func create(name string, desc string) {
	todo := models.Todolist{Name: name, Description: desc}
	db.Create(&todo)
}

func update(old string, newName string, newDescription string){
	db.Model(&models.Todolist{}).Where("name = ?", old).Update("Name", newName)
	db.Model(&models.Todolist{}).Where("name = ?", old).Update("Description", newDescription)
}

func printAll() {
	var todos []models.Todolist
	db.Find(&todos)
	for x := range todos {
		fmt.Println(todos[x].Name)
	}
}

func deleteByName(name string) {
	db.Where("name = ?", name).Delete(&models.Todolist{})
}

func findByName(name string) *models.Todolist {
	var ret models.Todolist
	db.Where("name = ?", name).First(ret)
	return &ret
}
