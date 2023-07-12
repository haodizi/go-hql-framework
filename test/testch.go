package main

import (
	"awesomeProject/utils"
	"fmt"
	"time"
)

func main() {
	ch := new(utils.ClickHouseUtils)
	sql := "CREATE TABLE IF NOT EXISTS example1 (id UInt64,username String,password FixedString(32),add_time DateTime,status Enum16('锁定'=0,'正常'=1))engine=MergeTree() order by id"
	ch.ExecuteUpdate(sql)
	ch.ExecuteUpdate("insert into example1 (id,username,password,add_time,status) values (?,?,?,?,?)",
		500, "hql223", "a86d3c427285001eb1febc190ffee277", time.Now(), "锁定")
	rows := ch.ExecuteQuery("select * from example1")
	fmt.Println(rows)
}
