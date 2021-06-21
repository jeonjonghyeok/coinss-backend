package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeonjonghyeok/coinss-backend/model"
	"github.com/jeonjonghyeok/coinss-backend/psql"
	t "github.com/jeonjonghyeok/coinss-backend/token"
)

// Register godoc
// @Summary Register
// @Description get string by ID
// @ID post-user-signup
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
	/*
		token, err := upbit.NewUpbit(id, user.Accesskey, user.Secretkey)
		if err != nil {
			panic(err)
		}
	*/
	token, err := t.New(id)
	if err != nil {
		panic(err)
	}

	ctx.String(200, token)
}

type emailPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login godoc
// @Summary Login
// @Description get string by ID
// @ID post-user-signin
// @Tags user
// @Accept  json
// @Produce  json
// @Param emailPassword body emailPassword true "User"
// @Success 200 {object} model.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /user/signin [post]
func (c *Controller) SigninUser(ctx *gin.Context) {
	var user emailPassword
	if err := ctx.BindJSON(&user); err != nil {
		ctx.String(http.StatusBadRequest, "bad request")
		return
	}
	log.Println(user)

	id, err := psql.FindUser(user.Email, user.Password)
	if err != nil {
		panic(err)
	}
	token, err := t.New(id)
	if err != nil {
		panic(err)
	}
	ctx.String(200, token)
}
