// Package ge slack

package transaction

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/zhitongLIU/lentslog/database"
)

type TransactionContent struct {
	gorm.Model
	Sum  float64
	From string
	To   string
}

func InitTransactionContentTable() {
	db, err := database.DBConn()
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	db.AutoMigrate(&TransactionContent{})
}

func NewTransactionContent(sum float64, from string, to string) *TransactionContent {
	db, err := database.DBConn()
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	transactionContent := TransactionContent{Sum: sum, From: from, To: to}
	db.Create(&transactionContent)

	return &transactionContent
}
