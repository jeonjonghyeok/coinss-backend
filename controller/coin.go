package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeonjonghyeok/coinss-backend/model"
	"github.com/jeonjonghyeok/coinss-backend/psql"
	"github.com/jeonjonghyeok/coinss-backend/rds"
	upbit "github.com/jeonjonghyeok/coinss-backend/token"
	"github.com/jeonjonghyeok/coinss-backend/utils"
)

// Coin-List godoc
// @Summary Coin-List
// @Description get coinlist
// @Tags coin
// @Accept  json
// @Produce  json
// @Param favorite body favorite true "Faavorite"
// @Success 200 {object} model.Coin
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/coin/list [get]
func (c *Controller) Coins(ctx *gin.Context) {
	var f favorite
	utils.HandleErr(ctx.BindJSON(&f))

	coins, err := rds.GetCoins(f.Name)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, coins)
}

type header struct {
	Token string `header:"Token"`
}

// Coin-Wallet godoc
// @Summary Coin-Wallet
// @Description get coinwallet
// @Tags coin
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} model.Wallet
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/coin/wallet [get]
func (c *Controller) Wallet(ctx *gin.Context) {
	const URL = "https://api.upbit.com/v1/accounts"

	h := header{}
	if err := ctx.ShouldBindHeader(&h); err != nil {
		panic(err)
	}
	id, err := upbit.Parse(h.Token)
	if err != nil {
		panic(err)
	}

	user, err := psql.FindUserById(id)
	token, err := upbit.NewUpbit(id, user.Accesskey, user.Secretkey)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	log.Println(resp.Status)
	if resp.Status == "401 Unauthorized" {
		panic("등록되지 않은 IP주소 입니다.")
	}
	respBody, _ := ioutil.ReadAll(resp.Body)

	var wallet []model.Wallet
	if err := json.Unmarshal(respBody, &wallet); err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, wallet)

}

type favorite struct {
	Name string `form:"name" json:"name" example:"Bitcoin" binding:"required"`
}

// Favorite godoc
// @Summary Register Favority Coin
// @Description 관심코인 등록
// @ID post-coin-favorite
// @Tags coin
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param favorite body favorite true "Faavorite"
// @Success 200 {object} model.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/coin/favorite [post]
func (c *Controller) Favorite(ctx *gin.Context) {
	h := header{}
	utils.HandleErr(ctx.ShouldBindHeader(&h))

	id, err := upbit.Parse(h.Token)
	utils.HandleErr(err)

	var f favorite
	utils.HandleErr(ctx.BindJSON(&f))
	names := psql.GetFavorites(id)
	if names == "" {
		names = f.Name
	} else {
		names += "," + f.Name
	}
	utils.HandleErr(psql.Favorite(id, names))
	ctx.String(http.StatusOK, "Success")
}

// Favorite godoc
// @Summary Register Favority Coin
// @Description 관심코인 조회
// @ID post-coin-favorites
// @Tags coin
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} model.User
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/coin/favorites [get]
func (c *Controller) Favorites(ctx *gin.Context) {
	log.Println("favorites call")
	h := header{}
	utils.HandleErr(ctx.ShouldBindHeader(&h))

	id, err := upbit.Parse(h.Token)
	utils.HandleErr(err)

	names := psql.GetFavorites(id)
	if names == "" {
		ctx.JSON(http.StatusOK, nil)
		return
	}
	log.Println("names:", names)
	coins, err := rds.GetCoins(names)
	utils.HandleErr(err)
	log.Println("coins:", coins)

	ctx.JSON(http.StatusOK, coins)
}
