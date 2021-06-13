package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func API() http.Handler {
	log.Println("Router Start")
	router := mux.NewRouter()
	//user
	router.HandleFunc("/signup", signup).Methods(http.MethodPost, http.MethodOptions)

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, r)
		})
	})
	router.Use(mux.CORSMethodMiddleware(router))

	return router
}
