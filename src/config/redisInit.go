package config

import (
	"fmt"
	util "testing/src/util"

	"github.com/go-redis/redis"
)

var rs *redis.Client

func GetRS() *redis.Client {
	if rs == nil {
		InitRedis()
	}
	return rs
}

func InitRedis() *redis.Client {
	rs = redis.NewClient(&redis.Options{
		Addr:     util.GetElement("redis.host") + ":" + util.GetElement("redis.port"),
		Password: "",                                               // no password set
		DB:       util.StrToInt(util.GetElement("redis.database")), // use default DB
	})
	fmt.Println("redis init.....")
	pong, err := rs.Ping().Result()
	fmt.Println(pong, err)
	return rs
}
