package rds

import (
	redis "github.com/go-redis/redis"
)

var rds *redis.Client

func Start() {
	go setCoins()
	//go readPump(rds_client)
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

/*
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

*/
