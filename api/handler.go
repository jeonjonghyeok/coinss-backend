package api

import (
	"log"
	"net/http"

	"github.com/jeonjonghyeok/coinss-backend/psql"
	upbit "github.com/jeonjonghyeok/coinss-backend/token"
	"github.com/jeonjonghyeok/coinss-backend/vo"
)

// CollectHost godoc
// @Summary Host information collection.
// @Description If it already exists, the changeable information is updated, and in the case of a new host, it is created and returned.
// @Accept json
// @Produce json
// @Param apiver path string true "API version"
// @Param project_id path string true "Project ID"
// @Success 200 {object} vo.User
// @Router /{apiver}/projects/{project_id}/hosts [post]
func signup(w http.ResponseWriter, r *http.Request) {
	var user vo.User
	parseJSON(r.Body, &user)
	exists, err := psql.IsExistUser(user.Email)
	must(err)
	log.Println(exists, user)
	if exists {
		panic(existUserError)
	}

	id, err := psql.CreateUser(user)
	must(err)

	token, err := upbit.New(id, user.Accesskey, user.Secretkey)
	must(err)

	writeJSON(w, struct {
		Token string `json:"token"`
	}{token})

}
func signin(w http.ResponseWriter, r *http.Request) {
}
