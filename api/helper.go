package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/jeonjonghyeok/coinss-backend/psql"
)

func parseJSON(r io.Reader, v interface{}) {
	if err := json.NewDecoder(r).Decode(v); err != nil {
		log.Panic(err)
	}
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Panic(err)
	}
}
func must(err error) {
	if err == psql.ErrUnauthorized {
		panic(unauthorizedError)
	} else if err != nil {
		log.Println("internal error:", err)
		panic(internalError)
	}
}
