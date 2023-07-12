package utils

import (
	conf2 "awesomeProject/conf"
	"database/sql"
	"fmt"
	_ "github.com/mailru/go-clickhouse/v2"
)

var chsqldb *sql.DB
var cherr error

func init() {
	InitClickHouseConfigFromIni()
	rdbstr := fmt.Sprintf("http://%s:%s@%s:%s/%s", conf2.ClickHouseConf.Username, conf2.ClickHouseConf.Password, conf2.ClickHouseConf.Host, conf2.ClickHouseConf.Port, conf2.ClickHouseConf.DbName)
	fmt.Println(rdbstr)
	chsqldb, cherr = sql.Open("chhttp", rdbstr)
	if cherr != nil {
		fmt.Println("open database failed:", cherr)
	}
	cherr = chsqldb.Ping()
	if cherr != nil {
		fmt.Println("connect database failed: " + cherr.Error())
	}
}
