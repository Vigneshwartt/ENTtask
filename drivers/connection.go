package drivers

import (
	"fmt"
	"os"
	"routersmux/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectingDatabase() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error occured", err)
		return nil
	}
	host := os.Getenv("host")
	user := os.Getenv("user")
	port := os.Getenv("port")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")
	path := fmt.Sprintf("host=%s user=%s port=%s password=%s dbname=%s", host, user, port, password, dbname)

	connection, err := gorm.Open(postgres.Open(path), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = connection.AutoMigrate(&models.Employee{}, &models.Address{})
	if err != nil {
		panic(err)
	}

	defer HandlePanic()
	fmt.Println("Connected sucesfully")
	return connection
}

func HandlePanic() {
	a := recover()
	if a != nil {
		fmt.Println("RECOVER", a)
	}
}
