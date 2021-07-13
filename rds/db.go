package rds

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	redis "github.com/go-redis/redis"
	"github.com/jeonjonghyeok/coinss-backend/model"
)

var rds *redis.Client

func Start() {
	DB()
	//go GetMarketPrice(rds)
	//time.Sleep(time.Second * 2)
	//go readPump(rds_client)
}

func DB() {
	if rds == nil {
		log.Println("rds nil")
	}
	log.Println("RDS Connect")
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

func readPump(rds_client *redis.Client) {
	var RespQuote model.Resp_Quote
	for {
		val, err := rds_client.Get("price").Result()
		if err != nil {
			log.Println(err)
			panic(err)
		}
		json.Unmarshal([]byte(val), &RespQuote)
		for i := 0; i < 50; i++ {
			fmt.Print(RespQuote.Data[i].Symbol, " ")
			fmt.Print(RespQuote.Data[i].Name, " ")
			fmt.Println(RespQuote.Data[i].Quote.Usd.Price)

		}
		time.Sleep(time.Second * 10)
	}

}
