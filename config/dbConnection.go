package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"os"
	models "learn-fiber/model"
)

var DB *gorm.DB

func Connect() {
	// ** DATABASE SETTINGS & CONNECT**
	//load environment variables
	godotenv.Load()
	psql_host := os.Getenv("POSTGRE_HOST")
	psql_user := os.Getenv("POSTGRE_USER")
	psql_password := os.Getenv("POSTGRE_PASSWORD")
	psql_dbname := os.Getenv("POSTGRE_DBNAME")
	psql_port := os.Getenv("POSTGRE_PORT")

	// connect to postgres
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei", psql_host, psql_user, psql_password, psql_dbname, psql_port)
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		panic("DB connection failed")
	}
	DB = db
	fmt.Println("DB connection successfully")
	
	// Auto Create Table (Dont need this for now)
	AutoMigrate(db)
}

func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		&models.Cashier{},
		&models.Category{},
		&models.Payment{},
		&models.PaymentType{},
		&models.Product{},
		&models.Discount{},
		&models.Order{},
	) 
}
