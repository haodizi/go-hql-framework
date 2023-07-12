package utils

import (
	"awesomeProject/conf"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var RedisConn redis.Conn

func init() {
	InitRedisConfigFromIni()
	dsn := fmt.Sprintf("%s:%s", conf.RedisConf.Host, conf.RedisConf.Port)
	c, error := redis.Dial("tcp", dsn)
	if error != nil {
		fmt.Println("Conn redis failed,", error)
		return
	}
	_, error1 := c.Do("AUTH", conf.RedisConf.Password)
	if error1 != nil {
		fmt.Println("Auth failed,", error1)
		return
	}
	RedisConn = c
}

func RedisSet(key string, value string, expire int) {
	_, error := RedisConn.Do("Set", key, value)
	if error != nil {
		fmt.Println("Set into redis failed,", error)
		return
	}
	_, error1 := RedisConn.Do("expire", key, expire)
	if error1 != nil {
		fmt.Println("Set expire failed,", error1)
		return
	}
}

func RedisGet(key string) string {
	r, error := redis.String(RedisConn.Do("Get", key))
	if error != nil {
		fmt.Println("get "+key+" failed,", error)
		r = ""
	}
	return r
}

func RedisDelete(key string) {
	_, error := RedisConn.Do("Del", key)
	if error != nil {
		fmt.Println("Delete redis key failed,", error)
		return
	}
}
