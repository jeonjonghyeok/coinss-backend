package rds

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/go-redis/redis"
)

const (
	START   = "1"
	LIMIT   = "50"
	CONVERT = "USD"
	TOKEN   = "1cc3129e-e274-48fc-88e0-0fba6e437cbd"
	URL     = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest"
)

func GetMarketPrice(rds_client *redis.Client) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("start", START)
	q.Add("limit", LIMIT)
	q.Add("convert", CONVERT)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", TOKEN)
	req.URL.RawQuery = q.Encode()

	for {
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request to server")
			os.Exit(1)
		}
		fmt.Println(resp.Status)
		respBody, _ := ioutil.ReadAll(resp.Body)

		err = rds_client.Set("price", respBody, 0).Err()
		if err != nil {
			log.Println(err)
		}

		time.Sleep(time.Second * 60)
	}

}
