package rds

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	redis "github.com/go-redis/redis"
	"github.com/jeonjonghyeok/coinss-backend/model"
	"github.com/jeonjonghyeok/coinss-backend/utils"
)

type resCoin struct {
	Market      string `json:"market"`
	KoreanName  string `json:"korean_name"`
	EnglishName string `json:"english_name"`
}

type resCoinPrice struct {
	Market     string  `json:"market"`
	HighPrice  float32 `json:"high_price"`
	LowPrice   float32 `json:"low_price"`
	TradePrice float32 `json:"trade_price"`
}

func GetCoins() (coins []model.Coin, err error) {
	const URL = "https://api.upbit.com/v1/market/all"
	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	var resCoins []resCoin
	respBody, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(respBody), &resCoins)

	var markets string
	for _, coin := range resCoins {
		if markets == "" {
			markets += coin.Market
		} else {
			markets += "," + coin.Market
		}
		coins = append(coins, model.Coin{Symbol: coin.EnglishName})
	}
	coinsPrice := getCoinsPrice(markets)
	for i, coin := range coinsPrice {
		coins[i].Market = coin.Market
		coins[i].Price = float32(coin.TradePrice)
		coins[i].HighPrice = coin.HighPrice
		coins[i].LowPrice = coin.LowPrice
		coinBytes, err := json.Marshal(coins[i])
		utils.HandleErr(err)
		utils.HandleErr(rds.Set(coins[i].Symbol, coinBytes, 0).Err())
	}
	return

}

func getCoinsPrice(markets string) []*resCoinPrice {
	fmt.Println(markets)
	const URL = "https://api.upbit.com/v1/ticker"
	var coins []*resCoinPrice
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	q := url.Values{}
	q.Add("markets", markets)
	req.Header.Set("Accepts", "application/json")
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	utils.HandleErr(err)
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))
	utils.HandleErr(json.Unmarshal([]byte(respBody), &coins))
	return coins
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
