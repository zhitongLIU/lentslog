package initialize

import (
	"github.com/zhitongLIU/lentslog/transaction"
)

func Execute() {
	transaction.InitTables()
}
