package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/jeonjonghyeok/exchange-go/quote"
	"github.com/jeonjonghyeok/exchange-go/rds"
	"github.com/jeonjonghyeok/exchange-go/vo"
)

func main() {
	rds_client, err := rds.InitializeRedisClient()
	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "50")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "1cc3129e-e274-48fc-88e0-0fba6e437cbd")
	req.URL.RawQuery = q.Encode()

	go quote.GetMarketPrice(client, rds_client, req)
	time.Sleep(time.Second * 2)
	go readPump(rds_client)
	fmt.Scanln()

}
func readPump(rds_client *redis.Client) {
	var RespQuote vo.Resp_Quote
	for {
		for i := 0; i < 50; i++ {
			val, err := rds_client.Get("price").Result()
			if err != nil {
				panic(err)
			}
			json.Unmarshal([]byte(val), &RespQuote)
			fmt.Print(RespQuote.Data[i].Symbol, " ")
			fmt.Println(RespQuote.Data[i].Quote.Usd.Price)

		}
		time.Sleep(time.Second * 10)
	}

}
