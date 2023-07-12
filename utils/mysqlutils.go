package utils

import (
	"database/sql"
	"fmt"
	"strings"
)

type MysqlUtilsInterface interface {
	GetRecordById(tableName string, recordId int) map[string]string
	GetRecordList(tableName string, condition []SearchCond) []map[string]string
	AddRecord(tableName string, record map[string]string) int
	UpdateRecord(tableName string, update map[string]string, condition []SearchCond) bool
	DeleteRecordById(tableName string, recordId int)
	DeleteRecord(tableName string, condition []SearchCond)
}

type MysqlUtils struct {
}

func (self *MysqlUtils) GetRecordById(tableName string, recordId int) map[string]string {
	query := "select * from `" + tableName + "` where id=?"
	rows, error := Db.Query(query, recordId)
	if error != nil {
		fmt.Println("Exec sql error,", error)
	}
	result := self.Rows2List(rows, true)
	return result[0]
}

func (self *MysqlUtils) GetRecordList(tableName string, condition []SearchCond) []map[string]string {
	sqlstr, sqlstr2, args := SearchCond2Sql(condition)
	query := "select * from `" + tableName + "` "
	if len(condition) > 0 {
		query += " where " + sqlstr2
	}
	fmt.Println(sqlstr)
	fmt.Println(query)
	rows, error := Db.Query(query, args...)
	if error != nil {
		fmt.Println("Execute sql error,", error)
	}
	return self.Rows2List(rows, false)
}

func (self *MysqlUtils) AddRecord(tableName string, record map[string]string) int {
	keys := []string{}
	vals := []string{}
	vals2 := make([]interface{}, len(record))
	for s, s2 := range record {
		keys = append(keys, s)
		vals = append(vals, "?")
		vals2 = append(vals2, s2)
	}
	query := "insert into `" + tableName + "` (" + strings.Join(keys, ",") + ") values (" + strings.Join(vals, ",") + ")"
	record1, error := Db.Exec(query, vals2...)
	if error != nil {
		fmt.Println("Execute sql error,", error)
	}

	id, error := record1.LastInsertId()
	if error != nil {
		fmt.Println("Execute sql error,", error)
	}
	fmt.Println("Insert success,", id)
	return int(id)
}

func (self *MysqlUtils) UpdateRecord(tableName string, update map[string]string, condition []SearchCond) bool {
	res := true
	query := "update `" + tableName + "` set "
	updatesqlstr, updatesqlstr2, updateargs := Update2Sql(update)
	fmt.Println(updatesqlstr)
	query += updatesqlstr2
	if len(condition) > 0 {
		wheresqlstr, wheresqlstr2, whereargs := SearchCond2Sql(condition)
		fmt.Println(wheresqlstr)
		query += " where " + wheresqlstr2
		updateargs = append(updateargs, whereargs...)
	}
	fmt.Println(query)
	fmt.Println(updateargs)
	res1, error := Db.Exec(query, updateargs...)
	if error != nil {
		fmt.Println("Execute sql error,", error)
		res = false
	}
	row, error := res1.RowsAffected()
	if error != nil {
		fmt.Println("rows failed, ", error)
	}
	fmt.Println(row)
	return res
}

func (self *MysqlUtils) DeleteRecordById(tableName string, recordId int) {
	query := "delete from `" + tableName + "` where id=?"
	res, error := Db.Exec(query, recordId)
	if error != nil {
		fmt.Println("exec failed, ", error)
		return
	}
	row, error := res.RowsAffected()
	if error != nil {
		fmt.Println("rows failed, ", error)
	}
	fmt.Println(row)
}

func (self *MysqlUtils) DeleteRecord(tableName string, condition []SearchCond) {
	query := "delete from `" + tableName + "` "
	sqlstr, sqlstr2, args := SearchCond2Sql(condition)
	if len(condition) > 0 {
		fmt.Println(sqlstr)
		query += " where " + sqlstr2
	}
	res, error := Db.Exec(query, args...)
	if error != nil {
		fmt.Println("exec failed, ", error)
		return
	}
	row, error := res.RowsAffected()
	if error != nil {
		fmt.Println("rows failed, ", error)
	}
	fmt.Println(row)
}

func (self *MysqlUtils) Rows2List(rows *sql.Rows, fetchFirst bool) []map[string]string {
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
