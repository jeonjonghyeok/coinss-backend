package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeonjonghyeok/coinss-backend/model"
	"github.com/jeonjonghyeok/coinss-backend/psql"
	upbit "github.com/jeonjonghyeok/coinss-backend/token"
)

// Register godoc
// @Summary Register
// @Description get string by ID
// @ID get-string-by-int
// @Tags user
// @Accept  json
// @Produce  json
// @Param model.User body model.User true "User"
// @Success 200 {object} model.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /user/signup [post]
func (c *Controller) AddUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.String(http.StatusBadRequest, "bad request")
		return
	}
	log.Println(user)
	exists, err := psql.IsExistUser(user.Email)
	if err != nil {
		panic(err)
	}
	log.Println(exists, user)
	if exists {
		panic("duplicate user")
	}

	id, err := psql.CreateUser(user)
	if err != nil {
		panic(err)
	}

	token, err := upbit.New(id, user.Accesskey, user.Secretkey)
	if err != nil {
		panic(err)
	}

	ctx.String(200, token)
}
