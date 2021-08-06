package main

import (
	"github.com/jeonjonghyeok/coinss-backend/api"
	_ "github.com/jeonjonghyeok/coinss-backend/docs"
	"github.com/jeonjonghyeok/coinss-backend/rds"
)

func main() {
	//redis start
	rds.Start()
	api.Start()
}
