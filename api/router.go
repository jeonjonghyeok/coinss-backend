package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func API() http.Handler {
	router := mux.NewRouter()

	//user
	router.HandleFunc("/signup", signup).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/signin", signin).Methods(http.MethodPost, http.MethodOptions)

	router.Use(handlePanic)

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, r)
		})
	})
	router.Use(mux.CORSMethodMiddleware(router))

	return router
}
