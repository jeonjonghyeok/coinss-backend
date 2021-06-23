package rds

import (
	"encoding/json"
	"log"

	redis "github.com/go-redis/redis"
)

func GetCoinlist() (RespQuote Resp_Quote, err error) {

	rds_client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 접근 url 및 port
		Password: "",               // password ""값은 없다는 뜻
		DB:       0,                // 기본 DB 사용
	})

	val, err := rds_client.Get("price").Result()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	err = json.Unmarshal([]byte(val), &RespQuote)
	return
}
