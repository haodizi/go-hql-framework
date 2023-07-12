package main

import (
	"awesomeProject/utils"
	"fmt"
)

func main() {
	mysqlutils := new(utils.MysqlUtils)
	record := mysqlutils.GetRecordById("sltt_admin", 1)
	fmt.Println("record=", record)
}
