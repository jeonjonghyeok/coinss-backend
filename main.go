package main

import (
	"github.com/jeonjonghyeok/coinss-backend/api"
	_ "github.com/jeonjonghyeok/coinss-backend/docs"
	"github.com/jeonjonghyeok/coinss-backend/psql"
	"github.com/jeonjonghyeok/coinss-backend/rds"
)

func main() {
	//postgres 연결
	psql.Start()
	//redis 연결
	rds.Start()
	api.Start()
}
