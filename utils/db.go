package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	conf2 "github.com/haodizi/go-hql-framework/conf"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	InitConfigFromIni()
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", conf2.DbConf.Username, conf2.DbConf.Password, conf2.DbConf.Host, conf2.DbConf.Port, conf2.DbConf.DbName, conf2.DbConf.Charset)
	//fmt.Println(dataSourceName)
	database, error := sqlx.Open("mysql", dataSourceName)
	if error != nil {
		fmt.Println("Connect to mysql fail,", error)
	}
	// fmt.Println(database)
	error = database.Ping()
	if error != nil {
		fmt.Println("Ping mysql fail,", error)
	}
	Db = database
	//defer Db.Close()
}

func CloseDb() {
	Db.Close()
}
