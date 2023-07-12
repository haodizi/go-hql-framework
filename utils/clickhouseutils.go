package utils

import (
	"database/sql"
	"fmt"
)

type ClickHouseUtilsInterface interface {
	ExecuteUpdate(sql string, args ...interface{})
	ExecuteQuery(sql string, args ...interface{}) []map[string]string
}

type ClickHouseUtils struct {
}

func (self *ClickHouseUtils) ExecuteUpdate(sql string, args ...interface{}) {
	prepare, err := chsqldb.Prepare(sql)
	if err != nil {
		fmt.Println("prepare sql error:", err)
	}
	_, err1 := prepare.Exec(args...)
	if err1 != nil {
		fmt.Println("execute sql error:", err1)
	}
}

func (self *ClickHouseUtils) ExecuteQuery(sql string, args ...interface{}) []map[string]string {
	prepare, err := chsqldb.Prepare(sql)
	if err != nil {
		fmt.Println("sql error:", err)
	}
	rows, err1 := prepare.Query(args...)
	if err1 != nil {
		fmt.Println("sql error:", err1)
	}
	return self.Rows2List(rows, false)
}

func (self *ClickHouseUtils) Rows2List(rows *sql.Rows, fetchFirst bool) []map[string]string {
	res := []map[string]string{}
	cols, error := rows.Columns()
	if error != nil {
		fmt.Println("Get sql rows columns failed,", error)
	}
	vals := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	for k, _ := range vals {
		scans[k] = &vals[k]
	}

	for rows.Next() {
		err := rows.Scan(scans...)
		if err != nil {
			fmt.Println("get data failed, error:[%v]", err.Error())
		}
		resItem := make(map[string]string, len(cols))
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			// fmt.Printf(string(v))
			//这里把[]byte数据转成string
			resItem[key] = string(v)
		}
		res = append(res, resItem)
		if fetchFirst {
			return res
		}
	}
	return res
}
