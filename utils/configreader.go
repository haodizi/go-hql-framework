package utils

import (
	iniconf "github.com/clod-moon/goconf"
	conf2 "github.com/haodizi/go-hql-framework/conf"
)

func InitConfigFromIni() {
	conf := iniconf.InitConfig("conf/config.ini")

	conf2.DbConf.Host = conf.GetValue("database", "host")
	conf2.DbConf.Username = conf.GetValue("database", "username")
	conf2.DbConf.Password = conf.GetValue("database", "password")
	conf2.DbConf.Port = conf.GetValue("database", "port")
	conf2.DbConf.DbName = conf.GetValue("database", "dbname")
	conf2.DbConf.Charset = conf.GetValue("database", "charset")
}

func InitRedisConfigFromIni() {
	conf := iniconf.InitConfig("conf/config.ini")

	conf2.RedisConf.Host = conf.GetValue("redis", "host")
	conf2.RedisConf.Port = conf.GetValue("redis", "port")
	conf2.RedisConf.Password = conf.GetValue("redis", "password")
}

func InitClickHouseConfigFromIni() {
	conf := iniconf.InitConfig("conf/config.ini")

	conf2.ClickHouseConf.Host = conf.GetValue("clickhouse", "host")
	conf2.ClickHouseConf.Username = conf.GetValue("clickhouse", "username")
	conf2.ClickHouseConf.Password = conf.GetValue("clickhouse", "password")
	conf2.ClickHouseConf.Port = conf.GetValue("clickhouse", "port")
	conf2.ClickHouseConf.DbName = conf.GetValue("clickhouse", "dbname")
}
