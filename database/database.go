package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

func DBConn() (*gorm.DB, error) {
	// db_config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PWD"))
	// db_config := "host=postgres port=5432 dbname=lentslog user=user password=password sslmode=disable"
	db_config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PWD"))
	fmt.Println(db_config)

	db, err := gorm.Open("postgres", db_config)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	return db, nil
}
