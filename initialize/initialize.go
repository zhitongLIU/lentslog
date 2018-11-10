package initialize

import (
	"github.com/zhitongLIU/lentslog/models"
)

func Execute() {
	models.InitTransactionContentTable()
}
