package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type simpleError struct {
	message string
	status  int
}

type apiErrorInterface interface {
	Write(w http.ResponseWriter)
}

func handlePanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
				if e, ok := r.(apiErrorInterface); ok {
					e.Write(w)
				} else {
					internalError.Write(w)
				}
			}
		}()
		next.ServeHTTP(w, r)
	})

}
func (e simpleError) Write(w http.ResponseWriter) {
	w.WriteHeader(e.status)
	json.NewEncoder(w).Encode(e.message)
}

var internalError = simpleError{
	message: "Internal Error",
	status:  http.StatusInternalServerError,
}
