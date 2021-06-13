package api

import (
	"net/http"

	"github.com/jeonjonghyeok/coinss-backend/vo"
)

type usernamePassword struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func signup(w http.ResponseWriter, r *http.Request) {
	var req usernamePassword
	var user vo.User
	parseJSON(r.Body, &req)
	user.Username = req.Username
	user.Password = req.Password
	writeJSON(w, struct {
		Token string `json:"token"`
	}{"success"})

}
