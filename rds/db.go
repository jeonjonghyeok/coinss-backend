package rds

import (
	redis "github.com/go-redis/redis"
)

func InitializeRedisClient() (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 접근 url 및 port
		Password: "",               // password ""값은 없다는 뜻
		DB:       0,                // 기본 DB 사용
	})

	_, err := client.Ping().Result()

	return client, err
}

/*
func ExecuteSomething() {

	Client, err := initializeRedisClient()

	if nil != err {
		panic(err)
	}

	err = Client.Set("value", "2", 0).Err() // 마지막 인자는 만료 시간 Redis 데이터의 만료 시간을 지정할 수 있다.
	if err != nil {
		panic(err)
	}

	val, err := Client.Get("value").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("value : ", val)
}
*/
