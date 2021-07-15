package rds

import (
	"fmt"

	redis "github.com/go-redis/redis"
)

var rds *redis.Client

func Start() {
	fmt.Println("Redis Connect")
	go setCoins()
}

func db() *redis.Client {
	if rds == nil {
		rds = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379", // 접근 url 및 port
			Password: "",               // password ""값은 없다는 뜻
			DB:       0,                // 기본 DB 사용
		})
		_, err := rds.Ping().Result()
		if err != nil {
			panic(err)
		}

	}
	return rds
}
