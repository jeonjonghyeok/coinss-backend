package server

import (
	"log"
	"net/http"

	"github.com/jeonjonghyeok/coinss-backend/api"
	"github.com/jeonjonghyeok/coinss-backend/psql"
	"github.com/jeonjonghyeok/coinss-backend/rds"
)

type Config struct {
	Address string
	Url     string
}

func ListenAndServe(c Config) error {
	log.Println("ListenAndServe")
	if err := psql.Connect(c.Url); err != nil {
		log.Fatal(err)
	}
	if err := rds.Connect(); err != nil {
		log.Fatal(err)
	}

	return http.ListenAndServe(c.Address, api.API())
}
