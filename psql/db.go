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

func DB() *sql.DB {
	if db == nil {
		url := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
			DB_USER, DB_PASSWORD, DB_NAME)
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
	}
	return db

}
