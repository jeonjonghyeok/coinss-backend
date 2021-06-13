package psql

import (
	"database/sql"
	"log"
	"time"

	"github.com/lib/pq"
)

var db *sql.DB
var listener *pq.Listener

func Connect(url string) error {
	log.Println("PSQL Connect")
	c, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}
	db = c

	listener = pq.NewListener(url,
		10*time.Second, time.Minute, func(ev pq.ListenerEventType, err error) {
			if err != nil {
				panic(err)
			}
		})

	return nil

}
