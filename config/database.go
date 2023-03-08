package config

import (
	"fmt"
	"os"

	"github.com/Caknoooo/Golang-BLOG/entities"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDatabaseConnection() *gorm.DB{
	if os.Getenv("APP_ENV") != "Production"{
		err := godotenv.Load(".env")
		if err != nil{
			fmt.Println(err)
			panic(err)
		}
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		// Menambahkan opsi berikut akan memungkinkan driver database
		// untuk mendukung tipe data UUID secara bawaan.
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil{
		fmt.Println(err)
		panic(err)
	}

	if err := db.AutoMigrate(
		entities.Blog{}, 
		entities.User{},
	); err != nil{
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Database Connected")
	return db
}

func ClosDatabaseConnection(db *gorm.DB){
	dbSQL, err := db.DB()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	dbSQL.Close()
}