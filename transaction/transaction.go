package transaction

import (
	"crypto/sha256"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/zhitongLIU/lentslog/database"
)

type Transaction struct {
	gorm.Model
	CheckSum string
	Head     string
	Tail     string
	Content  Content
}

type Content struct {
	gorm.Model
	Sum  float64
	From string
	To   string
}

func InitTables() {
	db, err := database.DBConn()
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	db.AutoMigrate(&Content{})
	db.AutoMigrate(&Transaction{})
}

func All() []Transaction {
	db, err := database.DBConn()
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	var transactions []Transaction
	db.First(&transactions)
	return transactions
}

// Create a new transaction
func New(sum float64, from string, to string) (*Transaction, error) {
	db, _ := database.DBConn()
	defer db.Close()

	// create transaction content
	content := Content{Sum: sum, From: from, To: to}
	db.Create(&content)

	// create transaction
	checkSum := encrypt(content)
	transaction := Transaction{
		Content:  content,
		CheckSum: checkSum,
	}

	var lastTransaction Transaction
	if db.Last(&lastTransaction).Error != gorm.ErrRecordNotFound {
		lastTransaction.Tail = transaction.CheckSum
		transaction.Head = lastTransaction.CheckSum
		db.Save(&lastTransaction)
	} else {
		transaction.Head = "genesis"
	}

	db.Create(&transaction)
	return &transaction, nil
}

func encrypt(content Content) string {
	s := fmt.Sprintf("%s", content)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}
