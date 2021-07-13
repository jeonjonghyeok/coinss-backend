package main

import (
	"github.com/jeonjonghyeok/coinss-backend/api"
	_ "github.com/jeonjonghyeok/coinss-backend/docs"
	"github.com/jeonjonghyeok/coinss-backend/rds"
)

const (
	DB_USER     = "jjh"
	DB_PASSWORD = "jjh"
	DB_NAME     = "jjh"
)

func main() {
	//postgres 연결
	/*if err := psql.Connect(fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)); err != nil {
		log.Fatal(err)
	}
	*/
	//redis 연결
	rds.Start()
	api.Start()
}
