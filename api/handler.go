package api

import (
	"net/http"

	"github.com/jeonjonghyeok/coinss-backend/psql"
	upbit "github.com/jeonjonghyeok/coinss-backend/token"
	"github.com/jeonjonghyeok/coinss-backend/vo"
)

func signup(w http.ResponseWriter, r *http.Request) {
	var user vo.User
	parseJSON(r.Body, &user)

	token, err := upbit.New(user.Accesskey, user.Secretkey)
	must(err)

	must(psql.CreateUser(user))

	writeJSON(w, struct {
		Token string `json:"token"`
	}{token})

}
func signin(w http.ResponseWriter, r *http.Request) {
}
