package api

import (
	"log"
	"net/http"

	"github.com/jeonjonghyeok/coinss-backend/psql"
	upbit "github.com/jeonjonghyeok/coinss-backend/token"
	"github.com/jeonjonghyeok/coinss-backend/vo"
)

func signup(w http.ResponseWriter, r *http.Request) {
	log.Println("signup")
	var user vo.User
	parseJSON(r.Body, &user)
	log.Println(user)
	token, err := upbit.New(user.Accesskey, user.Secretkey)
	if err != nil {
		log.Println(err)
		return
	}
	id, err := psql.CreateUser(user)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(id)

	writeJSON(w, struct {
		Token string `json:"token"`
	}{token})

}
func signin(w http.ResponseWriter, r *http.Request) {
}
