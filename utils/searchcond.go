package utils

import (
	"fmt"
	"strings"
)

const COND_TYPE_EQ string = "="
const COND_TYPE_NEQ string = "<>"
const COND_TYPE_GT string = ">"
const COND_TYPE_EGT string = ">="
const COND_TYPE_LT string = "<"
const COND_TYPE_ELT string = "<="
const COND_TYPE_LIKE string = "like"
const COND_TYPE_BETWEEN string = "between"
const COND_TYPE_NOT_BETWEEN string = "not between"
const COND_TYPE_IN string = "in"
const COND_TYPE_NOT_IN string = "not in"

type SearchCond struct {
	Field    string
	Type     string
	Value    string
	ValueArr []string
}

// 把条件数组转化为sql条件语句，返回1 原生拼装的sql 2 ？占位的sql 3 绑定数据的切片
func SearchCond2Sql(cond []SearchCond) (string, string, []interface{}) {
	res := ""
	res2 := ""
	var args []interface{}
	for i, searchCond := range cond {
		fmt.Println(i)
		var sqlstr, sqlstr2 string
		if searchCond.Field == "_string" {
			sqlstr = searchCond.Value + " and "
			sqlstr2 = searchCond.Value + " and "
		} else {
			var checkRes bool = searchCond.Type == COND_TYPE_EQ || searchCond.Type == COND_TYPE_NEQ
			checkRes = checkRes || searchCond.Type == COND_TYPE_GT || searchCond.Type == COND_TYPE_EGT
			checkRes = checkRes || searchCond.Type == COND_TYPE_LT || searchCond.Type == COND_TYPE_ELT
			checkRes = checkRes || searchCond.Type == COND_TYPE_LIKE
			if checkRes == true {
				sqlstr = searchCond.Field + searchCond.Type + "'" + searchCond.Value + "' and "
				sqlstr2 = searchCond.Field + searchCond.Type + "? and "
				args = append(args, searchCond.Value)
			} else if searchCond.Type == COND_TYPE_BETWEEN {
				sqlstr = "(" + searchCond.Field + " " + searchCond.Type + "'" + searchCond.ValueArr[0] + "' and '" + searchCond.ValueArr[1] + "') and "
				sqlstr2 = "(" + searchCond.Field + " " + searchCond.Type + "? and ?) and "
				args = append(args, searchCond.ValueArr[0])
				args = append(args, searchCond.ValueArr[1])
			} else if searchCond.Type == COND_TYPE_NOT_BETWEEN {
				sqlstr = "(" + searchCond.Field + "<='" + searchCond.ValueArr[0] + "' or " + searchCond.Field + ">='" + searchCond.ValueArr[1] + "') and "
				sqlstr2 = "(" + searchCond.Field + "<=? or " + searchCond.Field + ">=?) and "
				args = append(args, searchCond.ValueArr[0])
				args = append(args, searchCond.ValueArr[1])
			} else {
				var temp string
				for i2, s := range searchCond.ValueArr {
					fmt.Println(i2)
					temp += "'" + s + "',"
				}
				temp = strings.TrimRight(temp, ",")
				sqlstr = searchCond.Field + " " + searchCond.Type + " (" + temp + ") and "
				sqlstr2 = searchCond.Field + " " + searchCond.Type + " (" + temp + ") and "
			}
		}
		res += sqlstr
		res2 += sqlstr2
	}
	res = strings.TrimRight(res, " and ")
	res2 = strings.TrimRight(res2, " and ")
	return res, res2, args
}

// 把更新数组转化为sql语句, 返回1 原生sql 2 ?占位的sql 3 绑定数据的切片
func Update2Sql(update map[string]string) (string, string, []interface{}) {
	res := ""
	res2 := ""
	var args []interface{}
	for s, s2 := range update {
		res += s + "='" + s2 + "',"
		res2 += s + "=?,"
		args = append(args, s2)
	}
	res = strings.TrimRight(res, ",")
	res2 = strings.TrimRight(res2, ",")
	return res, res2, args
}
