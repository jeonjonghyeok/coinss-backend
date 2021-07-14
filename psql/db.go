package psql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
)

const (
	DB_USER     = "jjh"
	DB_PASSWORD = "jjh"
	DB_NAME     = "jjh"
)

var db *sql.DB
var listener *pq.Listener
var ErrUnauthorized = errors.New("db: unauthorized")

func Start() {
	if err := Connect(fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)); err != nil {
		log.Fatal(err)
	}
}

func Connect(url string) error {
	log.Println("PSQL Connect")
	c, err := sql.Open("postgres", url)
	if err != nil {
		log.Println(err)
		panic(ErrUnauthorized)
	}
	db = c

	listener = pq.NewListener(url,
		10*time.Second, time.Minute, func(ev pq.ListenerEventType, err error) {
			if err != nil {
				log.Println(err)
				panic(ErrUnauthorized)
			}
		})

	return nil

}
