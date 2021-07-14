package rds

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/jeonjonghyeok/coinss-backend/model"
	"github.com/jeonjonghyeok/coinss-backend/utils"
)

type nameResponse struct {
	Market      string `json:"market"`
	KoreanName  string `json:"korean_name"`
	EnglishName string `json:"english_name"`
}

type priceResponse struct {
	Market     string  `json:"market"`
	HighPrice  float32 `json:"high_price"`
	LowPrice   float32 `json:"low_price"`
	TradePrice float32 `json:"trade_price"`
}

func getCoinName() (coins []*model.Coin, markets string, err error) {
	const URL = "https://api.upbit.com/v1/market/all"
	resp, err := http.Get(URL)
	if err != nil {
		return nil, "", err
	}
	var resCoins []nameResponse
	respBody, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(respBody), &resCoins)
	if err != nil {
		return nil, "", err
	}

	for _, coin := range resCoins {
		if markets == "" {
			markets += coin.Market
		} else {
			markets += "," + coin.Market
		}
		coins = append(coins, &model.Coin{EnglishName: coin.EnglishName, KoreanName: coin.KoreanName})
	}
	return
}

func getPriceByName(markets string) (coins []*priceResponse, err error) {
	const URL = "https://api.upbit.com/v1/ticker"
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("markets", markets)
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Accepts", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(respBody, &coins); err != nil {
		return nil, err
	}
	return
}

func setCoins() {
	coins, markets, err := getCoinName()
	utils.HandleErr(err)

	price_coins, err := getPriceByName(markets)
	utils.HandleErr(err)

	for i, coin := range price_coins {
		coins[i].Market = coin.Market
		coins[i].Price = float32(coin.TradePrice)
		coins[i].HighPrice = coin.HighPrice
		coins[i].LowPrice = coin.LowPrice
		coinBytes, err := json.Marshal(coins[i])
		utils.HandleErr(err)
		utils.HandleErr(db().Set(coins[i].EnglishName, coinBytes, 0).Err())
	}
}

func GetCoins(names string) (coins []model.Coin, err error) {
	var val string
	splitNames := strings.Split(names, ",")
	for _, name := range splitNames {
		val, err = db().Get(name).Result()
		if err != nil {
			log.Println(err)
			return
		}
		var coin model.Coin

		if err = json.Unmarshal([]byte(val), &coin); err != nil {
			return
		}
		coins = append(coins, coin)
	}
	return
}
