package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func DBConn() (*gorm.DB, error) {
	// db_config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PWD"))
	// db_config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PWD"))
	db_config := "host=127.0.0.1 port=32768  dbname=lentslog user=postgres sslmode=disable"

	db, err := gorm.Open("postgres", db_config)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	return db, nil
}
