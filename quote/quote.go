package quote

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
)

func GetMarketPrice(client *http.Client, rds_client *redis.Client, req *http.Request) {
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
