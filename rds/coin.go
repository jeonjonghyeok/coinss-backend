package rds

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	redis "github.com/go-redis/redis"
	"github.com/jeonjonghyeok/coinss-backend/model"
)

func GetCoin() (coin []model.Coin, err error) {
	const URL = "https://api.upbit.com/v1/market/all"
	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(respBody))

	err = json.Unmarshal([]byte(respBody), &coin)
	return
}

func GetCoinlist() (RespQuote model.Resp_Quote, err error) {
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
