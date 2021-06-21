package psql

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/lib/pq"
)

var db *sql.DB
var listener *pq.Listener
var ErrUnauthorized = errors.New("db: unauthorized")

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
