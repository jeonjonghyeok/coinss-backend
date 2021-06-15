package main

import (
	"fmt"
	"log"

	"github.com/jeonjonghyeok/coinss-backend/server"
)

const (
	DB_USER     = "jjh"
	DB_PASSWORD = "jjh"
	DB_NAME     = "jjh"
)

// @title Example API
// @version 0.0.2
// @description This is a Example api server
// @contact.name Request permission of Example API
// @contact.url http://www.yonghochoi.com
// @contact.email yongho1037@gmail.com
// @host localhost
// @BasePath /api/v1
func main() {

	if err := server.ListenAndServe(
		server.Config{
			Address: ":5000",
			Url: fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
				DB_USER, DB_PASSWORD, DB_NAME),
		}); err != nil {
		log.Fatal(err)
	}
}
